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

package multiclusterrolebinding

import (
	"context"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/endpoints/request"
	"k8s.io/apiserver/pkg/registry/generic/registry"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/apiserver/pkg/storage/names"
	"k8s.io/client-go/tools/cache"
	"strings"
	"tkestack.io/tke/api/authz"
	platformversionedclient "tkestack.io/tke/api/client/clientset/versioned/typed/platform/v1"
	"tkestack.io/tke/pkg/apiserver/authentication"
	"tkestack.io/tke/pkg/authz/constant"
	authzprovider "tkestack.io/tke/pkg/authz/provider"
	"tkestack.io/tke/pkg/util/log"
	namesutil "tkestack.io/tke/pkg/util/names"
)

// Strategy implements verification logic for configmap.
type Strategy struct {
	runtime.ObjectTyper
	names.NameGenerator
	roleGetter     rest.Getter
	platformClient platformversionedclient.PlatformV1Interface
}

var _ rest.RESTCreateStrategy = &Strategy{}
var _ rest.RESTUpdateStrategy = &Strategy{}
var _ rest.RESTDeleteStrategy = &Strategy{}

const NamePrefix = "mcrb-"

// NewStrategy creates a strategy that is the default logic that applies when
// creating and updating namespace set objects.
func NewStrategy(roleGetter rest.Getter, platformClient platformversionedclient.PlatformV1Interface) *Strategy {
	return &Strategy{authz.Scheme, namesutil.Generator, roleGetter, platformClient}
}

// DefaultGarbageCollectionPolicy returns the default garbage collection behavior.
func (Strategy) DefaultGarbageCollectionPolicy(ctx context.Context) rest.GarbageCollectionPolicy {
	return rest.Unsupported
}

// NamespaceScoped is false for namespaceSets
func (Strategy) NamespaceScoped() bool {
	return true
}

// Export strips fields that can not be set by the user.
func (Strategy) Export(ctx context.Context, obj runtime.Object, exact bool) error {
	return nil
}

// PrepareForCreate is invoked on create before validation to normalize
// the object.
func (Strategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
	tenantID := request.NamespaceValue(ctx)
	if tenantID == "" {
		tenantID = "default"
	}
	mcrb, _ := obj.(*authz.MultiClusterRoleBinding)
	mcrb.Spec.TenantID = tenantID

	if mcrb.Name == "" && mcrb.GenerateName == "" {
		mcrb.Name = "mcrb-" + mcrb.Spec.Username + "-" + strings.ReplaceAll(mcrb.Spec.RoleName, "/", "-")
		mcrb.GenerateName = NamePrefix
	}
	roleNs, roleName, err := cache.SplitMetaNamespaceKey(mcrb.Spec.RoleName)
	if err != nil {
		return
	}
	labels := mcrb.Labels
	if labels == nil {
		labels = map[string]string{}
	}
	labels[constant.RoleNamespace] = roleNs
	labels[constant.RoleName] = roleName
	labels[constant.Username] = mcrb.Spec.Username
	if dispatchAllClusters(mcrb.Spec.Clusters) {
		labels[constant.DispatchAllClusters] = "true"
	}
	mcrb.Labels = labels

	annotation := mcrb.Annotations
	if annotation == nil {
		annotation = map[string]string{}
	}
	region := authentication.GetExtraValue("region", ctx)
	log.Debugf("region '%v'", region)
	if len(region) != 0 {
		annotation[authz.GroupName+"/region"] = region[0]
	}
	mcrb.Annotations = annotation
	mcrb.Status.Phase = authz.BindingActive
	mcrb.ObjectMeta.Finalizers = []string{string(authz.MultiClusterRoleBindingFinalize)}
}

func dispatchAllClusters(clusterIDs []string) bool {
	if len(clusterIDs) == 1 && clusterIDs[0] == "*" {
		return true
	}
	return false
}

// PrepareForUpdate is invoked on update before validation to normalize the
// object.
func (Strategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
	oldMcrb := old.(*authz.MultiClusterRoleBinding)
	mcrb := obj.(*authz.MultiClusterRoleBinding)
	if dispatchAllClusters(mcrb.Spec.Clusters) {
		mcrb.Labels[constant.DispatchAllClusters] = "true"
	} else {
		delete(mcrb.Labels, constant.DispatchAllClusters)
	}
	if mcrb.Spec.TenantID != oldMcrb.Spec.TenantID {
		log.Warnf("Unauthorized update mcrb tenantID '%s'", oldMcrb.Spec.TenantID)
		mcrb.Spec.TenantID = oldMcrb.Spec.TenantID
	}
	if mcrb.Spec.RoleName != oldMcrb.Spec.RoleName {
		log.Warnf("Unauthorized update mcrb roleName '%s'", oldMcrb.Spec.RoleName)
		mcrb.Spec.RoleName = oldMcrb.Spec.RoleName
	}
}

// Validate validates a new configmap.
func (s Strategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	mcrb := obj.(*authz.MultiClusterRoleBinding)
	provider, err := authzprovider.GetProvider(mcrb.Annotations)
	if err == nil {
		if fieldErr := provider.Validate(context.TODO(), mcrb, s.platformClient); fieldErr != nil {
			return field.ErrorList{fieldErr}
		}
	}
	if len(mcrb.Spec.Clusters) == 0 {
		return field.ErrorList{field.Required(field.NewPath("spec", "clusters"), "empty clusters")}
	}
	return ValidateMultiClusterRoleBinding(mcrb, s.roleGetter, s.platformClient)
}

// AllowCreateOnUpdate is false for persistent events
func (Strategy) AllowCreateOnUpdate() bool {
	return false
}

// AllowUnconditionalUpdate returns true if the object can be updated
// unconditionally (irrespective of the latest resource version), when there is
// no resource version specified in the object.
func (Strategy) AllowUnconditionalUpdate() bool {
	return false
}

// WarningsOnCreate returns warnings for the creation of the given object.
func (Strategy) WarningsOnCreate(ctx context.Context, obj runtime.Object) []string {
	return nil
}

// Canonicalize normalizes the object after validation.
func (Strategy) Canonicalize(obj runtime.Object) {
}

// ValidateUpdate is the default update validation for an end namespace set.
func (s Strategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	return ValidateMultiClusterRoleBindingUpdate(obj.(*authz.MultiClusterRoleBinding), old.(*authz.MultiClusterRoleBinding), s.roleGetter, s.platformClient)
}

// WarningsOnUpdate returns warnings for the given update.
func (Strategy) WarningsOnUpdate(ctx context.Context, obj, old runtime.Object) []string {
	return nil
}

func ShouldDeleteDuringUpdate(ctx context.Context, key string, obj, existing runtime.Object) bool {
	pol, ok := obj.(*authz.MultiClusterRoleBinding)
	if !ok {
		log.Errorf("unexpected object, key:%s", key)
		return false
	}
	return len(pol.Finalizers) == 0 && registry.ShouldDeleteDuringUpdate(ctx, key, obj, existing)
}

// StatusStrategy implements verification logic for status of roletemplate request.
type StatusStrategy struct {
	*Strategy
}

var _ rest.RESTUpdateStrategy = &StatusStrategy{}

// NewStatusStrategy create the StatusStrategy object by given strategy.
func NewStatusStrategy(strategy *Strategy) *StatusStrategy {
	return &StatusStrategy{strategy}
}

// PrepareForUpdate is invoked on update before validation to normalize
// the object.  For example: remove fields that are not to be persisted,
// sort order-insensitive list fields, etc.  This should not remove fields
// whose presence would be considered a validation error.
func (StatusStrategy) PrepareForUpdate(_ context.Context, obj, old runtime.Object) {
	newMultiClusterRoleBinding := obj.(*authz.MultiClusterRoleBinding)
	oldMultiClusterRoleBinding := old.(*authz.MultiClusterRoleBinding)
	status := newMultiClusterRoleBinding.Status
	newMultiClusterRoleBinding = oldMultiClusterRoleBinding
	newMultiClusterRoleBinding.Status = status
}

// ValidateUpdate is invoked after default fields in the object have been
// filled in before the object is persisted.  This method should not mutate
// the object.
func (s *StatusStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	return nil
}

// FinalizeStrategy implements finalizer logic for Machine.
type FinalizeStrategy struct {
	*Strategy
}

var _ rest.RESTUpdateStrategy = &FinalizeStrategy{}

// NewFinalizerStrategy create the FinalizeStrategy object by given strategy.
func NewFinalizerStrategy(strategy *Strategy) *FinalizeStrategy {
	return &FinalizeStrategy{strategy}
}

// PrepareForUpdate is invoked on update before validation to normalize
// the object.  For example: remove fields that are not to be persisted,
// sort order-insensitive list fields, etc.  This should not remove fields
// whose presence would be considered a validation error.
func (FinalizeStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
	newBinding := obj.(*authz.MultiClusterRoleBinding)
	oldBinding := old.(*authz.MultiClusterRoleBinding)
	finalizers := newBinding.Finalizers
	newBinding = oldBinding
	newBinding.Finalizers = finalizers
}

// ValidateUpdate is invoked after default fields in the object have been
// filled in before the object is persisted.  This method should not mutate
// the object.
func (s *FinalizeStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	return nil
}
