/*
Copyright 2020 The Kubernetes authors.

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

// K8sResourceClonerSpec defines the desired state of K8sResourceCloner
type K8sResourceClonerSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of K8sResourceCloner. Edit K8sResourceCloner_types.go to remove/update
	FromNamespace string   `json:"fromNamespace"`
	ToNamespaces  []string `json:"toNamespaces"`
	ResourceNames []string `json:"resourceNames"`
}

// K8sResourceClonerStatus defines the observed state of K8sResourceCloner
type K8sResourceClonerStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Nodes []string `json:"nodes"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// K8sResourceCloner is the Schema for the k8sresourcecloners API
type K8sResourceCloner struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   K8sResourceClonerSpec   `json:"spec,omitempty"`
	Status K8sResourceClonerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// K8sResourceClonerList contains a list of K8sResourceCloner
type K8sResourceClonerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []K8sResourceCloner `json:"items"`
}

func init() {
	SchemeBuilder.Register(&K8sResourceCloner{}, &K8sResourceClonerList{})
}
