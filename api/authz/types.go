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
	// +optional
	Status RoleTemplateStatus
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

type RoleTemplatePhase string

const (
	Installing RoleTemplatePhase = "Installing"
	Succeeded  RoleTemplatePhase = "Succeeded"
	Failed     RoleTemplatePhase = "Failed"
)

type RoleTemplateStatus struct {
	// Phase the release is in, one of ('Installing', 'Succeeded', 'Failed')
	// +optional
	Phase RoleTemplatePhase
	// The last time the condition transitioned from one status to another.
	// +optional
	LastTransitionTime metav1.Time
	// The reason for the condition's last transition.
	// +optional
	Reason string
	// A human readable message indicating details about the transition.
	// +optional
	Message string
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

// +genclient
// +genclient:skipVerbs=deleteCollection
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ClusterRoleTemplateBinding struct {
	metav1.TypeMeta
	// +optional
	metav1.ObjectMeta
	Spec ClusterRoleTemplateBindingSpec
	Status ClusterRoleTemplateBindingStatus
}

type ClusterRoleTemplateBindingSpec struct {
	// +optional
	UserName         string
	// +optional
	GroupName        string
	RoleTemplateName string
	Clusters         []string
}

type ClusterRoleTemplateBindingStatus struct {
	// Phase the release is in, one of ('Installing', 'Succeeded', 'Failed')
	// +optional
	Phase RoleTemplatePhase
	// The last time the condition transitioned from one status to another.
	// +optional
	LastTransitionTime metav1.Time
	// The reason for the condition's last transition.
	// +optional
	Reason string
	// A human readable message indicating details about the transition.
	// +optional
	Message string
	// +optional
	Clusters []ClusterRoleTemplateBindingStatusItem
}

type ClusterRoleTemplateBindingStatusItem struct {
	// Phase the release is in, one of ('Installing', 'Succeeded', 'Failed')
	// +optional
	Phase RoleTemplatePhase
	// The last time the condition transitioned from one status to another.
	// +optional
	LastTransitionTime metav1.Time
	// The reason for the condition's last transition.
	// +optional
	Reason string
	// A human readable message indicating details about the transition.
	// +optional
	Message string
}

// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterRoleTemplateBindingList is a resource containing a list of ClusterRoleTemplateBinding objects.
type ClusterRoleTemplateBindingList struct {
	metav1.TypeMeta
	// +optional
	metav1.ListMeta
	// Items is the list of ConfigMaps.
	Items []ConfigMap
}

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
