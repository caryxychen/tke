// +build !ignore_autogenerated

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

// Code generated by conversion-gen. DO NOT EDIT.

package v1

import (
	unsafe "unsafe"

	rbacv1 "k8s.io/api/rbac/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	authz "tkestack.io/tke/api/authz"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*ClusterRoleTemplateBinding)(nil), (*authz.ClusterRoleTemplateBinding)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ClusterRoleTemplateBinding_To_authz_ClusterRoleTemplateBinding(a.(*ClusterRoleTemplateBinding), b.(*authz.ClusterRoleTemplateBinding), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*authz.ClusterRoleTemplateBinding)(nil), (*ClusterRoleTemplateBinding)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_authz_ClusterRoleTemplateBinding_To_v1_ClusterRoleTemplateBinding(a.(*authz.ClusterRoleTemplateBinding), b.(*ClusterRoleTemplateBinding), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ClusterRoleTemplateBindingList)(nil), (*authz.ClusterRoleTemplateBindingList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ClusterRoleTemplateBindingList_To_authz_ClusterRoleTemplateBindingList(a.(*ClusterRoleTemplateBindingList), b.(*authz.ClusterRoleTemplateBindingList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*authz.ClusterRoleTemplateBindingList)(nil), (*ClusterRoleTemplateBindingList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_authz_ClusterRoleTemplateBindingList_To_v1_ClusterRoleTemplateBindingList(a.(*authz.ClusterRoleTemplateBindingList), b.(*ClusterRoleTemplateBindingList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ClusterRoleTemplateBindingSpec)(nil), (*authz.ClusterRoleTemplateBindingSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ClusterRoleTemplateBindingSpec_To_authz_ClusterRoleTemplateBindingSpec(a.(*ClusterRoleTemplateBindingSpec), b.(*authz.ClusterRoleTemplateBindingSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*authz.ClusterRoleTemplateBindingSpec)(nil), (*ClusterRoleTemplateBindingSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_authz_ClusterRoleTemplateBindingSpec_To_v1_ClusterRoleTemplateBindingSpec(a.(*authz.ClusterRoleTemplateBindingSpec), b.(*ClusterRoleTemplateBindingSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ClusterRoleTemplateBindingStatus)(nil), (*authz.ClusterRoleTemplateBindingStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ClusterRoleTemplateBindingStatus_To_authz_ClusterRoleTemplateBindingStatus(a.(*ClusterRoleTemplateBindingStatus), b.(*authz.ClusterRoleTemplateBindingStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*authz.ClusterRoleTemplateBindingStatus)(nil), (*ClusterRoleTemplateBindingStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_authz_ClusterRoleTemplateBindingStatus_To_v1_ClusterRoleTemplateBindingStatus(a.(*authz.ClusterRoleTemplateBindingStatus), b.(*ClusterRoleTemplateBindingStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ClusterRoleTemplateBindingStatusItem)(nil), (*authz.ClusterRoleTemplateBindingStatusItem)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ClusterRoleTemplateBindingStatusItem_To_authz_ClusterRoleTemplateBindingStatusItem(a.(*ClusterRoleTemplateBindingStatusItem), b.(*authz.ClusterRoleTemplateBindingStatusItem), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*authz.ClusterRoleTemplateBindingStatusItem)(nil), (*ClusterRoleTemplateBindingStatusItem)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_authz_ClusterRoleTemplateBindingStatusItem_To_v1_ClusterRoleTemplateBindingStatusItem(a.(*authz.ClusterRoleTemplateBindingStatusItem), b.(*ClusterRoleTemplateBindingStatusItem), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ConfigMap)(nil), (*authz.ConfigMap)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ConfigMap_To_authz_ConfigMap(a.(*ConfigMap), b.(*authz.ConfigMap), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*authz.ConfigMap)(nil), (*ConfigMap)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_authz_ConfigMap_To_v1_ConfigMap(a.(*authz.ConfigMap), b.(*ConfigMap), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ConfigMapList)(nil), (*authz.ConfigMapList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ConfigMapList_To_authz_ConfigMapList(a.(*ConfigMapList), b.(*authz.ConfigMapList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*authz.ConfigMapList)(nil), (*ConfigMapList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_authz_ConfigMapList_To_v1_ConfigMapList(a.(*authz.ConfigMapList), b.(*ConfigMapList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Role)(nil), (*authz.Role)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_Role_To_authz_Role(a.(*Role), b.(*authz.Role), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*authz.Role)(nil), (*Role)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_authz_Role_To_v1_Role(a.(*authz.Role), b.(*Role), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*RoleBinding)(nil), (*authz.RoleBinding)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_RoleBinding_To_authz_RoleBinding(a.(*RoleBinding), b.(*authz.RoleBinding), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*authz.RoleBinding)(nil), (*RoleBinding)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_authz_RoleBinding_To_v1_RoleBinding(a.(*authz.RoleBinding), b.(*RoleBinding), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*RoleBindingList)(nil), (*authz.RoleBindingList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_RoleBindingList_To_authz_RoleBindingList(a.(*RoleBindingList), b.(*authz.RoleBindingList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*authz.RoleBindingList)(nil), (*RoleBindingList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_authz_RoleBindingList_To_v1_RoleBindingList(a.(*authz.RoleBindingList), b.(*RoleBindingList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*RoleBindingSpec)(nil), (*authz.RoleBindingSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_RoleBindingSpec_To_authz_RoleBindingSpec(a.(*RoleBindingSpec), b.(*authz.RoleBindingSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*authz.RoleBindingSpec)(nil), (*RoleBindingSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_authz_RoleBindingSpec_To_v1_RoleBindingSpec(a.(*authz.RoleBindingSpec), b.(*RoleBindingSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*RoleList)(nil), (*authz.RoleList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_RoleList_To_authz_RoleList(a.(*RoleList), b.(*authz.RoleList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*authz.RoleList)(nil), (*RoleList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_authz_RoleList_To_v1_RoleList(a.(*authz.RoleList), b.(*RoleList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*RoleTemplate)(nil), (*authz.RoleTemplate)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_RoleTemplate_To_authz_RoleTemplate(a.(*RoleTemplate), b.(*authz.RoleTemplate), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*authz.RoleTemplate)(nil), (*RoleTemplate)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_authz_RoleTemplate_To_v1_RoleTemplate(a.(*authz.RoleTemplate), b.(*RoleTemplate), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*RoleTemplateList)(nil), (*authz.RoleTemplateList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_RoleTemplateList_To_authz_RoleTemplateList(a.(*RoleTemplateList), b.(*authz.RoleTemplateList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*authz.RoleTemplateList)(nil), (*RoleTemplateList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_authz_RoleTemplateList_To_v1_RoleTemplateList(a.(*authz.RoleTemplateList), b.(*RoleTemplateList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*RoleTemplateSpec)(nil), (*authz.RoleTemplateSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_RoleTemplateSpec_To_authz_RoleTemplateSpec(a.(*RoleTemplateSpec), b.(*authz.RoleTemplateSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*authz.RoleTemplateSpec)(nil), (*RoleTemplateSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_authz_RoleTemplateSpec_To_v1_RoleTemplateSpec(a.(*authz.RoleTemplateSpec), b.(*RoleTemplateSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*RoleTemplateStatus)(nil), (*authz.RoleTemplateStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_RoleTemplateStatus_To_authz_RoleTemplateStatus(a.(*RoleTemplateStatus), b.(*authz.RoleTemplateStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*authz.RoleTemplateStatus)(nil), (*RoleTemplateStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_authz_RoleTemplateStatus_To_v1_RoleTemplateStatus(a.(*authz.RoleTemplateStatus), b.(*RoleTemplateStatus), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1_ClusterRoleTemplateBinding_To_authz_ClusterRoleTemplateBinding(in *ClusterRoleTemplateBinding, out *authz.ClusterRoleTemplateBinding, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_ClusterRoleTemplateBindingSpec_To_authz_ClusterRoleTemplateBindingSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1_ClusterRoleTemplateBindingStatus_To_authz_ClusterRoleTemplateBindingStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_ClusterRoleTemplateBinding_To_authz_ClusterRoleTemplateBinding is an autogenerated conversion function.
func Convert_v1_ClusterRoleTemplateBinding_To_authz_ClusterRoleTemplateBinding(in *ClusterRoleTemplateBinding, out *authz.ClusterRoleTemplateBinding, s conversion.Scope) error {
	return autoConvert_v1_ClusterRoleTemplateBinding_To_authz_ClusterRoleTemplateBinding(in, out, s)
}

func autoConvert_authz_ClusterRoleTemplateBinding_To_v1_ClusterRoleTemplateBinding(in *authz.ClusterRoleTemplateBinding, out *ClusterRoleTemplateBinding, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_authz_ClusterRoleTemplateBindingSpec_To_v1_ClusterRoleTemplateBindingSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_authz_ClusterRoleTemplateBindingStatus_To_v1_ClusterRoleTemplateBindingStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_authz_ClusterRoleTemplateBinding_To_v1_ClusterRoleTemplateBinding is an autogenerated conversion function.
func Convert_authz_ClusterRoleTemplateBinding_To_v1_ClusterRoleTemplateBinding(in *authz.ClusterRoleTemplateBinding, out *ClusterRoleTemplateBinding, s conversion.Scope) error {
	return autoConvert_authz_ClusterRoleTemplateBinding_To_v1_ClusterRoleTemplateBinding(in, out, s)
}

func autoConvert_v1_ClusterRoleTemplateBindingList_To_authz_ClusterRoleTemplateBindingList(in *ClusterRoleTemplateBindingList, out *authz.ClusterRoleTemplateBindingList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]authz.ClusterRoleTemplateBinding)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1_ClusterRoleTemplateBindingList_To_authz_ClusterRoleTemplateBindingList is an autogenerated conversion function.
func Convert_v1_ClusterRoleTemplateBindingList_To_authz_ClusterRoleTemplateBindingList(in *ClusterRoleTemplateBindingList, out *authz.ClusterRoleTemplateBindingList, s conversion.Scope) error {
	return autoConvert_v1_ClusterRoleTemplateBindingList_To_authz_ClusterRoleTemplateBindingList(in, out, s)
}

func autoConvert_authz_ClusterRoleTemplateBindingList_To_v1_ClusterRoleTemplateBindingList(in *authz.ClusterRoleTemplateBindingList, out *ClusterRoleTemplateBindingList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]ClusterRoleTemplateBinding)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_authz_ClusterRoleTemplateBindingList_To_v1_ClusterRoleTemplateBindingList is an autogenerated conversion function.
func Convert_authz_ClusterRoleTemplateBindingList_To_v1_ClusterRoleTemplateBindingList(in *authz.ClusterRoleTemplateBindingList, out *ClusterRoleTemplateBindingList, s conversion.Scope) error {
	return autoConvert_authz_ClusterRoleTemplateBindingList_To_v1_ClusterRoleTemplateBindingList(in, out, s)
}

func autoConvert_v1_ClusterRoleTemplateBindingSpec_To_authz_ClusterRoleTemplateBindingSpec(in *ClusterRoleTemplateBindingSpec, out *authz.ClusterRoleTemplateBindingSpec, s conversion.Scope) error {
	out.UserName = in.UserName
	out.GroupName = in.GroupName
	out.RoleTemplateName = in.RoleTemplateName
	out.Clusters = *(*[]string)(unsafe.Pointer(&in.Clusters))
	return nil
}

// Convert_v1_ClusterRoleTemplateBindingSpec_To_authz_ClusterRoleTemplateBindingSpec is an autogenerated conversion function.
func Convert_v1_ClusterRoleTemplateBindingSpec_To_authz_ClusterRoleTemplateBindingSpec(in *ClusterRoleTemplateBindingSpec, out *authz.ClusterRoleTemplateBindingSpec, s conversion.Scope) error {
	return autoConvert_v1_ClusterRoleTemplateBindingSpec_To_authz_ClusterRoleTemplateBindingSpec(in, out, s)
}

func autoConvert_authz_ClusterRoleTemplateBindingSpec_To_v1_ClusterRoleTemplateBindingSpec(in *authz.ClusterRoleTemplateBindingSpec, out *ClusterRoleTemplateBindingSpec, s conversion.Scope) error {
	out.UserName = in.UserName
	out.GroupName = in.GroupName
	out.RoleTemplateName = in.RoleTemplateName
	out.Clusters = *(*[]string)(unsafe.Pointer(&in.Clusters))
	return nil
}

// Convert_authz_ClusterRoleTemplateBindingSpec_To_v1_ClusterRoleTemplateBindingSpec is an autogenerated conversion function.
func Convert_authz_ClusterRoleTemplateBindingSpec_To_v1_ClusterRoleTemplateBindingSpec(in *authz.ClusterRoleTemplateBindingSpec, out *ClusterRoleTemplateBindingSpec, s conversion.Scope) error {
	return autoConvert_authz_ClusterRoleTemplateBindingSpec_To_v1_ClusterRoleTemplateBindingSpec(in, out, s)
}

func autoConvert_v1_ClusterRoleTemplateBindingStatus_To_authz_ClusterRoleTemplateBindingStatus(in *ClusterRoleTemplateBindingStatus, out *authz.ClusterRoleTemplateBindingStatus, s conversion.Scope) error {
	out.Phase = authz.RoleTemplatePhase(in.Phase)
	out.LastTransitionTime = in.LastTransitionTime
	out.Reason = in.Reason
	out.Message = in.Message
	out.Clusters = *(*[]authz.ClusterRoleTemplateBindingStatusItem)(unsafe.Pointer(&in.Clusters))
	return nil
}

// Convert_v1_ClusterRoleTemplateBindingStatus_To_authz_ClusterRoleTemplateBindingStatus is an autogenerated conversion function.
func Convert_v1_ClusterRoleTemplateBindingStatus_To_authz_ClusterRoleTemplateBindingStatus(in *ClusterRoleTemplateBindingStatus, out *authz.ClusterRoleTemplateBindingStatus, s conversion.Scope) error {
	return autoConvert_v1_ClusterRoleTemplateBindingStatus_To_authz_ClusterRoleTemplateBindingStatus(in, out, s)
}

func autoConvert_authz_ClusterRoleTemplateBindingStatus_To_v1_ClusterRoleTemplateBindingStatus(in *authz.ClusterRoleTemplateBindingStatus, out *ClusterRoleTemplateBindingStatus, s conversion.Scope) error {
	out.Phase = RoleTemplatePhase(in.Phase)
	out.LastTransitionTime = in.LastTransitionTime
	out.Reason = in.Reason
	out.Message = in.Message
	out.Clusters = *(*[]ClusterRoleTemplateBindingStatusItem)(unsafe.Pointer(&in.Clusters))
	return nil
}

// Convert_authz_ClusterRoleTemplateBindingStatus_To_v1_ClusterRoleTemplateBindingStatus is an autogenerated conversion function.
func Convert_authz_ClusterRoleTemplateBindingStatus_To_v1_ClusterRoleTemplateBindingStatus(in *authz.ClusterRoleTemplateBindingStatus, out *ClusterRoleTemplateBindingStatus, s conversion.Scope) error {
	return autoConvert_authz_ClusterRoleTemplateBindingStatus_To_v1_ClusterRoleTemplateBindingStatus(in, out, s)
}

func autoConvert_v1_ClusterRoleTemplateBindingStatusItem_To_authz_ClusterRoleTemplateBindingStatusItem(in *ClusterRoleTemplateBindingStatusItem, out *authz.ClusterRoleTemplateBindingStatusItem, s conversion.Scope) error {
	out.Phase = authz.RoleTemplatePhase(in.Phase)
	out.LastTransitionTime = in.LastTransitionTime
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}

// Convert_v1_ClusterRoleTemplateBindingStatusItem_To_authz_ClusterRoleTemplateBindingStatusItem is an autogenerated conversion function.
func Convert_v1_ClusterRoleTemplateBindingStatusItem_To_authz_ClusterRoleTemplateBindingStatusItem(in *ClusterRoleTemplateBindingStatusItem, out *authz.ClusterRoleTemplateBindingStatusItem, s conversion.Scope) error {
	return autoConvert_v1_ClusterRoleTemplateBindingStatusItem_To_authz_ClusterRoleTemplateBindingStatusItem(in, out, s)
}

func autoConvert_authz_ClusterRoleTemplateBindingStatusItem_To_v1_ClusterRoleTemplateBindingStatusItem(in *authz.ClusterRoleTemplateBindingStatusItem, out *ClusterRoleTemplateBindingStatusItem, s conversion.Scope) error {
	out.Phase = RoleTemplatePhase(in.Phase)
	out.LastTransitionTime = in.LastTransitionTime
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}

// Convert_authz_ClusterRoleTemplateBindingStatusItem_To_v1_ClusterRoleTemplateBindingStatusItem is an autogenerated conversion function.
func Convert_authz_ClusterRoleTemplateBindingStatusItem_To_v1_ClusterRoleTemplateBindingStatusItem(in *authz.ClusterRoleTemplateBindingStatusItem, out *ClusterRoleTemplateBindingStatusItem, s conversion.Scope) error {
	return autoConvert_authz_ClusterRoleTemplateBindingStatusItem_To_v1_ClusterRoleTemplateBindingStatusItem(in, out, s)
}

func autoConvert_v1_ConfigMap_To_authz_ConfigMap(in *ConfigMap, out *authz.ConfigMap, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Data = *(*map[string]string)(unsafe.Pointer(&in.Data))
	out.BinaryData = *(*map[string][]byte)(unsafe.Pointer(&in.BinaryData))
	return nil
}

// Convert_v1_ConfigMap_To_authz_ConfigMap is an autogenerated conversion function.
func Convert_v1_ConfigMap_To_authz_ConfigMap(in *ConfigMap, out *authz.ConfigMap, s conversion.Scope) error {
	return autoConvert_v1_ConfigMap_To_authz_ConfigMap(in, out, s)
}

func autoConvert_authz_ConfigMap_To_v1_ConfigMap(in *authz.ConfigMap, out *ConfigMap, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Data = *(*map[string]string)(unsafe.Pointer(&in.Data))
	out.BinaryData = *(*map[string][]byte)(unsafe.Pointer(&in.BinaryData))
	return nil
}

// Convert_authz_ConfigMap_To_v1_ConfigMap is an autogenerated conversion function.
func Convert_authz_ConfigMap_To_v1_ConfigMap(in *authz.ConfigMap, out *ConfigMap, s conversion.Scope) error {
	return autoConvert_authz_ConfigMap_To_v1_ConfigMap(in, out, s)
}

func autoConvert_v1_ConfigMapList_To_authz_ConfigMapList(in *ConfigMapList, out *authz.ConfigMapList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]authz.ConfigMap)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1_ConfigMapList_To_authz_ConfigMapList is an autogenerated conversion function.
func Convert_v1_ConfigMapList_To_authz_ConfigMapList(in *ConfigMapList, out *authz.ConfigMapList, s conversion.Scope) error {
	return autoConvert_v1_ConfigMapList_To_authz_ConfigMapList(in, out, s)
}

func autoConvert_authz_ConfigMapList_To_v1_ConfigMapList(in *authz.ConfigMapList, out *ConfigMapList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]ConfigMap)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_authz_ConfigMapList_To_v1_ConfigMapList is an autogenerated conversion function.
func Convert_authz_ConfigMapList_To_v1_ConfigMapList(in *authz.ConfigMapList, out *ConfigMapList, s conversion.Scope) error {
	return autoConvert_authz_ConfigMapList_To_v1_ConfigMapList(in, out, s)
}

func autoConvert_v1_Role_To_authz_Role(in *Role, out *authz.Role, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.DisplayName = in.DisplayName
	out.TenantID = in.TenantID
	out.Username = in.Username
	out.Description = in.Description
	out.Policies = *(*[]string)(unsafe.Pointer(&in.Policies))
	return nil
}

// Convert_v1_Role_To_authz_Role is an autogenerated conversion function.
func Convert_v1_Role_To_authz_Role(in *Role, out *authz.Role, s conversion.Scope) error {
	return autoConvert_v1_Role_To_authz_Role(in, out, s)
}

func autoConvert_authz_Role_To_v1_Role(in *authz.Role, out *Role, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.DisplayName = in.DisplayName
	out.TenantID = in.TenantID
	out.Username = in.Username
	out.Description = in.Description
	out.Policies = *(*[]string)(unsafe.Pointer(&in.Policies))
	return nil
}

// Convert_authz_Role_To_v1_Role is an autogenerated conversion function.
func Convert_authz_Role_To_v1_Role(in *authz.Role, out *Role, s conversion.Scope) error {
	return autoConvert_authz_Role_To_v1_Role(in, out, s)
}

func autoConvert_v1_RoleBinding_To_authz_RoleBinding(in *RoleBinding, out *authz.RoleBinding, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_RoleBindingSpec_To_authz_RoleBindingSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_RoleBinding_To_authz_RoleBinding is an autogenerated conversion function.
func Convert_v1_RoleBinding_To_authz_RoleBinding(in *RoleBinding, out *authz.RoleBinding, s conversion.Scope) error {
	return autoConvert_v1_RoleBinding_To_authz_RoleBinding(in, out, s)
}

func autoConvert_authz_RoleBinding_To_v1_RoleBinding(in *authz.RoleBinding, out *RoleBinding, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_authz_RoleBindingSpec_To_v1_RoleBindingSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_authz_RoleBinding_To_v1_RoleBinding is an autogenerated conversion function.
func Convert_authz_RoleBinding_To_v1_RoleBinding(in *authz.RoleBinding, out *RoleBinding, s conversion.Scope) error {
	return autoConvert_authz_RoleBinding_To_v1_RoleBinding(in, out, s)
}

func autoConvert_v1_RoleBindingList_To_authz_RoleBindingList(in *RoleBindingList, out *authz.RoleBindingList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]authz.RoleBindingList)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1_RoleBindingList_To_authz_RoleBindingList is an autogenerated conversion function.
func Convert_v1_RoleBindingList_To_authz_RoleBindingList(in *RoleBindingList, out *authz.RoleBindingList, s conversion.Scope) error {
	return autoConvert_v1_RoleBindingList_To_authz_RoleBindingList(in, out, s)
}

func autoConvert_authz_RoleBindingList_To_v1_RoleBindingList(in *authz.RoleBindingList, out *RoleBindingList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]RoleBindingList)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_authz_RoleBindingList_To_v1_RoleBindingList is an autogenerated conversion function.
func Convert_authz_RoleBindingList_To_v1_RoleBindingList(in *authz.RoleBindingList, out *RoleBindingList, s conversion.Scope) error {
	return autoConvert_authz_RoleBindingList_To_v1_RoleBindingList(in, out, s)
}

func autoConvert_v1_RoleBindingSpec_To_authz_RoleBindingSpec(in *RoleBindingSpec, out *authz.RoleBindingSpec, s conversion.Scope) error {
	out.UserName = in.UserName
	out.GroupName = in.GroupName
	out.RoleName = in.RoleName
	out.Clusters = *(*[]string)(unsafe.Pointer(&in.Clusters))
	return nil
}

// Convert_v1_RoleBindingSpec_To_authz_RoleBindingSpec is an autogenerated conversion function.
func Convert_v1_RoleBindingSpec_To_authz_RoleBindingSpec(in *RoleBindingSpec, out *authz.RoleBindingSpec, s conversion.Scope) error {
	return autoConvert_v1_RoleBindingSpec_To_authz_RoleBindingSpec(in, out, s)
}

func autoConvert_authz_RoleBindingSpec_To_v1_RoleBindingSpec(in *authz.RoleBindingSpec, out *RoleBindingSpec, s conversion.Scope) error {
	out.UserName = in.UserName
	out.GroupName = in.GroupName
	out.RoleName = in.RoleName
	out.Clusters = *(*[]string)(unsafe.Pointer(&in.Clusters))
	return nil
}

// Convert_authz_RoleBindingSpec_To_v1_RoleBindingSpec is an autogenerated conversion function.
func Convert_authz_RoleBindingSpec_To_v1_RoleBindingSpec(in *authz.RoleBindingSpec, out *RoleBindingSpec, s conversion.Scope) error {
	return autoConvert_authz_RoleBindingSpec_To_v1_RoleBindingSpec(in, out, s)
}

func autoConvert_v1_RoleList_To_authz_RoleList(in *RoleList, out *authz.RoleList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]authz.Role)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1_RoleList_To_authz_RoleList is an autogenerated conversion function.
func Convert_v1_RoleList_To_authz_RoleList(in *RoleList, out *authz.RoleList, s conversion.Scope) error {
	return autoConvert_v1_RoleList_To_authz_RoleList(in, out, s)
}

func autoConvert_authz_RoleList_To_v1_RoleList(in *authz.RoleList, out *RoleList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]Role)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_authz_RoleList_To_v1_RoleList is an autogenerated conversion function.
func Convert_authz_RoleList_To_v1_RoleList(in *authz.RoleList, out *RoleList, s conversion.Scope) error {
	return autoConvert_authz_RoleList_To_v1_RoleList(in, out, s)
}

func autoConvert_v1_RoleTemplate_To_authz_RoleTemplate(in *RoleTemplate, out *authz.RoleTemplate, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_RoleTemplateSpec_To_authz_RoleTemplateSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1_RoleTemplateStatus_To_authz_RoleTemplateStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_RoleTemplate_To_authz_RoleTemplate is an autogenerated conversion function.
func Convert_v1_RoleTemplate_To_authz_RoleTemplate(in *RoleTemplate, out *authz.RoleTemplate, s conversion.Scope) error {
	return autoConvert_v1_RoleTemplate_To_authz_RoleTemplate(in, out, s)
}

func autoConvert_authz_RoleTemplate_To_v1_RoleTemplate(in *authz.RoleTemplate, out *RoleTemplate, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_authz_RoleTemplateSpec_To_v1_RoleTemplateSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_authz_RoleTemplateStatus_To_v1_RoleTemplateStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_authz_RoleTemplate_To_v1_RoleTemplate is an autogenerated conversion function.
func Convert_authz_RoleTemplate_To_v1_RoleTemplate(in *authz.RoleTemplate, out *RoleTemplate, s conversion.Scope) error {
	return autoConvert_authz_RoleTemplate_To_v1_RoleTemplate(in, out, s)
}

func autoConvert_v1_RoleTemplateList_To_authz_RoleTemplateList(in *RoleTemplateList, out *authz.RoleTemplateList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]authz.RoleTemplate)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1_RoleTemplateList_To_authz_RoleTemplateList is an autogenerated conversion function.
func Convert_v1_RoleTemplateList_To_authz_RoleTemplateList(in *RoleTemplateList, out *authz.RoleTemplateList, s conversion.Scope) error {
	return autoConvert_v1_RoleTemplateList_To_authz_RoleTemplateList(in, out, s)
}

func autoConvert_authz_RoleTemplateList_To_v1_RoleTemplateList(in *authz.RoleTemplateList, out *RoleTemplateList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]RoleTemplate)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_authz_RoleTemplateList_To_v1_RoleTemplateList is an autogenerated conversion function.
func Convert_authz_RoleTemplateList_To_v1_RoleTemplateList(in *authz.RoleTemplateList, out *RoleTemplateList, s conversion.Scope) error {
	return autoConvert_authz_RoleTemplateList_To_v1_RoleTemplateList(in, out, s)
}

func autoConvert_v1_RoleTemplateSpec_To_authz_RoleTemplateSpec(in *RoleTemplateSpec, out *authz.RoleTemplateSpec, s conversion.Scope) error {
	out.TenantID = in.TenantID
	out.DisplayName = in.DisplayName
	out.Description = in.Description
	out.Scope = authz.Scope(in.Scope)
	out.Rules = *(*[]rbacv1.PolicyRule)(unsafe.Pointer(&in.Rules))
	return nil
}

// Convert_v1_RoleTemplateSpec_To_authz_RoleTemplateSpec is an autogenerated conversion function.
func Convert_v1_RoleTemplateSpec_To_authz_RoleTemplateSpec(in *RoleTemplateSpec, out *authz.RoleTemplateSpec, s conversion.Scope) error {
	return autoConvert_v1_RoleTemplateSpec_To_authz_RoleTemplateSpec(in, out, s)
}

func autoConvert_authz_RoleTemplateSpec_To_v1_RoleTemplateSpec(in *authz.RoleTemplateSpec, out *RoleTemplateSpec, s conversion.Scope) error {
	out.TenantID = in.TenantID
	out.DisplayName = in.DisplayName
	out.Description = in.Description
	out.Scope = Scope(in.Scope)
	out.Rules = *(*[]rbacv1.PolicyRule)(unsafe.Pointer(&in.Rules))
	return nil
}

// Convert_authz_RoleTemplateSpec_To_v1_RoleTemplateSpec is an autogenerated conversion function.
func Convert_authz_RoleTemplateSpec_To_v1_RoleTemplateSpec(in *authz.RoleTemplateSpec, out *RoleTemplateSpec, s conversion.Scope) error {
	return autoConvert_authz_RoleTemplateSpec_To_v1_RoleTemplateSpec(in, out, s)
}

func autoConvert_v1_RoleTemplateStatus_To_authz_RoleTemplateStatus(in *RoleTemplateStatus, out *authz.RoleTemplateStatus, s conversion.Scope) error {
	out.Phase = authz.RoleTemplatePhase(in.Phase)
	out.LastTransitionTime = in.LastTransitionTime
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}

// Convert_v1_RoleTemplateStatus_To_authz_RoleTemplateStatus is an autogenerated conversion function.
func Convert_v1_RoleTemplateStatus_To_authz_RoleTemplateStatus(in *RoleTemplateStatus, out *authz.RoleTemplateStatus, s conversion.Scope) error {
	return autoConvert_v1_RoleTemplateStatus_To_authz_RoleTemplateStatus(in, out, s)
}

func autoConvert_authz_RoleTemplateStatus_To_v1_RoleTemplateStatus(in *authz.RoleTemplateStatus, out *RoleTemplateStatus, s conversion.Scope) error {
	out.Phase = RoleTemplatePhase(in.Phase)
	out.LastTransitionTime = in.LastTransitionTime
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}

// Convert_authz_RoleTemplateStatus_To_v1_RoleTemplateStatus is an autogenerated conversion function.
func Convert_authz_RoleTemplateStatus_To_v1_RoleTemplateStatus(in *authz.RoleTemplateStatus, out *RoleTemplateStatus, s conversion.Scope) error {
	return autoConvert_authz_RoleTemplateStatus_To_v1_RoleTemplateStatus(in, out, s)
}
