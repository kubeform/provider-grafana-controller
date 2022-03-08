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

type Report struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ReportSpec   `json:"spec,omitempty"`
	Status            ReportStatus `json:"status,omitempty"`
}

type ReportSpecSchedule struct {
	// Custom interval of the report.
	// **Note:** This field is only available when frequency is set to `custom`.
	// +optional
	CustomInterval *string `json:"customInterval,omitempty" tf:"custom_interval"`
	// End time of the report. If empty, the report will be sent indefinitely (according to frequency). Note that times will be saved as UTC in Grafana.
	// +optional
	EndTime *string `json:"endTime,omitempty" tf:"end_time"`
	// Frequency of the report. One of `never`, `once`, `hourly`, `daily`, `weekly`, `monthly` or `custom`.
	Frequency *string `json:"frequency" tf:"frequency"`
	// Start time of the report. If empty, the start date will be set to the creation time. Note that times will be saved as UTC in Grafana.
	// +optional
	StartTime *string `json:"startTime,omitempty" tf:"start_time"`
	// Whether to send the report only on work days.
	// +optional
	WorkdaysOnly *bool `json:"workdaysOnly,omitempty" tf:"workdays_only"`
}

type ReportSpecTimeRange struct {
	// Start of the time range.
	// +optional
	From *string `json:"from,omitempty" tf:"from"`
	// End of the time range.
	// +optional
	To *string `json:"to,omitempty" tf:"to"`
}

type ReportSpec struct {
	State *ReportSpecResource `json:"state,omitempty" tf:"-"`

	Resource ReportSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type ReportSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Dashboard to be sent in the report.
	DashboardID *int64 `json:"dashboardID" tf:"dashboard_id"`
	// Whether to include a link to the dashboard in the report.
	// +optional
	IncludeDashboardLink *bool `json:"includeDashboardLink,omitempty" tf:"include_dashboard_link"`
	// Whether to include a CSV file of table panel data.
	// +optional
	IncludeTableCsv *bool `json:"includeTableCsv,omitempty" tf:"include_table_csv"`
	// Layout of the report. `simple` or `grid`
	// +optional
	Layout *string `json:"layout,omitempty" tf:"layout"`
	// Message to be sent in the report.
	// +optional
	Message *string `json:"message,omitempty" tf:"message"`
	// Name of the report.
	Name *string `json:"name" tf:"name"`
	// Orientation of the report. `landscape` or `portrait`
	// +optional
	Orientation *string `json:"orientation,omitempty" tf:"orientation"`
	// List of recipients of the report.
	// +kubebuilder:validation:MinItems=1
	Recipients []string `json:"recipients" tf:"recipients"`
	// Reply-to email address of the report.
	// +optional
	ReplyTo *string `json:"replyTo,omitempty" tf:"reply_to"`
	// Schedule of the report.
	Schedule *ReportSpecSchedule `json:"schedule" tf:"schedule"`
	// Time range of the report.
	// +optional
	TimeRange *ReportSpecTimeRange `json:"timeRange,omitempty" tf:"time_range"`
}

type ReportStatus struct {
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

// ReportList is a list of Reports
type ReportList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Report CRD objects
	Items []Report `json:"items,omitempty"`
}