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

package clusterpolicybinding

import (
	"context"
	"fmt"
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/api/errors"
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
	"tkestack.io/tke/pkg/authz/controller/clusterpolicybinding/deletion"
	authzprovider "tkestack.io/tke/pkg/authz/provider"
	controllerutil "tkestack.io/tke/pkg/controller"
	clusterprovider "tkestack.io/tke/pkg/platform/provider/cluster"
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
	controllerName = "clusterpolicybinding-controller"
)

type Controller struct {
	client                      clientset.Interface
	platformClient              platformversionedclient.PlatformV1Interface
	queue                       workqueue.RateLimitingInterface
	policyLister                authzv1.PolicyLister
	clusterPolicyBindingLister  authzv1.ClusterPolicyBindingLister
	policySynced                cache.InformerSynced
	clusterPolicyBindingSynced  cache.InformerSynced
	clusterPolicyBindingDeleter deletion.ClusterPolicyBindingDeleter
	stopCh                      <-chan struct{}
}

// NewController creates a new Controller object.
func NewController(
	client clientset.Interface,
	platformClient platformversionedclient.PlatformV1Interface,
	policyInformer authzv1informer.PolicyInformer,
	clusterPolicyBindingInformer authzv1informer.ClusterPolicyBindingInformer,
	resyncPeriod time.Duration,
) *Controller {
	// create the controller so we can inject the enqueue function
	controller := &Controller{
		client:                      client,
		platformClient:              platformClient,
		queue:                       workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), controllerName),
		clusterPolicyBindingDeleter: deletion.New(client),
	}
	if client != nil &&
		client.AuthzV1().RESTClient() != nil &&
		!reflect.ValueOf(client.AuthzV1().RESTClient()).IsNil() &&
		client.AuthzV1().RESTClient().GetRateLimiter() != nil {
		_ = metrics.RegisterMetricAndTrackRateLimiterUsage(controllerName, client.AuthzV1().RESTClient().GetRateLimiter())
	}

	policyInformer.Informer().AddEventHandlerWithResyncPeriod(
		cache.FilteringResourceEventHandler{
			Handler: cache.ResourceEventHandlerFuncs{
				AddFunc: controller.addPolicy,
				UpdateFunc: func(oldObj, newObj interface{}) {
					old, ok1 := oldObj.(*apiauthzv1.Policy)
					cur, ok2 := newObj.(*apiauthzv1.Policy)
					if ok1 && ok2 && controller.needsAddPolicy(old, cur) {
						controller.addPolicy(newObj)
					}
				},
			},
			FilterFunc: func(obj interface{}) bool {
				return true
			},
		}, resyncPeriod,
	)

	clusterPolicyBindingInformer.Informer().AddEventHandlerWithResyncPeriod(
		cache.FilteringResourceEventHandler{
			Handler: cache.ResourceEventHandlerFuncs{
				AddFunc: controller.enqueue,
				UpdateFunc: func(oldObj, newObj interface{}) {
					old, ok1 := oldObj.(*apiauthzv1.ClusterPolicyBinding)
					cur, ok2 := newObj.(*apiauthzv1.ClusterPolicyBinding)
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
	controller.clusterPolicyBindingLister = clusterPolicyBindingInformer.Lister()
	controller.clusterPolicyBindingSynced = clusterPolicyBindingInformer.Informer().HasSynced
	controller.policyLister = policyInformer.Lister()
	controller.policySynced = policyInformer.Informer().HasSynced
	return controller
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

func (c *Controller) addPolicy(obj interface{}) {
	policy := obj.(*apiauthzv1.Policy)
	var (
		cpbs []*apiauthzv1.ClusterPolicyBinding
		err  error
	)
	selector := labels.SelectorFromSet(map[string]string{
		constant.PolicyNamespace: policy.Namespace,
		constant.PolicyName:      policy.Name,
	})
	if policy.Namespace == "" {
		// 如果是默认策略，允许绑定到任何ns
		cpbs, err = c.clusterPolicyBindingLister.List(selector)
	} else {
		cpbs, err = c.clusterPolicyBindingLister.ClusterPolicyBindings(policy.Namespace).List(selector)
	}
	if err != nil {
		log.Warnf("Failed to list clusterpolicybindings by policy %s/%s", policy.Namespace, policy.Name)
		return
	}
	for _, cpb := range cpbs {
		c.enqueue(cpb)
	}
}

func (c *Controller) enqueue(obj interface{}) {
	key, err := controllerutil.KeyFunc(obj)
	if err != nil {
		log.Error("Couldn't get key for object", log.Any("object", obj), log.Err(err))
		return
	}
	c.queue.AddAfter(key, appDeletionGracePeriod)
}

func (c *Controller) needsUpdate(old *apiauthzv1.ClusterPolicyBinding, new *apiauthzv1.ClusterPolicyBinding) bool {
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
	log.Info("Starting app controller")
	defer log.Info("Shutting down app controller")

	if ok := cache.WaitForCacheSync(stopCh, c.clusterPolicyBindingSynced, c.policySynced); !ok {
		log.Error("Failed to wait for app caches to sync")
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
		log.Info("Finished syncing clusterpolicybinding", log.String("clusterpolicybinding", key), log.Duration("processTime", time.Since(startTime)))
	}()
	ns, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		return err
	}

	cpb, err := c.clusterPolicyBindingLister.ClusterPolicyBindings(ns).Get(name)
	if err != nil {
		if errors.IsNotFound(err) {
			log.Info("ClusterPolicyBinding has been deleted. Attempting to cleanup resources",
				log.String("namespace", ns),
				log.String("name", name))
			return nil
		}
		log.Warn("Unable to retrieve clusterpolicybinding from store",
			log.String("namespace", ns),
			log.String("name", name), log.Err(err))
		return err
	}
	cpb = cpb.DeepCopy()
	provider, err := authzprovider.GetProvider(cpb.Annotations)
	if err != nil {
		log.Warn("Unable to retrieve provider",
			log.String("namespace", ns),
			log.String("name", name), log.Err(err))
		return err
	}
	ctx := provider.InitContext(cpb)

	switch cpb.Status.Phase {
	case apiauthzv1.BindingActive:
		return c.handleActive(ctx, cpb, provider)
	case apiauthzv1.BindingTerminating:
		return c.clusterPolicyBindingDeleter.Delete(ctx, cpb, provider)
	default:
		return fmt.Errorf("unknown clusterpolicybinding phase '%s'", cpb.Status.Phase)
	}
}

func (c *Controller) handleActive(ctx context.Context, cpb *apiauthzv1.ClusterPolicyBinding, provider authzprovider.Provider) error {
	policyNs, policyName, err := cache.SplitMetaNamespaceKey(cpb.Spec.PolicyName)
	if err != nil {
		log.Warnf("failed to parse clusterpolicybinding namespace/name '%s'", cpb.Spec.PolicyName)
		return err
	}
	policy, err := c.policyLister.Policies(policyNs).Get(policyName)
	if err != nil {
		log.Warn("Unable to retrieve clusterpolicybinding from store",
			log.String("namespace", policyNs),
			log.String("name", policyName), log.Err(err))
		return err
	}
	// 获取user在各个cluster内的subject
	clusterSubjects := map[string]*rbacv1.Subject{}
	for _, cls := range cpb.Spec.Clusters {
		cluster, err := clusterprovider.GetV1ClusterByName(ctx, c.platformClient, cls, cpb.Spec.UserName)
		if err != nil {
			log.Warnf("GetV1ClusterByName failed, cluster: '%s', user: '%s', err: '%#v'", cls, cpb.Spec.UserName, err)
			return err
		}
		subject, err := provider.GetSubject(ctx, cpb.Spec.UserName, cluster)
		if err != nil {
			log.Warnf("GetSubject failed, cluster: '%s',  user: '%s', err: '%#v'", cls, cpb.Spec.UserName, err)
			return err
		}
		clusterSubjects[cls] = subject
	}
	// 执行权限分发
	err = provider.DispatchPolicy(ctx, c.platformClient, policy, cpb, clusterSubjects)
	if err != nil {
		log.Warnf("DispatchPolicy failed, clusterpolicybinding: '%s', err: '%#v'", cpb.Name, err)
		return err
	}
	return nil
}
