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

package authz

import (
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:skipVerbs=deleteCollection
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RoleTemplate is a rbac template in TKE.
type RoleTemplate struct {
	metav1.TypeMeta
	// +optional
	metav1.ObjectMeta
	// +optional
	Spec RoleTemplateSpec
}

type Scope string

const (
	ClusterScope  Scope = "Cluster"
	BusinessScope Scope = "Business"
)

type RoleTemplateSpec struct {
	TenantID    string
	DisplayName string
	// +optional
	Description string
	// +optional
	Scope Scope
	Rules []rbacv1.PolicyRule
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RoleTemplateList is the whole list of all rbac templates.
type RoleTemplateList struct {
	metav1.TypeMeta
	// +optional
	metav1.ListMeta
	// List of bootstraps
	Items []RoleTemplate
}
