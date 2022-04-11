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
	client            clientset.Interface
	platformClient    platformversionedclient.PlatformV1Interface
	queue             workqueue.RateLimitingInterface
	roleLister        authzv1.RoleLister
	rolebindingLister authzv1.RoleBindingLister
	policyLister      authzv1.PolicyLister
	roleSynced        cache.InformerSynced
	rolebindingSynced cache.InformerSynced
	policySynced      cache.InformerSynced
	stopCh            <-chan struct{}
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
		client:         client,
		platformClient: platformClient,
		queue:          workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), controllerName),
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
			// TODO
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

	// 构造 Policy和ClusterPolicyBinding资源并提交
	roleNs, roleName, err := cache.SplitMetaNamespaceKey(rb.Spec.RoleName)
	if err != nil {
		return err
	}
	role, err := c.roleLister.Roles(roleNs).Get(roleName)
	if err != nil {
		return err
	}
	var policyRules []rbacv1.PolicyRule
	for _, policy := range role.Policies {
		policyNamespace, policyName, err := cache.SplitMetaNamespaceKey(policy)
		if err != nil {
			return err
		}
		pol, err := c.policyLister.Policies(policyNamespace).Get(policyName)
		if err != nil {
			if errors.IsNotFound(err) {
				continue
			}
			return err
		}
		policyRules = append(policyRules, pol.Rules...)
	}
	combinedPolicyName := fmt.Sprintf("role-%s", name)
	combinedPolicy := &apiauthzv1.Policy{
		ObjectMeta: metav1.ObjectMeta{
			Name:      combinedPolicyName,
			Namespace: ns,
		},
		Scope: apiauthzv1.ClusterScope,
		Rules: policyRules,
	}
	_, err = c.client.AuthzV1().Policies(ns).Get(context.Background(), combinedPolicyName, metav1.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			if _, err = c.client.AuthzV1().Policies(ns).Create(context.Background(), combinedPolicy, metav1.CreateOptions{}); err != nil {
				return err
			}
		} else {
			return err
		}
	} else {
		if _, err = c.client.AuthzV1().Policies(ns).Update(context.Background(), combinedPolicy, metav1.UpdateOptions{}); err != nil {
			return err
		}
	}

	provider, err := authzprovider.GetProvider(rb.Annotations)
	if err != nil {
		log.Warn("Unable to retrieve provider",
			log.String("namespace", ns),
			log.String("name", name), log.Err(err))
		return err
	}

	// 创建ClusterPolicyBinding
	cpb := &apiauthzv1.ClusterPolicyBinding{
		ObjectMeta: metav1.ObjectMeta{
			Name:      combinedPolicyName,
			Namespace: ns,
		},
		Spec: apiauthzv1.ClusterPolicyBindingSpec{
			PolicyName: fmt.Sprintf("%s/%s", ns, combinedPolicyName),
			UserName:   rb.Spec.UserName,
			Clusters:   rb.Spec.Clusters,
		},
	}
	provider.RenderClusterPolicyBinding(context.Background(), cpb)
	if _, err = c.client.AuthzV1().ClusterPolicyBindings(ns).Get(context.Background(), cpb.Name, metav1.GetOptions{}); err != nil {
		if errors.IsNotFound(err) {
			if _, err = c.client.AuthzV1().ClusterPolicyBindings(ns).Create(context.Background(), cpb, metav1.CreateOptions{}); err != nil {
				return err
			}
		} else {
			return err
		}
	} else {
		if _, err = c.client.AuthzV1().ClusterPolicyBindings(ns).Update(context.Background(), cpb, metav1.UpdateOptions{}); err != nil {
			return err
		}
	}
	return err
}
