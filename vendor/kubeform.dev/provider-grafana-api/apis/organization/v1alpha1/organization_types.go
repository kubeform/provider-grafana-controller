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

type Organization struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OrganizationSpec   `json:"spec,omitempty"`
	Status            OrganizationStatus `json:"status,omitempty"`
}

type OrganizationSpec struct {
	State *OrganizationSpecResource `json:"state,omitempty" tf:"-"`

	Resource OrganizationSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type OrganizationSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	//
	// The login name of the configured default admin user for the Grafana
	// installation. If unset, this value defaults to admin, the Grafana default.
	// Grafana adds the default admin user to all organizations automatically upon
	// creation, and this parameter keeps Terraform from removing it from
	// organizations.
	//
	// +optional
	AdminUser *string `json:"adminUser,omitempty" tf:"admin_user"`
	//
	// A list of email addresses corresponding to users who should be given admin
	// access to the organization. Note: users specified here must already exist in
	// Grafana unless 'create_users' is set to true.
	//
	// +optional
	Admins []string `json:"admins,omitempty" tf:"admins"`
	//
	// Whether or not to create Grafana users specified in the organization's
	// membership if they don't already exist in Grafana. If unspecified, this
	// parameter defaults to true, creating placeholder users with the name, login,
	// and email set to the email of the user, and a random password. Setting this
	// option to false will cause an error to be thrown for any users that do not
	// already exist in Grafana.
	//
	// +optional
	CreateUsers *bool `json:"createUsers,omitempty" tf:"create_users"`
	//
	// A list of email addresses corresponding to users who should be given editor
	// access to the organization. Note: users specified here must already exist in
	// Grafana unless 'create_users' is set to true.
	//
	// +optional
	Editors []string `json:"editors,omitempty" tf:"editors"`
	// The display name for the Grafana organization created.
	Name *string `json:"name" tf:"name"`
	// The organization id assigned to this organization by Grafana.
	// +optional
	OrgID *int64 `json:"orgID,omitempty" tf:"org_id"`
	//
	// A list of email addresses corresponding to users who should be given viewer
	// access to the organization. Note: users specified here must already exist in
	// Grafana unless 'create_users' is set to true.
	//
	// +optional
	Viewers []string `json:"viewers,omitempty" tf:"viewers"`
}

type OrganizationStatus struct {
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

// OrganizationList is a list of Organizations
type OrganizationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Organization CRD objects
	Items []Organization `json:"items,omitempty"`
}
