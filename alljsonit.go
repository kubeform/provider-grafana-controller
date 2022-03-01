/*
Copyright AppsCode Inc. and Contributors

Licensed under the AppsCode Community License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Community-1.0.0.md

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by Kubeform. DO NOT EDIT.

package main

import (
	jsoniter "github.com/json-iterator/go"
	"k8s.io/apimachinery/pkg/runtime/schema"
	alertv1alpha1 "kubeform.dev/provider-grafana-api/apis/alert/v1alpha1"
	builtinv1alpha1 "kubeform.dev/provider-grafana-api/apis/builtin/v1alpha1"
	dashboardv1alpha1 "kubeform.dev/provider-grafana-api/apis/dashboard/v1alpha1"
	datav1alpha1 "kubeform.dev/provider-grafana-api/apis/data/v1alpha1"
	folderv1alpha1 "kubeform.dev/provider-grafana-api/apis/folder/v1alpha1"
	organizationv1alpha1 "kubeform.dev/provider-grafana-api/apis/organization/v1alpha1"
	rolev1alpha1 "kubeform.dev/provider-grafana-api/apis/role/v1alpha1"
	syntheticv1alpha1 "kubeform.dev/provider-grafana-api/apis/synthetic/v1alpha1"
	teamv1alpha1 "kubeform.dev/provider-grafana-api/apis/team/v1alpha1"
	userv1alpha1 "kubeform.dev/provider-grafana-api/apis/user/v1alpha1"
	"kubeform.dev/provider-grafana-controller/controllers"
)

type Data struct {
	JsonIt       jsoniter.API
	ResourceType string
}

var allJsonIt = map[schema.GroupVersionResource]Data{
	{
		Group:    "alert.grafana.kubeform.com",
		Version:  "v1alpha1",
		Resource: "notifications",
	}: {
		JsonIt:       controllers.GetJSONItr(alertv1alpha1.GetEncoder(), alertv1alpha1.GetDecoder()),
		ResourceType: "grafana_alert_notification",
	},
	{
		Group:    "builtin.grafana.kubeform.com",
		Version:  "v1alpha1",
		Resource: "roleassignments",
	}: {
		JsonIt:       controllers.GetJSONItr(builtinv1alpha1.GetEncoder(), builtinv1alpha1.GetDecoder()),
		ResourceType: "grafana_builtin_role_assignment",
	},
	{
		Group:    "dashboard.grafana.kubeform.com",
		Version:  "v1alpha1",
		Resource: "dashboards",
	}: {
		JsonIt:       controllers.GetJSONItr(dashboardv1alpha1.GetEncoder(), dashboardv1alpha1.GetDecoder()),
		ResourceType: "grafana_dashboard",
	},
	{
		Group:    "dashboard.grafana.kubeform.com",
		Version:  "v1alpha1",
		Resource: "permissions",
	}: {
		JsonIt:       controllers.GetJSONItr(dashboardv1alpha1.GetEncoder(), dashboardv1alpha1.GetDecoder()),
		ResourceType: "grafana_dashboard_permission",
	},
	{
		Group:    "data.grafana.kubeform.com",
		Version:  "v1alpha1",
		Resource: "sources",
	}: {
		JsonIt:       controllers.GetJSONItr(datav1alpha1.GetEncoder(), datav1alpha1.GetDecoder()),
		ResourceType: "grafana_data_source",
	},
	{
		Group:    "folder.grafana.kubeform.com",
		Version:  "v1alpha1",
		Resource: "folders",
	}: {
		JsonIt:       controllers.GetJSONItr(folderv1alpha1.GetEncoder(), folderv1alpha1.GetDecoder()),
		ResourceType: "grafana_folder",
	},
	{
		Group:    "folder.grafana.kubeform.com",
		Version:  "v1alpha1",
		Resource: "permissions",
	}: {
		JsonIt:       controllers.GetJSONItr(folderv1alpha1.GetEncoder(), folderv1alpha1.GetDecoder()),
		ResourceType: "grafana_folder_permission",
	},
	{
		Group:    "organization.grafana.kubeform.com",
		Version:  "v1alpha1",
		Resource: "organizations",
	}: {
		JsonIt:       controllers.GetJSONItr(organizationv1alpha1.GetEncoder(), organizationv1alpha1.GetDecoder()),
		ResourceType: "grafana_organization",
	},
	{
		Group:    "role.grafana.kubeform.com",
		Version:  "v1alpha1",
		Resource: "roles",
	}: {
		JsonIt:       controllers.GetJSONItr(rolev1alpha1.GetEncoder(), rolev1alpha1.GetDecoder()),
		ResourceType: "grafana_role",
	},
	{
		Group:    "synthetic.grafana.kubeform.com",
		Version:  "v1alpha1",
		Resource: "monitoringchecks",
	}: {
		JsonIt:       controllers.GetJSONItr(syntheticv1alpha1.GetEncoder(), syntheticv1alpha1.GetDecoder()),
		ResourceType: "grafana_synthetic_monitoring_check",
	},
	{
		Group:    "synthetic.grafana.kubeform.com",
		Version:  "v1alpha1",
		Resource: "monitoringprobes",
	}: {
		JsonIt:       controllers.GetJSONItr(syntheticv1alpha1.GetEncoder(), syntheticv1alpha1.GetDecoder()),
		ResourceType: "grafana_synthetic_monitoring_probe",
	},
	{
		Group:    "team.grafana.kubeform.com",
		Version:  "v1alpha1",
		Resource: "teams",
	}: {
		JsonIt:       controllers.GetJSONItr(teamv1alpha1.GetEncoder(), teamv1alpha1.GetDecoder()),
		ResourceType: "grafana_team",
	},
	{
		Group:    "team.grafana.kubeform.com",
		Version:  "v1alpha1",
		Resource: "externalgroups",
	}: {
		JsonIt:       controllers.GetJSONItr(teamv1alpha1.GetEncoder(), teamv1alpha1.GetDecoder()),
		ResourceType: "grafana_team_external_group",
	},
	{
		Group:    "team.grafana.kubeform.com",
		Version:  "v1alpha1",
		Resource: "preferences",
	}: {
		JsonIt:       controllers.GetJSONItr(teamv1alpha1.GetEncoder(), teamv1alpha1.GetDecoder()),
		ResourceType: "grafana_team_preferences",
	},
	{
		Group:    "user.grafana.kubeform.com",
		Version:  "v1alpha1",
		Resource: "users",
	}: {
		JsonIt:       controllers.GetJSONItr(userv1alpha1.GetEncoder(), userv1alpha1.GetDecoder()),
		ResourceType: "grafana_user",
	},
}

func getJsonItAndResType(gvr schema.GroupVersionResource) Data {
	return allJsonIt[gvr]
}
