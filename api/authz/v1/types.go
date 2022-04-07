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

package v1

import (
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:skipVerbs=deleteCollection
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Role is a collection with multiple policies.
type Role struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	DisplayName string `json:"displayName" protobuf:"bytes,2,opt,name=displayName"`
	TenantID    string `json:"tenantID" protobuf:"bytes,3,opt,name=tenantID"`
	// Username is Creator
	Username    string   `json:"username" protobuf:"bytes,4,opt,name=username"`
	Description string   `json:"description" protobuf:"bytes,5,opt,name=description"`
	Policies    []string `json:"policies" protobuf:"bytes,6,rep,name=policies"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RoleList is the whole list of policy.
type RoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	// List of rules.
	Items []Role `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +genclient
// +genclient:skipVerbs=deleteCollection
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type RoleBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec              RoleBindingSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}

type RoleBindingSpec struct {
	// +optional
	UserName string `json:"userName" protobuf:"bytes,1,opt,name=userName"`
	// +optional
	GroupName string   `json:"groupName" protobuf:"bytes,2,opt,name=groupName"`
	RoleName  string   `json:"roleName" protobuf:"bytes,3,opt,name=roleName"`
	Clusters  []string `json:"clusters" protobuf:"bytes,4,rep,name=clusters"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type RoleBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	// List of rules.
	Items []RoleBindingList `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +genclient
// +genclient:skipVerbs=deleteCollection
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RoleTemplate is a rbac template in TKE.
type RoleTemplate struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// +optional
	Spec RoleTemplateSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`

	// +optional
	Status RoleTemplateStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type Scope string

const (
	ClusterScope  Scope = "Cluster"
	BusinessScope Scope = "Business"
)

type RoleTemplateSpec struct {
	TenantID    string `json:"tenantID" protobuf:"bytes,1,opt,name=tenantID"`
	DisplayName string `json:"displayName" protobuf:"bytes,2,opt,name=displayName"`
	// +optional
	Description string `json:"description" protobuf:"bytes,3,opt,name=description"`
	// +optional
	Scope Scope               `json:"scope" protobuf:"bytes,4,opt,name=scope"`
	Rules []rbacv1.PolicyRule `json:"rules" protobuf:"bytes,5,rep,name=rules"`
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
	Phase RoleTemplatePhase `json:"phase" protobuf:"bytes,1,opt,name=phase"`
	// The last time the condition transitioned from one status to another.
	// +optional
	LastTransitionTime metav1.Time `json:"lastTransitionTime" protobuf:"bytes,2,opt,name=lastTransitionTime"`
	// The reason for the condition's last transition.
	// +optional
	Reason string `json:"reason" protobuf:"bytes,3,opt,name=reason"`
	// A human readable message indicating details about the transition.
	// +optional
	Message string `json:"message" protobuf:"bytes,4,opt,name=message"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RoleTemplateList is the whole list of all rbac templates.
type RoleTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// List of bootstraps
	Items []RoleTemplate `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +genclient
// +genclient:skipVerbs=deleteCollection
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ClusterRoleTemplateBinding struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec              ClusterRoleTemplateBindingSpec   `json:"spec" protobuf:"bytes,2,opt,name=spec"`
	Status            ClusterRoleTemplateBindingStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type ClusterRoleTemplateBindingSpec struct {
	// +optional
	UserName string `json:"userName" protobuf:"bytes,1,opt,name=userName"`
	// +optional
	GroupName        string   `json:"groupName" protobuf:"bytes,2,opt,name=groupName"`
	RoleTemplateName string   `json:"roleTemplateName" protobuf:"bytes,3,opt,name=roleTemplateName"`
	Clusters         []string `json:"clusters" protobuf:"bytes,4,rep,name=clusters"`
}

type ClusterRoleTemplateBindingStatus struct {
	// Phase the release is in, one of ('Installing', 'Succeeded', 'Failed')
	// +optional
	Phase RoleTemplatePhase `json:"phase" protobuf:"bytes,1,opt,name=phase"`
	// The last time the condition transitioned from one status to another.
	// +optional
	LastTransitionTime metav1.Time `json:"lastTransitionTime" protobuf:"bytes,2,opt,name=lastTransitionTime"`
	// The reason for the condition's last transition.
	// +optional
	Reason string `json:"reason" protobuf:"bytes,3,opt,name=reason"`
	// A human readable message indicating details about the transition.
	// +optional
	Message string `json:"message" protobuf:"bytes,4,opt,name=message"`
	// +optional
	Clusters []ClusterRoleTemplateBindingStatusItem `json:"clusters" protobuf:"bytes,5,rep,name=clusters"`
}

type ClusterRoleTemplateBindingStatusItem struct {
	// Phase the release is in, one of ('Installing', 'Succeeded', 'Failed')
	// +optional
	Phase RoleTemplatePhase `json:"phase" protobuf:"bytes,1,opt,name=phase"`
	// The last time the condition transitioned from one status to another.
	// +optional
	LastTransitionTime metav1.Time `json:"lastTransitionTime" protobuf:"bytes,2,opt,name=lastTransitionTime"`
	// The reason for the condition's last transition.
	// +optional
	Reason string `json:"reason" protobuf:"bytes,3,opt,name=reason"`
	// A human readable message indicating details about the transition.
	// +optional
	Message string `json:"message" protobuf:"bytes,4,opt,name=message"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterRoleTemplateBindingList is a resource containing a list of ClusterRoleTemplateBinding objects.
type ClusterRoleTemplateBindingList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	// Items is the list of ConfigMaps.
	Items []ClusterRoleTemplateBinding `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +genclient
// +genclient:nonNamespaced
// +genclient:skipVerbs=deleteCollection
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ConfigMap holds configuration data for tke to consume.
type ConfigMap struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Data contains the configuration data.
	// Each key must consist of alphanumeric characters, '-', '_' or '.'.
	// Values with non-UTF-8 byte sequences must use the BinaryData field.
	// The keys stored in Data must not overlap with the keys in
	// the BinaryData field, this is enforced during validation process.
	// +optional
	Data map[string]string `json:"data,omitempty" protobuf:"bytes,2,rep,name=data"`

	// BinaryData contains the binary data.
	// Each key must consist of alphanumeric characters, '-', '_' or '.'.
	// BinaryData can contain byte sequences that are not in the UTF-8 range.
	// The keys stored in BinaryData must not overlap with the ones in
	// the Data field, this is enforced during validation process.
	// +optional
	BinaryData map[string][]byte `json:"binaryData,omitempty" protobuf:"bytes,3,rep,name=binaryData"`
}

// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ConfigMapList is a resource containing a list of ConfigMap objects.
type ConfigMapList struct {
	metav1.TypeMeta `json:",inline"`

	// +optional
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Items is the list of ConfigMaps.
	Items []ConfigMap `json:"items" protobuf:"bytes,2,rep,name=items"`
}
