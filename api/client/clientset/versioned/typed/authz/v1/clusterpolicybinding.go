/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2020 Tencent. All Rights Reserved.
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

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1 "tkestack.io/tke/api/authz/v1"
	scheme "tkestack.io/tke/api/client/clientset/versioned/scheme"
)

// ClusterPolicyBindingsGetter has a method to return a ClusterPolicyBindingInterface.
// A group's client should implement this interface.
type ClusterPolicyBindingsGetter interface {
	ClusterPolicyBindings(namespace string) ClusterPolicyBindingInterface
}

// ClusterPolicyBindingInterface has methods to work with ClusterPolicyBinding resources.
type ClusterPolicyBindingInterface interface {
	Create(ctx context.Context, clusterPolicyBinding *v1.ClusterPolicyBinding, opts metav1.CreateOptions) (*v1.ClusterPolicyBinding, error)
	Update(ctx context.Context, clusterPolicyBinding *v1.ClusterPolicyBinding, opts metav1.UpdateOptions) (*v1.ClusterPolicyBinding, error)
	UpdateStatus(ctx context.Context, clusterPolicyBinding *v1.ClusterPolicyBinding, opts metav1.UpdateOptions) (*v1.ClusterPolicyBinding, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.ClusterPolicyBinding, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.ClusterPolicyBindingList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ClusterPolicyBinding, err error)
	ClusterPolicyBindingExpansion
}

// clusterPolicyBindings implements ClusterPolicyBindingInterface
type clusterPolicyBindings struct {
	client rest.Interface
	ns     string
}

// newClusterPolicyBindings returns a ClusterPolicyBindings
func newClusterPolicyBindings(c *AuthzV1Client, namespace string) *clusterPolicyBindings {
	return &clusterPolicyBindings{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the clusterPolicyBinding, and returns the corresponding clusterPolicyBinding object, and an error if there is any.
func (c *clusterPolicyBindings) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.ClusterPolicyBinding, err error) {
	result = &v1.ClusterPolicyBinding{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("clusterpolicybindings").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ClusterPolicyBindings that match those selectors.
func (c *clusterPolicyBindings) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ClusterPolicyBindingList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.ClusterPolicyBindingList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("clusterpolicybindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested clusterPolicyBindings.
func (c *clusterPolicyBindings) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("clusterpolicybindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a clusterPolicyBinding and creates it.  Returns the server's representation of the clusterPolicyBinding, and an error, if there is any.
func (c *clusterPolicyBindings) Create(ctx context.Context, clusterPolicyBinding *v1.ClusterPolicyBinding, opts metav1.CreateOptions) (result *v1.ClusterPolicyBinding, err error) {
	result = &v1.ClusterPolicyBinding{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("clusterpolicybindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterPolicyBinding).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a clusterPolicyBinding and updates it. Returns the server's representation of the clusterPolicyBinding, and an error, if there is any.
func (c *clusterPolicyBindings) Update(ctx context.Context, clusterPolicyBinding *v1.ClusterPolicyBinding, opts metav1.UpdateOptions) (result *v1.ClusterPolicyBinding, err error) {
	result = &v1.ClusterPolicyBinding{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("clusterpolicybindings").
		Name(clusterPolicyBinding.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterPolicyBinding).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *clusterPolicyBindings) UpdateStatus(ctx context.Context, clusterPolicyBinding *v1.ClusterPolicyBinding, opts metav1.UpdateOptions) (result *v1.ClusterPolicyBinding, err error) {
	result = &v1.ClusterPolicyBinding{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("clusterpolicybindings").
		Name(clusterPolicyBinding.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterPolicyBinding).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the clusterPolicyBinding and deletes it. Returns an error if one occurs.
func (c *clusterPolicyBindings) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("clusterpolicybindings").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched clusterPolicyBinding.
func (c *clusterPolicyBindings) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ClusterPolicyBinding, err error) {
	result = &v1.ClusterPolicyBinding{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("clusterpolicybindings").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
