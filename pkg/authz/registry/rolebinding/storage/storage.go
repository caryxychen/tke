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
	"fmt"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metainternal "k8s.io/apimachinery/pkg/apis/meta/internalversion"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	genericregistry "k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/registry/generic/registry"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/apiserver/pkg/storage"
	storageerr "k8s.io/apiserver/pkg/storage/errors"
	"k8s.io/apiserver/pkg/util/dryrun"
	"tkestack.io/tke/api/authz"
	apiserverutil "tkestack.io/tke/pkg/apiserver/util"
	"tkestack.io/tke/pkg/authz/registry/rolebinding"
	"tkestack.io/tke/pkg/util/log"
)

// Storage includes storage for configmap and all sub resources.
type Storage struct {
	RoleBinding *REST
	Status      *StatusREST
	Finalize    *FinalizeREST
}

// NewStorage returns a Storage object that will work against configmap.
func NewStorage(optsGetter genericregistry.RESTOptionsGetter) *Storage {
	strategy := rolebinding.NewStrategy()
	store := &registry.Store{
		NewFunc:                  func() runtime.Object { return &authz.RoleBinding{} },
		NewListFunc:              func() runtime.Object { return &authz.RoleBindingList{} },
		DefaultQualifiedResource: authz.Resource("rolebindings"),
		ShouldDeleteDuringUpdate: rolebinding.ShouldDeleteDuringUpdate,

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

	statusStore := *store
	statusStore.UpdateStrategy = rolebinding.NewStatusStrategy(strategy)

	finalizeStore := *store
	finalizeStore.UpdateStrategy = rolebinding.NewFinalizerStrategy(strategy)

	return &Storage{
		RoleBinding: &REST{store},
		Status:      &StatusREST{&statusStore},
		Finalize:    &FinalizeREST{&finalizeStore},
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

// Delete enforces life-cycle rules for policy termination
func (r *REST) Delete(ctx context.Context, name string, deleteValidation rest.ValidateObjectFunc, options *metav1.DeleteOptions) (runtime.Object, bool, error) {
	object, err := r.Get(ctx, name, &metav1.GetOptions{})
	if err != nil {
		return nil, false, err
	}
	roleBinding := object.(*authz.RoleBinding)

	// Ensure we have a UID precondition
	if options == nil {
		options = metav1.NewDeleteOptions(0)
	}
	if options.Preconditions == nil {
		options.Preconditions = &metav1.Preconditions{}
	}
	if options.Preconditions.UID == nil {
		options.Preconditions.UID = &roleBinding.UID
	} else if *options.Preconditions.UID != roleBinding.UID {
		err = apierrors.NewConflict(
			authz.Resource("RoleBindings"),
			name,
			fmt.Errorf("precondition failed: UID in precondition: %v, UID in object meta: %v", *options.Preconditions.UID, roleBinding.UID),
		)
		return nil, false, err
	}

	// upon first request to delete, we switch the phase to start roleBinding termination
	if roleBinding.DeletionTimestamp.IsZero() {
		key, err := r.Store.KeyFunc(ctx, name)
		if err != nil {
			return nil, false, err
		}

		preconditions := storage.Preconditions{UID: options.Preconditions.UID}

		out := r.Store.NewFunc()
		err = r.Store.Storage.GuaranteedUpdate(
			ctx, key, out, false, &preconditions,
			storage.SimpleUpdate(func(existing runtime.Object) (runtime.Object, error) {
				existingBinding, ok := existing.(*authz.RoleBinding)
				if !ok {
					// wrong type
					return nil, fmt.Errorf("expected *auth.RoleBinding, got %v", existing)
				}
				if err := deleteValidation(ctx, existingBinding); err != nil {
					return nil, err
				}
				// Set the deletion timestamp if needed
				if existingBinding.DeletionTimestamp.IsZero() {
					now := metav1.Now()
					existingBinding.DeletionTimestamp = &now
				}
				// Set the roleBinding phase to terminating, if needed
				if existingBinding.Status.Phase != authz.BindingTerminating {
					existingBinding.Status.Phase = authz.BindingTerminating
				}

				// the current finalizers which are on namespace
				currentFinalizers := map[string]bool{}
				for _, f := range existingBinding.Finalizers {
					currentFinalizers[f] = true
				}
				// the finalizers we should ensure on rule
				shouldHaveFinalizers := map[string]bool{
					metav1.FinalizerOrphanDependents: apiserverutil.ShouldHaveOrphanFinalizer(options, currentFinalizers[metav1.FinalizerOrphanDependents]),
					metav1.FinalizerDeleteDependents: apiserverutil.ShouldHaveDeleteDependentsFinalizer(options, currentFinalizers[metav1.FinalizerDeleteDependents]),
				}
				// determine whether there are changes
				changeNeeded := false
				for finalizer, shouldHave := range shouldHaveFinalizers {
					changeNeeded = currentFinalizers[finalizer] != shouldHave || changeNeeded
					if shouldHave {
						currentFinalizers[finalizer] = true
					} else {
						delete(currentFinalizers, finalizer)
					}
				}
				// make the changes if needed
				if changeNeeded {
					var newFinalizers []string
					for f := range currentFinalizers {
						newFinalizers = append(newFinalizers, f)
					}
					existingBinding.Finalizers = newFinalizers
				}
				return existingBinding, nil
			}),
			dryrun.IsDryRun(options.DryRun),
			nil,
		)

		if err != nil {
			err = storageerr.InterpretGetError(err, authz.Resource("RoleBindings"), name)
			err = storageerr.InterpretUpdateError(err, authz.Resource("RoleBindings"), name)
			if _, ok := err.(*apierrors.StatusError); !ok {
				err = apierrors.NewInternalError(err)
			}
			return nil, false, err
		}

		return out, false, nil
	}

	// prior to final deletion, we must ensure that finalizers is empty
	if len(roleBinding.Finalizers) != 0 {
		err = apierrors.NewConflict(authz.Resource("RoleBindings"), roleBinding.Name, fmt.Errorf("the system is ensuring all content is removed from this roleBinding.  Upon completion, this roleBinding will automatically be purged by the system"))
		return nil, false, err
	}
	return r.Store.Delete(ctx, name, deleteValidation, options)
}

// StatusREST implements the GenericREST endpoint for changing the status of a policy request.
type StatusREST struct {
	*registry.Store
}

// StatusREST implements Patcher.
var _ = rest.Patcher(&StatusREST{})

// New returns an empty object that can be used with Create and Update after request data has been put into it.
func (r *StatusREST) New() runtime.Object {
	return r.Store.New()
}

// FinalizeREST implements Patcher.
var _ = rest.Patcher(&FinalizeREST{})

// FinalizeREST implements the REST endpoint for finalizing a policy.
type FinalizeREST struct {
	*registry.Store
}

// New returns an empty object that can be used with Create and Update after
// request data has been put into it.
func (r *FinalizeREST) New() runtime.Object {
	return r.Store.New()
}
