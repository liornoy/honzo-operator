/*
Copyright 2021.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// HonzoSpec defines the desired state of Honzo
type HonzoSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// text is the string that honzo prints to the console
	// +kubebuilder:validation:Pattern:="Honzo says:+"
	// +kubebuilder:validation:MinLength:=14
	Text string `json:"text,omitempty"`
	// +kubebuilder:validation:Pattern:="Honzo says:+"
	// +kubebuilder:validation:MinLength:=14
	DeleteText string `json:"deleteText,omitempty"`
}

// HonzoStatus defines the observed state of Honzo
type HonzoStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Honzo is the Schema for the honzoes API
type Honzo struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HonzoSpec   `json:"spec,omitempty"`
	Status HonzoStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// HonzoList contains a list of Honzo
type HonzoList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Honzo `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Honzo{}, &HonzoList{})
}
