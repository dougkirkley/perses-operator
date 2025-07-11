/*
Copyright 2025.

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

package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PersesDashboardStatus defines the observed state of PersesDashboard
type PersesDashboardStatus struct {
	// +operator-sdk:csv:customresourcedefinitions:type=status
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,1,rep,name=conditions"`
}

type PersesDashboardSpec struct {
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	Config Dashboard `json:"config"`
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	// +optional
	InstanceSelector *metav1.LabelSelector `json:"instanceSelector,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:resource:shortName=perdb
//+kubebuilder:conversion:hub
//+versionName=v1alpha2
//+kubebuilder:storageversion

// PersesDashboard is the Schema for the persesdashboards API
type PersesDashboard struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PersesDashboardSpec   `json:"spec,omitempty"`
	Status PersesDashboardStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// PersesDashboardList contains a list of PersesDashboard
type PersesDashboardList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PersesDashboard `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PersesDashboard{}, &PersesDashboardList{})
}
