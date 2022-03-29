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

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1 "tkestack.io/tke/api/authz/v1"
)

// ClusterRoleTemplateBindingLister helps list ClusterRoleTemplateBindings.
// All objects returned here must be treated as read-only.
type ClusterRoleTemplateBindingLister interface {
	// List lists all ClusterRoleTemplateBindings in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ClusterRoleTemplateBinding, err error)
	// ClusterRoleTemplateBindings returns an object that can list and get ClusterRoleTemplateBindings.
	ClusterRoleTemplateBindings(namespace string) ClusterRoleTemplateBindingNamespaceLister
	ClusterRoleTemplateBindingListerExpansion
}

// clusterRoleTemplateBindingLister implements the ClusterRoleTemplateBindingLister interface.
type clusterRoleTemplateBindingLister struct {
	indexer cache.Indexer
}

// NewClusterRoleTemplateBindingLister returns a new ClusterRoleTemplateBindingLister.
func NewClusterRoleTemplateBindingLister(indexer cache.Indexer) ClusterRoleTemplateBindingLister {
	return &clusterRoleTemplateBindingLister{indexer: indexer}
}

// List lists all ClusterRoleTemplateBindings in the indexer.
func (s *clusterRoleTemplateBindingLister) List(selector labels.Selector) (ret []*v1.ClusterRoleTemplateBinding, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ClusterRoleTemplateBinding))
	})
	return ret, err
}

// ClusterRoleTemplateBindings returns an object that can list and get ClusterRoleTemplateBindings.
func (s *clusterRoleTemplateBindingLister) ClusterRoleTemplateBindings(namespace string) ClusterRoleTemplateBindingNamespaceLister {
	return clusterRoleTemplateBindingNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ClusterRoleTemplateBindingNamespaceLister helps list and get ClusterRoleTemplateBindings.
// All objects returned here must be treated as read-only.
type ClusterRoleTemplateBindingNamespaceLister interface {
	// List lists all ClusterRoleTemplateBindings in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ClusterRoleTemplateBinding, err error)
	// Get retrieves the ClusterRoleTemplateBinding from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.ClusterRoleTemplateBinding, error)
	ClusterRoleTemplateBindingNamespaceListerExpansion
}

// clusterRoleTemplateBindingNamespaceLister implements the ClusterRoleTemplateBindingNamespaceLister
// interface.
type clusterRoleTemplateBindingNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ClusterRoleTemplateBindings in the indexer for a given namespace.
func (s clusterRoleTemplateBindingNamespaceLister) List(selector labels.Selector) (ret []*v1.ClusterRoleTemplateBinding, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ClusterRoleTemplateBinding))
	})
	return ret, err
}

// Get retrieves the ClusterRoleTemplateBinding from the indexer for a given namespace and name.
func (s clusterRoleTemplateBindingNamespaceLister) Get(name string) (*v1.ClusterRoleTemplateBinding, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("clusterroletemplatebinding"), name)
	}
	return obj.(*v1.ClusterRoleTemplateBinding), nil
}