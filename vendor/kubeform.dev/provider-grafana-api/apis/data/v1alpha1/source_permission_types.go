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

type SourcePermission struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SourcePermissionSpec   `json:"spec,omitempty"`
	Status            SourcePermissionStatus `json:"status,omitempty"`
}

type SourcePermissionSpecPermissions struct {
	// Permission to associate with item. Must be `Query`.
	Permission *string `json:"permission" tf:"permission"`
	// ID of the team to manage permissions for.
	// +optional
	TeamID *int64 `json:"teamID,omitempty" tf:"team_id"`
	// ID of the user to manage permissions for.
	// +optional
	UserID *int64 `json:"userID,omitempty" tf:"user_id"`
}

type SourcePermissionSpec struct {
	State *SourcePermissionSpecResource `json:"state,omitempty" tf:"-"`

	Resource SourcePermissionSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type SourcePermissionSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// ID of the datasource to apply permissions to.
	DatasourceID *int64 `json:"datasourceID" tf:"datasource_id"`
	// The permission items to add/update. Items that are omitted from the list will be removed.
	Permissions []SourcePermissionSpecPermissions `json:"permissions" tf:"permissions"`
}

type SourcePermissionStatus struct {
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

// SourcePermissionList is a list of SourcePermissions
type SourcePermissionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SourcePermission CRD objects
	Items []SourcePermission `json:"items,omitempty"`
}