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

package fake

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	authz "tkestack.io/tke/api/authz"
)

// FakePolicies implements PolicyInterface
type FakePolicies struct {
	Fake *FakeAuthz
	ns   string
}

var policiesResource = schema.GroupVersionResource{Group: "authz.tkestack.io", Version: "", Resource: "policies"}

var policiesKind = schema.GroupVersionKind{Group: "authz.tkestack.io", Version: "", Kind: "Policy"}

// Get takes name of the policy, and returns the corresponding policy object, and an error if there is any.
func (c *FakePolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *authz.Policy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(policiesResource, c.ns, name), &authz.Policy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*authz.Policy), err
}

// List takes label and field selectors, and returns the list of Policies that match those selectors.
func (c *FakePolicies) List(ctx context.Context, opts v1.ListOptions) (result *authz.PolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(policiesResource, policiesKind, c.ns, opts), &authz.PolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &authz.PolicyList{ListMeta: obj.(*authz.PolicyList).ListMeta}
	for _, item := range obj.(*authz.PolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested policies.
func (c *FakePolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(policiesResource, c.ns, opts))

}

// Create takes the representation of a policy and creates it.  Returns the server's representation of the policy, and an error, if there is any.
func (c *FakePolicies) Create(ctx context.Context, policy *authz.Policy, opts v1.CreateOptions) (result *authz.Policy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(policiesResource, c.ns, policy), &authz.Policy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*authz.Policy), err
}

// Update takes the representation of a policy and updates it. Returns the server's representation of the policy, and an error, if there is any.
func (c *FakePolicies) Update(ctx context.Context, policy *authz.Policy, opts v1.UpdateOptions) (result *authz.Policy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(policiesResource, c.ns, policy), &authz.Policy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*authz.Policy), err
}

// Delete takes name of the policy and deletes it. Returns an error if one occurs.
func (c *FakePolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(policiesResource, c.ns, name), &authz.Policy{})

	return err
}

// Patch applies the patch and returns the patched policy.
func (c *FakePolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *authz.Policy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(policiesResource, c.ns, name, pt, data, subresources...), &authz.Policy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*authz.Policy), err
}
