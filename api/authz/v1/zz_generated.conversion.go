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
	if err := s.AddGeneratedConversionFunc((*ClusterPolicyBinding)(nil), (*authz.ClusterPolicyBinding)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ClusterPolicyBinding_To_authz_ClusterPolicyBinding(a.(*ClusterPolicyBinding), b.(*authz.ClusterPolicyBinding), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*authz.ClusterPolicyBinding)(nil), (*ClusterPolicyBinding)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_authz_ClusterPolicyBinding_To_v1_ClusterPolicyBinding(a.(*authz.ClusterPolicyBinding), b.(*ClusterPolicyBinding), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ClusterPolicyBindingList)(nil), (*authz.ClusterPolicyBindingList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ClusterPolicyBindingList_To_authz_ClusterPolicyBindingList(a.(*ClusterPolicyBindingList), b.(*authz.ClusterPolicyBindingList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*authz.ClusterPolicyBindingList)(nil), (*ClusterPolicyBindingList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_authz_ClusterPolicyBindingList_To_v1_ClusterPolicyBindingList(a.(*authz.ClusterPolicyBindingList), b.(*ClusterPolicyBindingList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ClusterPolicyBindingSpec)(nil), (*authz.ClusterPolicyBindingSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ClusterPolicyBindingSpec_To_authz_ClusterPolicyBindingSpec(a.(*ClusterPolicyBindingSpec), b.(*authz.ClusterPolicyBindingSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*authz.ClusterPolicyBindingSpec)(nil), (*ClusterPolicyBindingSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_authz_ClusterPolicyBindingSpec_To_v1_ClusterPolicyBindingSpec(a.(*authz.ClusterPolicyBindingSpec), b.(*ClusterPolicyBindingSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ClusterPolicyBindingStatus)(nil), (*authz.ClusterPolicyBindingStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ClusterPolicyBindingStatus_To_authz_ClusterPolicyBindingStatus(a.(*ClusterPolicyBindingStatus), b.(*authz.ClusterPolicyBindingStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*authz.ClusterPolicyBindingStatus)(nil), (*ClusterPolicyBindingStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_authz_ClusterPolicyBindingStatus_To_v1_ClusterPolicyBindingStatus(a.(*authz.ClusterPolicyBindingStatus), b.(*ClusterPolicyBindingStatus), scope)
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
	if err := s.AddGeneratedConversionFunc((*Policy)(nil), (*authz.Policy)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_Policy_To_authz_Policy(a.(*Policy), b.(*authz.Policy), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*authz.Policy)(nil), (*Policy)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_authz_Policy_To_v1_Policy(a.(*authz.Policy), b.(*Policy), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*PolicyList)(nil), (*authz.PolicyList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_PolicyList_To_authz_PolicyList(a.(*PolicyList), b.(*authz.PolicyList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*authz.PolicyList)(nil), (*PolicyList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_authz_PolicyList_To_v1_PolicyList(a.(*authz.PolicyList), b.(*PolicyList), scope)
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
	if err := s.AddGeneratedConversionFunc((*RoleBindingStatus)(nil), (*authz.RoleBindingStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_RoleBindingStatus_To_authz_RoleBindingStatus(a.(*RoleBindingStatus), b.(*authz.RoleBindingStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*authz.RoleBindingStatus)(nil), (*RoleBindingStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_authz_RoleBindingStatus_To_v1_RoleBindingStatus(a.(*authz.RoleBindingStatus), b.(*RoleBindingStatus), scope)
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
	return nil
}

func autoConvert_v1_ClusterPolicyBinding_To_authz_ClusterPolicyBinding(in *ClusterPolicyBinding, out *authz.ClusterPolicyBinding, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_ClusterPolicyBindingSpec_To_authz_ClusterPolicyBindingSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1_ClusterPolicyBindingStatus_To_authz_ClusterPolicyBindingStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_ClusterPolicyBinding_To_authz_ClusterPolicyBinding is an autogenerated conversion function.
func Convert_v1_ClusterPolicyBinding_To_authz_ClusterPolicyBinding(in *ClusterPolicyBinding, out *authz.ClusterPolicyBinding, s conversion.Scope) error {
	return autoConvert_v1_ClusterPolicyBinding_To_authz_ClusterPolicyBinding(in, out, s)
}

func autoConvert_authz_ClusterPolicyBinding_To_v1_ClusterPolicyBinding(in *authz.ClusterPolicyBinding, out *ClusterPolicyBinding, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_authz_ClusterPolicyBindingSpec_To_v1_ClusterPolicyBindingSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_authz_ClusterPolicyBindingStatus_To_v1_ClusterPolicyBindingStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_authz_ClusterPolicyBinding_To_v1_ClusterPolicyBinding is an autogenerated conversion function.
func Convert_authz_ClusterPolicyBinding_To_v1_ClusterPolicyBinding(in *authz.ClusterPolicyBinding, out *ClusterPolicyBinding, s conversion.Scope) error {
	return autoConvert_authz_ClusterPolicyBinding_To_v1_ClusterPolicyBinding(in, out, s)
}

func autoConvert_v1_ClusterPolicyBindingList_To_authz_ClusterPolicyBindingList(in *ClusterPolicyBindingList, out *authz.ClusterPolicyBindingList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]authz.ClusterPolicyBinding)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1_ClusterPolicyBindingList_To_authz_ClusterPolicyBindingList is an autogenerated conversion function.
func Convert_v1_ClusterPolicyBindingList_To_authz_ClusterPolicyBindingList(in *ClusterPolicyBindingList, out *authz.ClusterPolicyBindingList, s conversion.Scope) error {
	return autoConvert_v1_ClusterPolicyBindingList_To_authz_ClusterPolicyBindingList(in, out, s)
}

func autoConvert_authz_ClusterPolicyBindingList_To_v1_ClusterPolicyBindingList(in *authz.ClusterPolicyBindingList, out *ClusterPolicyBindingList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]ClusterPolicyBinding)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_authz_ClusterPolicyBindingList_To_v1_ClusterPolicyBindingList is an autogenerated conversion function.
func Convert_authz_ClusterPolicyBindingList_To_v1_ClusterPolicyBindingList(in *authz.ClusterPolicyBindingList, out *ClusterPolicyBindingList, s conversion.Scope) error {
	return autoConvert_authz_ClusterPolicyBindingList_To_v1_ClusterPolicyBindingList(in, out, s)
}

func autoConvert_v1_ClusterPolicyBindingSpec_To_authz_ClusterPolicyBindingSpec(in *ClusterPolicyBindingSpec, out *authz.ClusterPolicyBindingSpec, s conversion.Scope) error {
	out.UserName = in.UserName
	out.GroupName = in.GroupName
	out.PolicyName = in.PolicyName
	out.Clusters = *(*[]string)(unsafe.Pointer(&in.Clusters))
	return nil
}

// Convert_v1_ClusterPolicyBindingSpec_To_authz_ClusterPolicyBindingSpec is an autogenerated conversion function.
func Convert_v1_ClusterPolicyBindingSpec_To_authz_ClusterPolicyBindingSpec(in *ClusterPolicyBindingSpec, out *authz.ClusterPolicyBindingSpec, s conversion.Scope) error {
	return autoConvert_v1_ClusterPolicyBindingSpec_To_authz_ClusterPolicyBindingSpec(in, out, s)
}

func autoConvert_authz_ClusterPolicyBindingSpec_To_v1_ClusterPolicyBindingSpec(in *authz.ClusterPolicyBindingSpec, out *ClusterPolicyBindingSpec, s conversion.Scope) error {
	out.UserName = in.UserName
	out.GroupName = in.GroupName
	out.PolicyName = in.PolicyName
	out.Clusters = *(*[]string)(unsafe.Pointer(&in.Clusters))
	return nil
}

// Convert_authz_ClusterPolicyBindingSpec_To_v1_ClusterPolicyBindingSpec is an autogenerated conversion function.
func Convert_authz_ClusterPolicyBindingSpec_To_v1_ClusterPolicyBindingSpec(in *authz.ClusterPolicyBindingSpec, out *ClusterPolicyBindingSpec, s conversion.Scope) error {
	return autoConvert_authz_ClusterPolicyBindingSpec_To_v1_ClusterPolicyBindingSpec(in, out, s)
}

func autoConvert_v1_ClusterPolicyBindingStatus_To_authz_ClusterPolicyBindingStatus(in *ClusterPolicyBindingStatus, out *authz.ClusterPolicyBindingStatus, s conversion.Scope) error {
	out.Phase = authz.BindingPhase(in.Phase)
	return nil
}

// Convert_v1_ClusterPolicyBindingStatus_To_authz_ClusterPolicyBindingStatus is an autogenerated conversion function.
func Convert_v1_ClusterPolicyBindingStatus_To_authz_ClusterPolicyBindingStatus(in *ClusterPolicyBindingStatus, out *authz.ClusterPolicyBindingStatus, s conversion.Scope) error {
	return autoConvert_v1_ClusterPolicyBindingStatus_To_authz_ClusterPolicyBindingStatus(in, out, s)
}

func autoConvert_authz_ClusterPolicyBindingStatus_To_v1_ClusterPolicyBindingStatus(in *authz.ClusterPolicyBindingStatus, out *ClusterPolicyBindingStatus, s conversion.Scope) error {
	out.Phase = BindingPhase(in.Phase)
	return nil
}

// Convert_authz_ClusterPolicyBindingStatus_To_v1_ClusterPolicyBindingStatus is an autogenerated conversion function.
func Convert_authz_ClusterPolicyBindingStatus_To_v1_ClusterPolicyBindingStatus(in *authz.ClusterPolicyBindingStatus, out *ClusterPolicyBindingStatus, s conversion.Scope) error {
	return autoConvert_authz_ClusterPolicyBindingStatus_To_v1_ClusterPolicyBindingStatus(in, out, s)
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

func autoConvert_v1_Policy_To_authz_Policy(in *Policy, out *authz.Policy, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.TenantID = in.TenantID
	out.DisplayName = in.DisplayName
	out.Description = in.Description
	out.Scope = authz.Scope(in.Scope)
	out.Rules = *(*[]rbacv1.PolicyRule)(unsafe.Pointer(&in.Rules))
	return nil
}

// Convert_v1_Policy_To_authz_Policy is an autogenerated conversion function.
func Convert_v1_Policy_To_authz_Policy(in *Policy, out *authz.Policy, s conversion.Scope) error {
	return autoConvert_v1_Policy_To_authz_Policy(in, out, s)
}

func autoConvert_authz_Policy_To_v1_Policy(in *authz.Policy, out *Policy, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.TenantID = in.TenantID
	out.DisplayName = in.DisplayName
	out.Description = in.Description
	out.Scope = Scope(in.Scope)
	out.Rules = *(*[]rbacv1.PolicyRule)(unsafe.Pointer(&in.Rules))
	return nil
}

// Convert_authz_Policy_To_v1_Policy is an autogenerated conversion function.
func Convert_authz_Policy_To_v1_Policy(in *authz.Policy, out *Policy, s conversion.Scope) error {
	return autoConvert_authz_Policy_To_v1_Policy(in, out, s)
}

func autoConvert_v1_PolicyList_To_authz_PolicyList(in *PolicyList, out *authz.PolicyList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]authz.Policy)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1_PolicyList_To_authz_PolicyList is an autogenerated conversion function.
func Convert_v1_PolicyList_To_authz_PolicyList(in *PolicyList, out *authz.PolicyList, s conversion.Scope) error {
	return autoConvert_v1_PolicyList_To_authz_PolicyList(in, out, s)
}

func autoConvert_authz_PolicyList_To_v1_PolicyList(in *authz.PolicyList, out *PolicyList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]Policy)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_authz_PolicyList_To_v1_PolicyList is an autogenerated conversion function.
func Convert_authz_PolicyList_To_v1_PolicyList(in *authz.PolicyList, out *PolicyList, s conversion.Scope) error {
	return autoConvert_authz_PolicyList_To_v1_PolicyList(in, out, s)
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
	if err := Convert_v1_RoleBindingStatus_To_authz_RoleBindingStatus(&in.Status, &out.Status, s); err != nil {
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
	if err := Convert_authz_RoleBindingStatus_To_v1_RoleBindingStatus(&in.Status, &out.Status, s); err != nil {
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
	out.Items = *(*[]authz.RoleBinding)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1_RoleBindingList_To_authz_RoleBindingList is an autogenerated conversion function.
func Convert_v1_RoleBindingList_To_authz_RoleBindingList(in *RoleBindingList, out *authz.RoleBindingList, s conversion.Scope) error {
	return autoConvert_v1_RoleBindingList_To_authz_RoleBindingList(in, out, s)
}

func autoConvert_authz_RoleBindingList_To_v1_RoleBindingList(in *authz.RoleBindingList, out *RoleBindingList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]RoleBinding)(unsafe.Pointer(&in.Items))
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

func autoConvert_v1_RoleBindingStatus_To_authz_RoleBindingStatus(in *RoleBindingStatus, out *authz.RoleBindingStatus, s conversion.Scope) error {
	out.Phase = authz.BindingPhase(in.Phase)
	return nil
}

// Convert_v1_RoleBindingStatus_To_authz_RoleBindingStatus is an autogenerated conversion function.
func Convert_v1_RoleBindingStatus_To_authz_RoleBindingStatus(in *RoleBindingStatus, out *authz.RoleBindingStatus, s conversion.Scope) error {
	return autoConvert_v1_RoleBindingStatus_To_authz_RoleBindingStatus(in, out, s)
}

func autoConvert_authz_RoleBindingStatus_To_v1_RoleBindingStatus(in *authz.RoleBindingStatus, out *RoleBindingStatus, s conversion.Scope) error {
	out.Phase = BindingPhase(in.Phase)
	return nil
}

// Convert_authz_RoleBindingStatus_To_v1_RoleBindingStatus is an autogenerated conversion function.
func Convert_authz_RoleBindingStatus_To_v1_RoleBindingStatus(in *authz.RoleBindingStatus, out *RoleBindingStatus, s conversion.Scope) error {
	return autoConvert_authz_RoleBindingStatus_To_v1_RoleBindingStatus(in, out, s)
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
