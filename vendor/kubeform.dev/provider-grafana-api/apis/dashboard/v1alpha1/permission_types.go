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

type Permission struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PermissionSpec   `json:"spec,omitempty"`
	Status            PermissionStatus `json:"status,omitempty"`
}

type PermissionSpecPermissions struct {
	// Permission to associate with item. Must be one of `View`, `Edit`, or `Admin`.
	Permission *string `json:"permission" tf:"permission"`
	// Manage permissions for `Viewer` or `Editor` roles.
	// +optional
	Role *string `json:"role,omitempty" tf:"role"`
	// ID of the team to manage permissions for.
	// +optional
	TeamID *int64 `json:"teamID,omitempty" tf:"team_id"`
	// ID of the user to manage permissions for.
	// +optional
	UserID *int64 `json:"userID,omitempty" tf:"user_id"`
}

type PermissionSpec struct {
	State *PermissionSpecResource `json:"state,omitempty" tf:"-"`

	Resource PermissionSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type PermissionSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// ID of the dashboard to apply permissions to.
	DashboardID *int64 `json:"dashboardID" tf:"dashboard_id"`
	// The permission items to add/update. Items that are omitted from the list will be removed.
	Permissions []PermissionSpecPermissions `json:"permissions" tf:"permissions"`
}

type PermissionStatus struct {
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

// PermissionList is a list of Permissions
type PermissionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Permission CRD objects
	Items []Permission `json:"items,omitempty"`
}
