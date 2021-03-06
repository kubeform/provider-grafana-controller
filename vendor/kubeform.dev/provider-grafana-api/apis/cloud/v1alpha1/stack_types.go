/*
Copyright AppsCode Inc. and Contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	base "kubeform.dev/apimachinery/api/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kmapi "kmodules.xyz/client-go/api/v1"
	"sigs.k8s.io/cli-utils/pkg/kstatus/status"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type Stack struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StackSpec   `json:"spec,omitempty"`
	Status            StackStatus `json:"status,omitempty"`
}

type StackSpec struct {
	State *StackSpecResource `json:"state,omitempty" tf:"-"`

	Resource StackSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type StackSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of the Alertmanager instance configured for this stack.
	// +optional
	AlertmanagerName *string `json:"alertmanagerName,omitempty" tf:"alertmanager_name"`
	// Status of the Alertmanager instance configured for this stack.
	// +optional
	AlertmanagerStatus *string `json:"alertmanagerStatus,omitempty" tf:"alertmanager_status"`
	// Base URL of the Alertmanager instance configured for this stack.
	// +optional
	AlertmanagerURL *string `json:"alertmanagerURL,omitempty" tf:"alertmanager_url"`
	// User ID of the Alertmanager instance configured for this stack.
	// +optional
	AlertmanagerUserID *int64 `json:"alertmanagerUserID,omitempty" tf:"alertmanager_user_id"`
	// Description of stack.
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// Name of stack. Conventionally matches the url of the instance (e.g. ???<stack_slug>.grafana.net???).
	Name *string `json:"name" tf:"name"`
	// Organization id to assign to this stack.
	// +optional
	OrgID *int64 `json:"orgID,omitempty" tf:"org_id"`
	// Organization name to assign to this stack.
	// +optional
	OrgName *string `json:"orgName,omitempty" tf:"org_name"`
	// Organization slug to assign to this stack.
	// +optional
	OrgSlug *string `json:"orgSlug,omitempty" tf:"org_slug"`
	// Prometheus name for this instance.
	// +optional
	PrometheusName *string `json:"prometheusName,omitempty" tf:"prometheus_name"`
	// Use this URL to query hosted metrics data e.g. Prometheus data source in Grafana
	// +optional
	PrometheusRemoteEndpoint *string `json:"prometheusRemoteEndpoint,omitempty" tf:"prometheus_remote_endpoint"`
	// Use this URL to send prometheus metrics to Grafana cloud
	// +optional
	PrometheusRemoteWriteEndpoint *string `json:"prometheusRemoteWriteEndpoint,omitempty" tf:"prometheus_remote_write_endpoint"`
	// Prometheus status for this instance.
	// +optional
	PrometheusStatus *string `json:"prometheusStatus,omitempty" tf:"prometheus_status"`
	// Prometheus url for this instance.
	// +optional
	PrometheusURL *string `json:"prometheusURL,omitempty" tf:"prometheus_url"`
	// Promehteus user ID. Used for e.g. remote_write.
	// +optional
	PrometheusUserID *int64 `json:"prometheusUserID,omitempty" tf:"prometheus_user_id"`
	// Region slug to assign to this stack.
	// Changing region will destroy the existing stack and create a new one in the desired region
	// +optional
	RegionSlug *string `json:"regionSlug,omitempty" tf:"region_slug"`
	//
	// Subdomain that the Grafana instance will be available at (i.e. setting slug to ???<stack_slug>??? will make the instance
	// available at ???https://<stack_slug>.grafana.net".
	Slug *string `json:"slug" tf:"slug"`
	// Status of the stack.
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	// Custom URL for the Grafana instance. Must have a CNAME setup to point to `.grafana.net` before creating the stack
	// +optional
	Url *string `json:"url,omitempty" tf:"url"`
}

type StackStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Phase status.Status `json:"phase,omitempty"`
	// +optional
	Conditions []kmapi.Condition `json:"conditions,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// StackList is a list of Stacks
type StackList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Stack CRD objects
	Items []Stack `json:"items,omitempty"`
}
