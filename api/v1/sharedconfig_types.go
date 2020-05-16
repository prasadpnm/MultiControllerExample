/*


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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// SharedConfigSpec defines the desired state of SharedConfig
type SharedConfigSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of SharedConfig. Edit SharedConfig_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// SharedConfigStatus defines the observed state of SharedConfig
type SharedConfigStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true

// SharedConfig is the Schema for the sharedconfigs API
type SharedConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SharedConfigSpec   `json:"spec,omitempty"`
	Status SharedConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SharedConfigList contains a list of SharedConfig
type SharedConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SharedConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SharedConfig{}, &SharedConfigList{})
}
