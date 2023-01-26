/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ProcessorObservation struct {

	// an identifier for the resource with format projects/{{project}}/locations/{{location}}/processors/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The resource name of the processor.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type ProcessorParameters struct {

	// The display name. Must be unique.
	// +kubebuilder:validation:Required
	DisplayName *string `json:"displayName" tf:"display_name,omitempty"`

	// The KMS key used for encryption/decryption in CMEK scenarios. See https://cloud.google.com/security-key-management.
	// +kubebuilder:validation:Optional
	KMSKeyName *string `json:"kmsKeyName,omitempty" tf:"kms_key_name,omitempty"`

	// The location of the resource.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The type of processor. For possible types see the official list
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

// ProcessorSpec defines the desired state of Processor
type ProcessorSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProcessorParameters `json:"forProvider"`
}

// ProcessorStatus defines the observed state of Processor.
type ProcessorStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProcessorObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Processor is the Schema for the Processors API. The first-class citizen for Document AI.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Processor struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProcessorSpec   `json:"spec"`
	Status            ProcessorStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProcessorList contains a list of Processors
type ProcessorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Processor `json:"items"`
}

// Repository type metadata.
var (
	Processor_Kind             = "Processor"
	Processor_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Processor_Kind}.String()
	Processor_KindAPIVersion   = Processor_Kind + "." + CRDGroupVersion.String()
	Processor_GroupVersionKind = CRDGroupVersion.WithKind(Processor_Kind)
)

func init() {
	SchemeBuilder.Register(&Processor{}, &ProcessorList{})
}
