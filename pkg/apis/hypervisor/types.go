package hypervisor

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type VirtualMachine struct {
	metav1.TypeMeta
	metav1.ObjectMeta

	Spec   VirtualMachineSpec
	Status VirtualMachineStatus
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type VirtualMachineList struct {
	metav1.TypeMeta
	metav1.ObjectMeta

	Items []VirtualMachine
}

type VirtualMachineSpec struct {
	Name string
}

type VirtualMachineStatus struct {
	Running bool
}
