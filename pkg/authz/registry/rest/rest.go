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

package rest

import (
	"k8s.io/apiserver/pkg/authorization/authorizer"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/registry/rest"
	genericserver "k8s.io/apiserver/pkg/server"
	serverstorage "k8s.io/apiserver/pkg/server/storage"
	restclient "k8s.io/client-go/rest"
	"tkestack.io/tke/api/authz"
	authzv1 "tkestack.io/tke/api/authz/v1"
	"tkestack.io/tke/pkg/apiserver/storage"
	roletemplatestorage "tkestack.io/tke/pkg/authz/registry/roletemplate/storage"
)

// StorageProvider is a REST type for core resources storage that implement
// RestStorageProvider interface
type StorageProvider struct {
	LoopbackClientConfig *restclient.Config
	Authorizer           authorizer.Authorizer
}

// Implement RESTStorageProvider
var _ storage.RESTStorageProvider = &StorageProvider{}

// NewRESTStorage is a factory constructor to creates and returns the APIGroupInfo
func (s *StorageProvider) NewRESTStorage(apiResourceConfigSource serverstorage.APIResourceConfigSource, restOptionsGetter generic.RESTOptionsGetter) (genericserver.APIGroupInfo, bool) {
	apiGroupInfo := genericserver.NewDefaultAPIGroupInfo(authz.GroupName, authz.Scheme, authz.ParameterCodec, authz.Codecs)

	if apiResourceConfigSource.VersionEnabled(authzv1.SchemeGroupVersion) {
		apiGroupInfo.VersionedResourcesStorageMap[authzv1.SchemeGroupVersion.Version] =
			s.v1Storage(apiResourceConfigSource, restOptionsGetter)
	}
	return apiGroupInfo, true
}

// GroupName return the api group name
func (*StorageProvider) GroupName() string {
	return authz.GroupName
}

func (s *StorageProvider) v1Storage(apiResourceConfigSource serverstorage.APIResourceConfigSource, restOptionsGetter generic.RESTOptionsGetter) map[string]rest.Storage {
	storageMap := make(map[string]rest.Storage)
	{
		roletemplateREST := roletemplatestorage.NewStorage(restOptionsGetter)
		storageMap["roletemplates"] = roletemplateREST.RoleTemplate
	}
	return storageMap
}
