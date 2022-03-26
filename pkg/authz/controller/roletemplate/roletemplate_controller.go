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

package roletemplate

import (
	"context"
	v1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
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
	controllerName = "roletemplate-controller"
)

type Controller struct {
	client         clientset.Interface
	platformClient platformversionedclient.PlatformV1Interface
	queue          workqueue.RateLimitingInterface
	lister         authzv1.RoleTemplateLister
	listerSynced   cache.InformerSynced
	stopCh         <-chan struct{}
}

// NewController creates a new Controller object.
func NewController(
	client clientset.Interface,
	platformClient platformversionedclient.PlatformV1Interface,
	roletemplateInformer authzv1informer.RoleTemplateInformer,
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

	roletemplateInformer.Informer().AddEventHandlerWithResyncPeriod(
		cache.FilteringResourceEventHandler{
			Handler: cache.ResourceEventHandlerFuncs{
				AddFunc: controller.enqueue,
				UpdateFunc: func(oldObj, newObj interface{}) {
					old, ok1 := oldObj.(*apiauthzv1.RoleTemplate)
					cur, ok2 := newObj.(*apiauthzv1.RoleTemplate)
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
	controller.lister = roletemplateInformer.Lister()
	controller.listerSynced = roletemplateInformer.Informer().HasSynced
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

func (c *Controller) needsUpdate(old *apiauthzv1.RoleTemplate, new *apiauthzv1.RoleTemplate) bool {
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

	if ok := cache.WaitForCacheSync(stopCh, c.listerSynced); !ok {
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
		log.Info("Finished syncing roletemplate", log.String("roletemplate", key), log.Duration("processTime", time.Since(startTime)))
	}()
	clusterName, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		return err
	}
	roletemplate, err := c.lister.RoleTemplates(clusterName).Get(name)
	switch {
	case errors.IsNotFound(err):
		log.Info("App has been deleted. Attempting to cleanup resources",
			log.String("namespace", clusterName),
			log.String("name", name))
		return nil
	case err != nil:
		log.Warn("Unable to retrieve app from store",
			log.String("namespace", clusterName),
			log.String("name", name), log.Err(err))
		return err
	default:
		provider, err := authzprovider.GetProvider(&roletemplate.ObjectMeta)
		if err != nil {
			log.Warn("Unable to retrieve provider",
				log.String("namespace", clusterName),
				log.String("name", name), log.Err(err))
			return err
		}
		err = provider.ReconcileRoleTemplate(roletemplate, c.platformClient)
		if err != nil {
			return err
		}
		// 更新Status
		_, err = c.client.AuthzV1().RoleTemplates(clusterName).UpdateStatus(context.Background(), roletemplate, metav1.UpdateOptions{})
		return err

		// ========================= TODO remove ===========================
		cluster, err := c.platformClient.Clusters().Get(context.Background(), clusterName, metav1.GetOptions{})
		if err != nil {
			log.Warnf("Unable to retrieve cluster '%s'", clusterName)
			return err
		}
		v1Cluster, err := clusterprovider.GetV1Cluster(context.Background(), c.platformClient, cluster, clusterprovider.AdminUsername)
		if err != nil {
			log.Warnf("Unable to retrieve cluster '%s'", clusterName)
			return err
		}
		config, err := v1Cluster.RESTConfig()
		if err != nil {
			log.Warnf("Unable to retrieve cluster '%s'", clusterName)
			return err
		}
		client, err := kubernetes.NewForConfig(config)
		if err != nil {
			log.Warnf("Unable to retrieve cluster '%s'", clusterName)
			return err
		}
		expected := convertClusterRole(roletemplate)
		_, err = client.RbacV1().ClusterRoles().Get(context.Background(), roletemplate.Name, metav1.GetOptions{})
		if err != nil {
			if errors.IsNotFound(err) {
				_, err = client.RbacV1().ClusterRoles().Create(context.Background(), expected, metav1.CreateOptions{})
				return err
			}
			return err
		}
		_, err = client.RbacV1().ClusterRoles().Update(context.Background(), expected, metav1.UpdateOptions{})
		return err
		// ========================= TODO remove ===========================
	}
}

func convertClusterRole(template *apiauthzv1.RoleTemplate) *v1.ClusterRole {
	return &v1.ClusterRole{
		ObjectMeta: metav1.ObjectMeta{
			Name: template.Name,
		},
		Rules: template.Spec.Rules,
	}
}
