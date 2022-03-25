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
	if err := s.AddGeneratedConversionFunc((*PolicyRule)(nil), (*authz.PolicyRule)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_PolicyRule_To_authz_PolicyRule(a.(*PolicyRule), b.(*authz.PolicyRule), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*authz.PolicyRule)(nil), (*PolicyRule)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_authz_PolicyRule_To_v1_PolicyRule(a.(*authz.PolicyRule), b.(*PolicyRule), scope)
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
	return nil
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

func autoConvert_v1_PolicyRule_To_authz_PolicyRule(in *PolicyRule, out *authz.PolicyRule, s conversion.Scope) error {
	out.Verbs = *(*[]string)(unsafe.Pointer(&in.Verbs))
	out.APIGroups = *(*[]string)(unsafe.Pointer(&in.APIGroups))
	out.Resources = *(*[]string)(unsafe.Pointer(&in.Resources))
	out.ResourceNames = *(*[]string)(unsafe.Pointer(&in.ResourceNames))
	out.NonResourceURLs = *(*[]string)(unsafe.Pointer(&in.NonResourceURLs))
	return nil
}

// Convert_v1_PolicyRule_To_authz_PolicyRule is an autogenerated conversion function.
func Convert_v1_PolicyRule_To_authz_PolicyRule(in *PolicyRule, out *authz.PolicyRule, s conversion.Scope) error {
	return autoConvert_v1_PolicyRule_To_authz_PolicyRule(in, out, s)
}

func autoConvert_authz_PolicyRule_To_v1_PolicyRule(in *authz.PolicyRule, out *PolicyRule, s conversion.Scope) error {
	out.Verbs = *(*[]string)(unsafe.Pointer(&in.Verbs))
	out.APIGroups = *(*[]string)(unsafe.Pointer(&in.APIGroups))
	out.Resources = *(*[]string)(unsafe.Pointer(&in.Resources))
	out.ResourceNames = *(*[]string)(unsafe.Pointer(&in.ResourceNames))
	out.NonResourceURLs = *(*[]string)(unsafe.Pointer(&in.NonResourceURLs))
	return nil
}

// Convert_authz_PolicyRule_To_v1_PolicyRule is an autogenerated conversion function.
func Convert_authz_PolicyRule_To_v1_PolicyRule(in *authz.PolicyRule, out *PolicyRule, s conversion.Scope) error {
	return autoConvert_authz_PolicyRule_To_v1_PolicyRule(in, out, s)
}

func autoConvert_v1_RoleTemplate_To_authz_RoleTemplate(in *RoleTemplate, out *authz.RoleTemplate, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_RoleTemplateSpec_To_authz_RoleTemplateSpec(&in.Spec, &out.Spec, s); err != nil {
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
	out.Rules = *(*[]authz.PolicyRule)(unsafe.Pointer(&in.Rules))
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
	out.Rules = *(*[]PolicyRule)(unsafe.Pointer(&in.Rules))
	return nil
}

// Convert_authz_RoleTemplateSpec_To_v1_RoleTemplateSpec is an autogenerated conversion function.
func Convert_authz_RoleTemplateSpec_To_v1_RoleTemplateSpec(in *authz.RoleTemplateSpec, out *RoleTemplateSpec, s conversion.Scope) error {
	return autoConvert_authz_RoleTemplateSpec_To_v1_RoleTemplateSpec(in, out, s)
}
