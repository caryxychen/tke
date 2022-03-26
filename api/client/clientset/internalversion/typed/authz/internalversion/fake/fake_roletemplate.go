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

// FakeRoleTemplates implements RoleTemplateInterface
type FakeRoleTemplates struct {
	Fake *FakeAuthz
	ns   string
}

var roletemplatesResource = schema.GroupVersionResource{Group: "authz.tkestack.io", Version: "", Resource: "roletemplates"}

var roletemplatesKind = schema.GroupVersionKind{Group: "authz.tkestack.io", Version: "", Kind: "RoleTemplate"}

// Get takes name of the roleTemplate, and returns the corresponding roleTemplate object, and an error if there is any.
func (c *FakeRoleTemplates) Get(ctx context.Context, name string, options v1.GetOptions) (result *authz.RoleTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(roletemplatesResource, c.ns, name), &authz.RoleTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*authz.RoleTemplate), err
}

// List takes label and field selectors, and returns the list of RoleTemplates that match those selectors.
func (c *FakeRoleTemplates) List(ctx context.Context, opts v1.ListOptions) (result *authz.RoleTemplateList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(roletemplatesResource, roletemplatesKind, c.ns, opts), &authz.RoleTemplateList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &authz.RoleTemplateList{ListMeta: obj.(*authz.RoleTemplateList).ListMeta}
	for _, item := range obj.(*authz.RoleTemplateList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested roleTemplates.
func (c *FakeRoleTemplates) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(roletemplatesResource, c.ns, opts))

}

// Create takes the representation of a roleTemplate and creates it.  Returns the server's representation of the roleTemplate, and an error, if there is any.
func (c *FakeRoleTemplates) Create(ctx context.Context, roleTemplate *authz.RoleTemplate, opts v1.CreateOptions) (result *authz.RoleTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(roletemplatesResource, c.ns, roleTemplate), &authz.RoleTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*authz.RoleTemplate), err
}

// Update takes the representation of a roleTemplate and updates it. Returns the server's representation of the roleTemplate, and an error, if there is any.
func (c *FakeRoleTemplates) Update(ctx context.Context, roleTemplate *authz.RoleTemplate, opts v1.UpdateOptions) (result *authz.RoleTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(roletemplatesResource, c.ns, roleTemplate), &authz.RoleTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*authz.RoleTemplate), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRoleTemplates) UpdateStatus(ctx context.Context, roleTemplate *authz.RoleTemplate, opts v1.UpdateOptions) (*authz.RoleTemplate, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(roletemplatesResource, "status", c.ns, roleTemplate), &authz.RoleTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*authz.RoleTemplate), err
}

// Delete takes name of the roleTemplate and deletes it. Returns an error if one occurs.
func (c *FakeRoleTemplates) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(roletemplatesResource, c.ns, name), &authz.RoleTemplate{})

	return err
}

// Patch applies the patch and returns the patched roleTemplate.
func (c *FakeRoleTemplates) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *authz.RoleTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(roletemplatesResource, c.ns, name, pt, data, subresources...), &authz.RoleTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*authz.RoleTemplate), err
}
