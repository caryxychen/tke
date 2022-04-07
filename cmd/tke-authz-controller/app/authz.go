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

package app

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"net/http"
	"time"
	authzv1 "tkestack.io/tke/api/authz/v1"
	"tkestack.io/tke/pkg/authz/controller/clusterroletemplatebinding"
	"tkestack.io/tke/pkg/authz/controller/rolebinding"
	"tkestack.io/tke/pkg/authz/controller/roletemplate"
)

func startRoleTemplateController(ctx ControllerContext) (http.Handler, bool, error) {
	if !ctx.AvailableResources[schema.GroupVersionResource{Group: authzv1.GroupName, Version: "v1", Resource: "roletemplates"}] {
		return nil, false, nil
	}
	ctrl := roletemplate.NewController(
		ctx.ClientBuilder.ClientOrDie("roletemplate-controller"),
		ctx.PlatformClient,
		ctx.InformerFactory.Authz().V1().RoleTemplates(),
		5 * time.Second,
	)
	go ctrl.Run(4, ctx.Stop)
	return nil, true, nil
}

func startClusterRoleTemplateBindingController(ctx ControllerContext) (http.Handler, bool, error) {
	if !ctx.AvailableResources[schema.GroupVersionResource{Group: authzv1.GroupName, Version: "v1", Resource: "clusterroletemplatebindings"}] {
		return nil, false, nil
	}
	ctrl := clusterroletemplatebinding.NewController(
		ctx.ClientBuilder.ClientOrDie("clusterroletemplatebinding-controller"),
		ctx.PlatformClient,
		ctx.InformerFactory.Authz().V1().RoleTemplates(),
		ctx.InformerFactory.Authz().V1().ClusterRoleTemplateBindings(),
		5 * time.Second,
	)
	go ctrl.Run(4, ctx.Stop)
	return nil, true, nil
}

func startRoleBindingController(ctx ControllerContext) (http.Handler, bool, error) {
	if !ctx.AvailableResources[schema.GroupVersionResource{Group: authzv1.GroupName, Version: "v1", Resource: "rolebindings"}] {
		return nil, false, nil
	}
	ctrl := rolebinding.NewController(
		ctx.ClientBuilder.ClientOrDie("rolebinding-controller"),
		ctx.PlatformClient,
		ctx.InformerFactory.Authz().V1().Roles(),
		ctx.InformerFactory.Authz().V1().RoleBindings(),
		ctx.InformerFactory.Authz().V1().RoleTemplates(),
		5 * time.Second,
	)
	go ctrl.Run(4, ctx.Stop)
	return nil, true, nil
}
