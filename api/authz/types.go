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

// Role is a collection with multiple policies.
type Role struct {
	metav1.TypeMeta
	metav1.ObjectMeta

	DisplayName string
	TenantID    string
	// Username is Creator
	Username    string
	Description string
	// policyNamespace/policyName
	Policies []string
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RoleList is the whole list of policy.
type RoleList struct {
	metav1.TypeMeta
	metav1.ListMeta
	// List of rules.
	Items []Role
}

// +genclient
// +genclient:skipVerbs=deleteCollection
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type RoleBinding struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec   RoleBindingSpec
	Status RoleBindingStatus
}

type RoleBindingSpec struct {
	// +optional
	UserName string
	// +optional
	GroupName string
	RoleName  string
	Clusters  []string
}

type RoleBindingStatus struct {
	// +optional
	Phase BindingPhase
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type RoleBindingList struct {
	metav1.TypeMeta
	metav1.ListMeta
	// List of rules.
	Items []RoleBinding
}

// +genclient
// +genclient:skipVerbs=deleteCollection
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Policy struct {
	metav1.TypeMeta
	// +optional
	metav1.ObjectMeta

	// +optional
	TenantID    string
	DisplayName string
	// +optional
	Description string
	// +optional
	Scope Scope
	Rules []rbacv1.PolicyRule
}

type Scope string

const (
	PlatformScope Scope = "Platform"
	ClusterScope  Scope = "Cluster"
	BusinessScope Scope = "Business"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PolicyList is the whole list of all policies.
type PolicyList struct {
	metav1.TypeMeta
	// +optional
	metav1.ListMeta
	// List of policies
	Items []Policy
}

// +genclient
// +genclient:skipVerbs=deleteCollection
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ClusterPolicyBinding struct {
	metav1.TypeMeta
	// +optional
	metav1.ObjectMeta
	Spec   ClusterPolicyBindingSpec
	Status ClusterPolicyBindingStatus
}

type ClusterPolicyBindingSpec struct {
	// +optional
	UserName string
	// +optional
	GroupName string
	// PolicyNamespace/PolicyName
	PolicyName string
	Clusters   []string
}

type BindingPhase string

const (
	BindingActive      BindingPhase = "Active"
	BindingTerminating BindingPhase = "Terminating"
)

type ClusterPolicyBindingStatus struct {
	// +optional
	Phase BindingPhase
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterPolicyBindingList is a resource containing a list of ClusterPolicyBinding objects.
type ClusterPolicyBindingList struct {
	metav1.TypeMeta
	// +optional
	metav1.ListMeta
	// Items is the list of ConfigMaps.
	Items []ClusterPolicyBinding
}

type FinalizerName string

const (
	PolicyFinalize               FinalizerName = "policy"
	ClusterPolicyBindingFinalize FinalizerName = "clusterpolicybinding"
	RoleBindingFinalize          FinalizerName = "rolebinding"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ConfigMap holds configuration data for tke to consume.
type ConfigMap struct {
	metav1.TypeMeta
	// +optional
	metav1.ObjectMeta

	// Data contains the configuration data.
	// Each key must consist of alphanumeric characters, '-', '_' or '.'.
	// Values with non-UTF-8 byte sequences must use the BinaryData field.
	// The keys stored in Data must not overlap with the keys in
	// the BinaryData field, this is enforced during validation process.
	// +optional
	Data map[string]string

	// BinaryData contains the binary data.
	// Each key must consist of alphanumeric characters, '-', '_' or '.'.
	// BinaryData can contain byte sequences that are not in the UTF-8 range.
	// The keys stored in BinaryData must not overlap with the ones in
	// the Data field, this is enforced during validation process.
	// +optional
	BinaryData map[string][]byte
}

// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ConfigMapList is a resource containing a list of ConfigMap objects.
type ConfigMapList struct {
	metav1.TypeMeta
	// +optional
	metav1.ListMeta
	// Items is the list of ConfigMaps.
	Items []ConfigMap
}
