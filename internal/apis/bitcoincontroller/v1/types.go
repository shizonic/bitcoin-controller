package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen=true

// BitcoinNetwork specifies a bitcoin test network
type BitcoinNetwork struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BitcoinNetworkSpec   `json:"spec"`
	Status BitcoinNetworkStatus `json:"status"`
}

// BitcoinNetworkSpec is the to-be state of a bitcoin network
type BitcoinNetworkSpec struct {
	Nodes int `json:",nodes"`
}

// BitcoinNetworkStatus is the as-is state of a bitcoin network
type BitcoinNetworkStatus struct {
	Nodes int `json:",nodes"`
}

// +k8s:deepcopy-gen=true

// BitcoinNetworkList is a list of bitcoin networks
type BitcoinNetworkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []BitcoinNetwork `json:"items"`
}