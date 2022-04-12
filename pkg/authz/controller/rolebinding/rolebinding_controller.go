/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2019 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

package rolebinding

import (
	"context"
	"fmt"
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	"reflect"
	"time"
	apiauthzv1 "tkestack.io/tke/api/authz/v1"
	clientset "tkestack.io/tke/api/client/clientset/versioned"
	platformversionedclient "tkestack.io/tke/api/client/clientset/versioned/typed/platform/v1"
	authzv1informer "tkestack.io/tke/api/client/informers/externalversions/authz/v1"
	authzv1 "tkestack.io/tke/api/client/listers/authz/v1"
	"tkestack.io/tke/pkg/authz/constant"
	"tkestack.io/tke/pkg/authz/controller/rolebinding/deletion"
	"tkestack.io/tke/pkg/authz/controller/rolebinding/policyrolecache"
	authzprovider "tkestack.io/tke/pkg/authz/provider"
	controllerutil "tkestack.io/tke/pkg/controller"
	"tkestack.io/tke/pkg/util/log"
	"tkestack.io/tke/pkg/util/metrics"
)

const (
	// appDeletionGracePeriod is the time period to wait before processing a received channel event.
	// This allows time for the following to occur:
	// * lifecycle admission plugins on HA apiservers to also observe a channel
	//   deletion and prevent new objects from being created in the terminating channel
	// * non-leader etcd servers to observe last-minute object creations in a channel
	//   so this controller's cleanup can actually clean up all objects
	appDeletionGracePeriod = 5 * time.Second
)

const (
	controllerName = "rolebinding-controller"
)

type Controller struct {
	client             clientset.Interface
	platformClient     platformversionedclient.PlatformV1Interface
	queue              workqueue.RateLimitingInterface
	roleLister         authzv1.RoleLister
	rolebindingLister  authzv1.RoleBindingLister
	policyLister       authzv1.PolicyLister
	roleSynced         cache.InformerSynced
	rolebindingSynced  cache.InformerSynced
	policySynced       cache.InformerSynced
	rolebindingDeleter deletion.RoleBindingDeleter
	stopCh             <-chan struct{}
}

// NewController creates a new Controller object.
func NewController(
	client clientset.Interface,
	platformClient platformversionedclient.PlatformV1Interface,
	roleInformer authzv1informer.RoleInformer,
	rolebindingInformer authzv1informer.RoleBindingInformer,
	policyInformer authzv1informer.PolicyInformer,
	resyncPeriod time.Duration,
) *Controller {
	// create the controller so we can inject the enqueue function
	controller := &Controller{
		client:             client,
		platformClient:     platformClient,
		queue:              workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), controllerName),
		rolebindingDeleter: deletion.New(client),
	}
	if client != nil &&
		client.AuthzV1().RESTClient() != nil &&
		!reflect.ValueOf(client.AuthzV1().RESTClient()).IsNil() &&
		client.AuthzV1().RESTClient().GetRateLimiter() != nil {
		_ = metrics.RegisterMetricAndTrackRateLimiterUsage(controllerName, client.AuthzV1().RESTClient().GetRateLimiter())
	}

	rolebindingInformer.Informer().AddEventHandlerWithResyncPeriod(
		cache.FilteringResourceEventHandler{
			Handler: cache.ResourceEventHandlerFuncs{
				AddFunc: controller.enqueue,
				UpdateFunc: func(oldObj, newObj interface{}) {
					old, ok1 := oldObj.(*apiauthzv1.RoleBinding)
					cur, ok2 := newObj.(*apiauthzv1.RoleBinding)
					if ok1 && ok2 && controller.needsUpdate(old, cur) {
						controller.enqueue(newObj)
					}
				},
				DeleteFunc: controller.enqueue,
			},
			FilterFunc: func(obj interface{}) bool {
				return true
			},
		},
		resyncPeriod,
	)

	roleInformer.Informer().AddEventHandlerWithResyncPeriod(
		cache.FilteringResourceEventHandler{
			Handler: cache.ResourceEventHandlerFuncs{
				AddFunc: func(obj interface{}) {
					policyrolecache.Cache.PutByRole(obj.(*apiauthzv1.Role))
					controller.addRole(obj)
				},
				UpdateFunc: func(oldObj, newObj interface{}) {
					old, ok1 := oldObj.(*apiauthzv1.Role)
					cur, ok2 := newObj.(*apiauthzv1.Role)
					if ok1 && ok2 && controller.needsAddRole(old, cur) {
						policyrolecache.Cache.UpdateByRole(cur)
						controller.addRole(cur)
					}
				},
				DeleteFunc: func(obj interface{}) {
					policyrolecache.Cache.DeleteRole(obj.(*apiauthzv1.Role))
				},
			},
			FilterFunc: func(obj interface{}) bool {
				return true
			},
		},
		resyncPeriod,
	)

	policyInformer.Informer().AddEventHandlerWithResyncPeriod(
		cache.FilteringResourceEventHandler{
			Handler: cache.ResourceEventHandlerFuncs{
				AddFunc: controller.addPolicy,
				UpdateFunc: func(oldObj, newObj interface{}) {
					old, ok1 := oldObj.(*apiauthzv1.Policy)
					cur, ok2 := newObj.(*apiauthzv1.Policy)
					if ok1 && ok2 && controller.needsAddPolicy(old, cur) {
						controller.addPolicy(cur)
					}
				},
				DeleteFunc: func(obj interface{}) {
					controller.addPolicy(obj)
					policyrolecache.Cache.DeletePolicy(obj.(*apiauthzv1.Policy))
				},
			},
			FilterFunc: func(obj interface{}) bool {
				return true
			},
		},
		resyncPeriod,
	)

	controller.roleLister = roleInformer.Lister()
	controller.roleSynced = roleInformer.Informer().HasSynced
	controller.rolebindingLister = rolebindingInformer.Lister()
	controller.rolebindingSynced = rolebindingInformer.Informer().HasSynced
	controller.policyLister = policyInformer.Lister()
	controller.policySynced = policyInformer.Informer().HasSynced
	return controller
}

func (c *Controller) addPolicy(obj interface{}) {
	pol := obj.(*apiauthzv1.Policy)
	roleSet := policyrolecache.Cache.GetRolesByPolicy(pol)
	for roleName, _ := range roleSet {
		ns, name, _ := cache.SplitMetaNamespaceKey(roleName)
		role, err := c.roleLister.Roles(ns).Get(name)
		if err != nil {
			log.Warnf("Unable to get role '%/%s', err: '%v'", ns, name, err)
			continue
		}
		c.addRole(role)
	}
}

func (c *Controller) needsAddPolicy(old *apiauthzv1.Policy, new *apiauthzv1.Policy) bool {
	if old.UID != new.UID {
		return true
	}
	if !reflect.DeepEqual(old.Rules, new.Rules) {
		return true
	}
	return false
}

func (c *Controller) addRole(obj interface{}) {
	role := obj.(*apiauthzv1.Role)
	var (
		rbs []*apiauthzv1.RoleBinding
		err error
	)
	selector := labels.SelectorFromSet(map[string]string{
		constant.RoleNamespace: role.Namespace,
		constant.RoleName:      role.Name,
	})
	if role.Namespace == "" {
		rbs, err = c.rolebindingLister.List(selector)
	} else {
		rbs, err = c.rolebindingLister.RoleBindings(role.Namespace).List(selector)
	}
	if err != nil {
		log.Warnf("Failed to list rolebindings by role %s/%s", role.Namespace, role.Name)
		return
	}
	for _, rb := range rbs {
		c.enqueue(rb)
	}
}

func (c *Controller) needsAddRole(old *apiauthzv1.Role, new *apiauthzv1.Role) bool {
	if old.UID != new.UID {
		return true
	}
	if !reflect.DeepEqual(old.Policies, new.Policies) {
		return true
	}
	return false
}

func (c *Controller) enqueue(obj interface{}) {
	key, err := controllerutil.KeyFunc(obj)
	if err != nil {
		log.Error("Couldn't get key for object", log.Any("object", obj), log.Err(err))
		return
	}
	c.queue.AddAfter(key, appDeletionGracePeriod)
}

func (c *Controller) needsUpdate(old *apiauthzv1.RoleBinding, new *apiauthzv1.RoleBinding) bool {
	if old.UID != new.UID {
		return true
	}
	if !reflect.DeepEqual(old.Spec, new.Spec) {
		return true
	}
	return false
}

// Run will set up the event handlers for types we are interested in, as well
// as syncing informer caches and starting workers.
func (c *Controller) Run(workers int, stopCh <-chan struct{}) {
	defer runtime.HandleCrash()
	defer c.queue.ShutDown()

	// Start the informer factories to begin populating the informer caches
	log.Info("Starting rolebinding controller")
	defer log.Info("Shutting down rolebinding controller")

	if ok := cache.WaitForCacheSync(stopCh, c.roleSynced, c.rolebindingSynced, c.policySynced); !ok {
		log.Error("Failed to wait for rolebinding caches to sync")
		return
	}

	c.stopCh = stopCh
	for i := 0; i < workers; i++ {
		go wait.Until(c.worker, time.Second, stopCh)
	}

	<-stopCh
}

// worker processes the queue of app objects.
// Each app can be in the queue at most once.
// The system ensures that no two workers can process
// the same app at the same time.
func (c *Controller) worker() {
	workFunc := func() bool {
		key, quit := c.queue.Get()
		if quit {
			return true
		}
		defer c.queue.Done(key)

		err := c.syncItem(key.(string))
		if err == nil {
			// no error, forget this entry and return
			c.queue.Forget(key)
			return false
		}

		// rather than wait for a full resync, re-add the app to the queue to be processed
		c.queue.AddRateLimited(key)
		runtime.HandleError(err)
		return false
	}

	for {
		quit := workFunc()
		if quit {
			return
		}
	}
}

func (c *Controller) syncItem(key string) error {
	startTime := time.Now()
	defer func() {
		log.Info("Finished syncing rolebinding", log.String("rolebinding", key), log.Duration("processTime", time.Since(startTime)))
	}()
	ns, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		return err
	}

	rb, err := c.rolebindingLister.RoleBindings(ns).Get(name)
	if err != nil {
		if errors.IsNotFound(err) {
			log.Info("RoleBinding has been deleted. Attempting to cleanup resources",
				log.String("namespace", ns),
				log.String("name", name))
			return nil
		}
		log.Warn("Unable to retrieve rolebinding from store",
			log.String("namespace", ns),
			log.String("name", name), log.Err(err))
		return err
	}
	rb = rb.DeepCopy()
	switch rb.Status.Phase {
	case apiauthzv1.BindingActive:
		return c.handleActive(context.Background(), rb)
	case apiauthzv1.BindingTerminating:
		return c.rolebindingDeleter.Delete(context.Background(), rb)
	default:
		return fmt.Errorf("unknown rolebinding '%s/%s' phase '%s'", rb.Namespace, rb.Name, rb.Status.Phase)
	}
}

func (c *Controller) handleActive(ctx context.Context, rb *apiauthzv1.RoleBinding) error {
	// 构造Policy和ClusterPolicyBinding资源并提交
	roleNs, roleName, _ := cache.SplitMetaNamespaceKey(rb.Spec.RoleName)
	role, err := c.roleLister.Roles(roleNs).Get(roleName)
	if err != nil {
		log.Warnf("Unable get role '%s/%s', err: '%v'", roleNs, roleName, err)
		return err
	}

	// 将Role关联的多个Policy合并为一个特殊的Policy
	var policyRules []rbacv1.PolicyRule
	for _, policy := range role.Policies {
		policyNamespace, policyName, _ := cache.SplitMetaNamespaceKey(policy)
		pol, err := c.policyLister.Policies(policyNamespace).Get(policyName)
		if err != nil {
			if errors.IsNotFound(err) {
				log.Warnf("Policy '%s/%s' is not exist", policyNamespace, policyName)
				continue
			}
			log.Warnf("Unable get policy '%s/%s', err: '%v'", policyNamespace, policyName, err)
			return err
		}
		policyRules = append(policyRules, pol.Rules...)
	}
	combinedPolicyName := fmt.Sprintf("role-%s", rb.Name)
	combinedPolicy := &apiauthzv1.Policy{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: rb.Namespace,
			Name:      combinedPolicyName,
		},
		Scope: apiauthzv1.ClusterScope,
		Rules: policyRules,
	}
	_, err = c.client.AuthzV1().Policies(rb.Namespace).Get(ctx, combinedPolicyName, metav1.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			if _, err = c.client.AuthzV1().Policies(rb.Namespace).Create(ctx, combinedPolicy, metav1.CreateOptions{}); err != nil {
				log.Warnf("Failed to crete policy '%s/%s', err: '%v'", rb.Namespace, combinedPolicyName, err)
				return err
			}
		} else {
			log.Warnf("Failed to get policy '%s/%s', err: '%v'", rb.Namespace, combinedPolicyName, err)
			return err
		}
	} else {
		if _, err = c.client.AuthzV1().Policies(rb.Namespace).Update(ctx, combinedPolicy, metav1.UpdateOptions{}); err != nil {
			log.Warnf("Failed to update policy '%s/%s', err: '%v'", rb.Namespace, combinedPolicyName, err)
			return err
		}
	}

	// 创建ClusterPolicyBinding
	cpb := &apiauthzv1.ClusterPolicyBinding{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: rb.Namespace,
			Name:      combinedPolicyName,
		},
		Spec: apiauthzv1.ClusterPolicyBindingSpec{
			PolicyName: fmt.Sprintf("%s/%s", rb.Namespace, combinedPolicyName),
			UserName:   rb.Spec.UserName,
			Clusters:   rb.Spec.Clusters,
		},
	}
	provider, err := authzprovider.GetProvider(rb.Annotations)
	if err != nil {
		log.Warn("Unable to retrieve provider",
			log.String("namespace", rb.Namespace),
			log.String("name", rb.Name), log.Err(err))
		return err
	}
	if err = provider.RenderClusterPolicyBinding(ctx, cpb); err != nil {
		log.Warnf("Unable render clusterpolicybinding, err '%v'", err)
		return err
	}
	if _, err = c.client.AuthzV1().ClusterPolicyBindings(rb.Namespace).Get(ctx, cpb.Name, metav1.GetOptions{}); err != nil {
		if errors.IsNotFound(err) {
			if _, err = c.client.AuthzV1().ClusterPolicyBindings(rb.Namespace).Create(ctx, cpb, metav1.CreateOptions{}); err != nil {
				log.Warnf("Failed to create clusterpolicybinding '%s/%s', err: '%v'", rb.Namespace, cpb.Name, err)
				return err
			}
		} else {
			log.Warnf("Failed to get clusterpolicybinding '%s/%s', err: '%v'", rb.Namespace, cpb.Name, err)
			return err
		}
	} else {
		if _, err = c.client.AuthzV1().ClusterPolicyBindings(rb.Namespace).Update(ctx, cpb, metav1.UpdateOptions{}); err != nil {
			log.Warnf("Failed to update clusterpolicybinding '%s/%s', err: '%v'", rb.Namespace, cpb.Name, err)
			return err
		}
	}
	return nil
}
