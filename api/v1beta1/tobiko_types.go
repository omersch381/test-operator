/*
Copyright 2023.

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

package v1beta1

import (
	"github.com/openstack-k8s-operators/lib-common/modules/common/condition"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Hash - struct to add hashes to status
/*
type Hash struct {
	// Name of hash referencing the parameter
	Name string `json:"name,omitempty"`
	// Hash
	Hash string `json:"hash,omitempty"`
}
*/

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// TobikoSpec defines the desired state of Tobiko
type TobikoSpec struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:default="local-storage"
        // StorageClass used to create PVCs that store the logs
	StorageClass string `json:"storageClass"`

	// +kubebuilder:validation:Optional
        // +kubebuilder:default:=true
        // Run tests in parallel
        Debug bool `json:"debug,omitempty"`

	// +kubebuilder:validation:Optional
        // +kubebuilder:default:="py3"
        // Test environment
        Testenv string `json:"testenv,omitempty"`

        // +kubebuilder:validation:Optional
        // +kubebuilder:default:=""
        // Tobiko version
        Version string `json:"version,omitempty"`

        // +kubebuilder:validation:Optional
        // +kubebuilder:default:=""
        // tobiko.conf
        Config string `json:"config,omitempty"`

        // +kubebuilder:validation:Optional
        // +kubebuilder:default:=""
        // Private Key
        PrivateKey string `json:"privateKey,omitempty"`

        // +kubebuilder:validation:Optional
        // +kubebuilder:default:=""
        // Public Key
        PublicKey string `json:"publicKey,omitempty"`

	// +kubebuilder:validation:Optional
        // +kubebuilder:default:="quay.io/podified-antelope-centos9/openstack-tobiko:current-podified"
        // Container image for tobiko
        ContainerImage string `json:"containerImage,omitempty"`

        // +kubebuilder:validation:Optional
        // +kubebuilder:default:=false
        // Container image for tobiko
        Parallel bool `json:"parallel,omitempty"`


	// BackoffLimimt allows to define the maximum number of retried executions (defaults to 6).
        // +kubebuilder:default:=0
        // +operator-sdk:csv:customresourcedefinitions:type=spec,xDescriptors={"urn:alm:descriptor:com.tectonic.ui:number"}
        BackoffLimit *int32 `json:"backoffLimit,omitempty"`
}

// TobikoStatus defines the observed state of Tobiko
type TobikoStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Map of hashes to track e.g. job status
	Hash map[string]string `json:"hash,omitempty"`

	// Conditions
	Conditions condition.Conditions `json:"conditions,omitempty" optional:"true"`

	// NetworkAttachments status of the deployment pods
	NetworkAttachments map[string][]string `json:"networkAttachments,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Tobiko is the Schema for the tobikoes API
type Tobiko struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TobikoSpec   `json:"spec,omitempty"`
	Status TobikoStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// TobikoList contains a list of Tobiko
type TobikoList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Tobiko `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Tobiko{}, &TobikoList{})
}

// RbacConditionsSet - set the conditions for the rbac object
func (instance Tobiko) RbacConditionsSet(c *condition.Condition) {
	instance.Status.Conditions.Set(c)
}

// RbacNamespace - return the namespace
func (instance Tobiko) RbacNamespace() string {
	return instance.Namespace
}

// RbacResourceName - return the name to be used for rbac objects (serviceaccount, role, rolebinding)
func (instance Tobiko) RbacResourceName() string {
	return instance.Name
}
