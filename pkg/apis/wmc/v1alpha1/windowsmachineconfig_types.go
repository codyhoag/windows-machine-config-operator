package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// WindowsMachineConfigSpec defines the desired state of WindowsMachineConfig
type WindowsMachineConfigSpec struct {
	// Replicas represent how many Windows nodes to be added to the
	// OpenShift cluster
	// +kubebuilder:validation:Minimum=0
	Replicas int `json:"replicas"`
	// InstanceType represents the flavor of instance to be used while
	// creating the virtual machines. Please note that this is common
	// across all the Windows nodes in the cluster
	InstanceType string `json:"instanceType"`
	// AWS holds AWS specific cloud provider information
	// +optional
	AWS *AWS `json:"aws,omitempty"`
	// Azure holds Azure specific cloud provider information
	// +optional
	Azure *Azure `json:"azure,omitempty"`
}

// AWS holds the information related to AWS cloud provider
type AWS struct {
	// SSHKeyPair is the sshKeyPair associated with cloudprovider. AWS
	// asks a keypair to be present for encrypting the Windows VM password
	SSHKeyPair string `json:"sshKeyPair"`
	// CredentialAccountID is account id associated with AWS provider
	CredentialAccountID string `json:"credentialAccountId"`
}

// AzureProvider holds the information related to Azure cloud provider.
// TODO: Populate when we are working on azure.
type Azure struct {
}

// WindowsMachineConfigStatus defines the observed state of WindowsMachineConfig
type WindowsMachineConfigStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// WindowsMachineConfig is the Schema for the windowsmachineconfigs API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=windowsmachineconfigs,scope=Namespaced
type WindowsMachineConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WindowsMachineConfigSpec   `json:"spec,omitempty"`
	Status WindowsMachineConfigStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// WindowsMachineConfigList contains a list of WindowsMachineConfig
type WindowsMachineConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WindowsMachineConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&WindowsMachineConfig{}, &WindowsMachineConfigList{})
}
