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

package storage

import (
	"context"
	metainternal "k8s.io/apimachinery/pkg/apis/meta/internalversion"
	"k8s.io/apimachinery/pkg/runtime"
	genericregistry "k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/registry/generic/registry"
	"k8s.io/apiserver/pkg/registry/rest"
	"tkestack.io/tke/api/authz"
	apiserverutil "tkestack.io/tke/pkg/apiserver/util"
	"tkestack.io/tke/pkg/authz/registry/rolebinding"
	"tkestack.io/tke/pkg/util/log"
)

// Storage includes storage for configmap and all sub resources.
type Storage struct {
	RoleBinding *REST
}

// NewStorage returns a Storage object that will work against configmap.
func NewStorage(optsGetter genericregistry.RESTOptionsGetter) *Storage {
	strategy := rolebinding.NewStrategy()
	store := &registry.Store{
		NewFunc:                  func() runtime.Object { return &authz.RoleBinding{} },
		NewListFunc:              func() runtime.Object { return &authz.RoleBindingList{} },
		DefaultQualifiedResource: authz.Resource("rolebindings"),

		CreateStrategy: strategy,
		UpdateStrategy: strategy,
		DeleteStrategy: strategy,
	}
	store.TableConvertor = rest.NewDefaultTableConvertor(store.DefaultQualifiedResource)
	options := &genericregistry.StoreOptions{
		RESTOptions: optsGetter,
	}

	if err := store.CompleteWithOptions(options); err != nil {
		log.Panic("Failed to create configmap etcd rest storage", log.Err(err))
	}

	return &Storage{
		RoleBinding: &REST{store},
	}
}

// REST implements a RESTStorage for configmap against etcd.
type REST struct {
	*registry.Store
}

var _ rest.ShortNamesProvider = &REST{}

// ShortNames implements the ShortNamesProvider interface. Returns a list of short names for a resource.
func (r *REST) ShortNames() []string {
	return []string{"rb"}
}

// List selects resources in the storage which match to the selector. 'options' can be nil.
func (r *REST) List(ctx context.Context, options *metainternal.ListOptions) (runtime.Object, error) {
	wrappedOptions := apiserverutil.PredicateListOptions(ctx, options)
	return r.Store.List(ctx, wrappedOptions)
}

// StatusREST implements the GenericREST endpoint for changing the status of a rolebinding request.
type StatusREST struct {
	store *registry.Store
}

// New returns an empty object that can be used with Create and Update after request data has been put into it.
func (r *StatusREST) New() runtime.Object {
	return r.store.New()
}
