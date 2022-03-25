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
	authzv1 "tkestack.io/tke/api/authz/v1"
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
		5,
	)
	go ctrl.Run(4, ctx.Stop)
	return nil, true, nil
}
