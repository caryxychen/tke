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
 * Unless required by policylicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

package policy

import (
	"context"
	"fmt"
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
	authzv1informer "tkestack.io/tke/api/client/informers/externalversions/authz/v1"
	authzv1 "tkestack.io/tke/api/client/listers/authz/v1"
	"tkestack.io/tke/pkg/authz/constant"
	controllerutil "tkestack.io/tke/pkg/controller"
	"tkestack.io/tke/pkg/util/log"
	"tkestack.io/tke/pkg/util/metrics"
)

const (
	// policyDeletionGracePeriod is the time period to wait before processing a received channel event.
	// This allows time for the following to occur:
	// * lifecycle admission plugins on HA apiservers to also observe a channel
	//   deletion and prevent new objects from being created in the terminating channel
	// * non-leader etcd servers to observe last-minute object creations in a channel
	//   so this controller's cleanup can actually clean up all objects
	policyDeletionGracePeriod = 5 * time.Second
)

const (
	controllerName = "policy-controller"
)

type Controller struct {
	client                     clientset.Interface
	queue                      workqueue.RateLimitingInterface
	policyLister               authzv1.PolicyLister
	clusterPolicyBindingLister authzv1.ClusterPolicyBindingLister
	policySynced               cache.InformerSynced
	clusterPolicyBindingSynced cache.InformerSynced
	stopCh                     <-chan struct{}
}

// NewController creates a new Controller object.
func NewController(
	client clientset.Interface,
	policyInformer authzv1informer.PolicyInformer,
	clusterPolicyBindingInformer authzv1informer.ClusterPolicyBindingInformer,
	resyncPeriod time.Duration,
) *Controller {
	// create the controller so we can inject the enqueue function
	controller := &Controller{
		client: client,
		queue:  workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), controllerName),
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
				AddFunc: controller.enqueue,
				UpdateFunc: func(oldObj, newObj interface{}) {
					controller.enqueue(newObj)
				},
				DeleteFunc: controller.enqueue,
			},
			FilterFunc: func(obj interface{}) bool {
				return true
			},
		}, resyncPeriod,
	)
	controller.clusterPolicyBindingLister = clusterPolicyBindingInformer.Lister()
	controller.clusterPolicyBindingSynced = clusterPolicyBindingInformer.Informer().HasSynced
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
	policy := obj.(*apiauthzv1.Policy)
	if policy.DeletionTimestamp != nil {
		c.queue.AddAfter(key, policyDeletionGracePeriod)
	}
	return
}

// Run will set up the event handlers for types we are interested in, as well
// as syncing informer caches and starting workers.
func (c *Controller) Run(workers int, stopCh <-chan struct{}) {
	defer runtime.HandleCrash()
	defer c.queue.ShutDown()

	// Start the informer factories to begin populating the informer caches
	log.Info("Starting policy controller")
	defer log.Info("Shutting down policy controller")

	if ok := cache.WaitForCacheSync(stopCh, c.clusterPolicyBindingSynced, c.policySynced); !ok {
		log.Error("Failed to wait for policy caches to sync")
		return
	}

	c.stopCh = stopCh
	for i := 0; i < workers; i++ {
		go wait.Until(c.worker, time.Second, stopCh)
	}

	<-stopCh
}

// worker processes the queue of policy objects.
// Each policy can be in the queue at most once.
// The system ensures that no two workers can process
// the same policy at the same time.
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

		// rather than wait for a full resync, re-add the policy to the queue to be processed
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
		log.Info("Finished syncing policy", log.String("policy", key), log.Duration("processTime", time.Since(startTime)))
	}()
	ns, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		return err
	}
	pol, err := c.policyLister.Policies(ns).Get(name)
	if err != nil {
		if errors.IsNotFound(err) {
			log.Info("Policy has been deleted. Attempting to cleanup resources",
				log.String("namespace", ns),
				log.String("name", name))
			return nil
		}
		log.Warn("Unable to retrieve policy from store",
			log.String("namespace", ns),
			log.String("name", name), log.Err(err))
		return err
	}

	selector := labels.SelectorFromSet(map[string]string{
		constant.PolicyNamespace: ns,
		constant.PolicyName:      name,
	})
	var cpbs []*apiauthzv1.ClusterPolicyBinding
	if ns == "" {
		// 如果是默认策略，允许绑定到任何ns
		cpbs, err = c.clusterPolicyBindingLister.List(selector)
	} else {
		cpbs, err = c.clusterPolicyBindingLister.ClusterPolicyBindings(ns).List(selector)
	}

	if len(cpbs) == 0 {
		policyFinalize := apiauthzv1.Policy{}
		policyFinalize.ObjectMeta = pol.ObjectMeta
		policyFinalize.Finalizers = []string{}
		if err := c.client.AuthzV1().RESTClient().Put().Resource("policies").
			Namespace(ns).
			Name(name).
			SubResource("finalize").
			Body(&policyFinalize).
			Do(context.Background()).
			Into(&policyFinalize); err != nil {
			log.Warnf("Unable to finalize policy '%s/%s', err: %v", ns, name, err)
			return err
		}
		return c.client.AuthzV1().Policies(ns).Delete(context.Background(), name, metav1.DeleteOptions{})
	} else {
		for _, cpb := range cpbs {
			if cpb.Status.Phase == apiauthzv1.BindingTerminating {
				continue
			}
			deleteOpt := metav1.DeletePropagationBackground
			if err = c.client.AuthzV1().ClusterPolicyBindings(cpb.Namespace).Delete(context.Background(), cpb.Name, metav1.DeleteOptions{PropagationPolicy: &deleteOpt}); err != nil {
				log.Warnf("Unable to delete clusterpolicybinding '%s/%s', err: '%v'", cpb.Namespace, cpb.Name, err)
				return err
			}
		}
		return fmt.Errorf("policy '%s/%s' will be retried later", ns, name)
	}
}
