// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors.

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

// This file was autogenerated by conversion-gen. Do not edit it manually!

package v1alpha1

import (
	hypervisor "github.com/dippynark/kube-qemu/pkg/apis/hypervisor"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	unsafe "unsafe"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1alpha1_VirtualMachine_To_hypervisor_VirtualMachine,
		Convert_hypervisor_VirtualMachine_To_v1alpha1_VirtualMachine,
		Convert_v1alpha1_VirtualMachineList_To_hypervisor_VirtualMachineList,
		Convert_hypervisor_VirtualMachineList_To_v1alpha1_VirtualMachineList,
		Convert_v1alpha1_VirtualMachineSpec_To_hypervisor_VirtualMachineSpec,
		Convert_hypervisor_VirtualMachineSpec_To_v1alpha1_VirtualMachineSpec,
		Convert_v1alpha1_VirtualMachineStatus_To_hypervisor_VirtualMachineStatus,
		Convert_hypervisor_VirtualMachineStatus_To_v1alpha1_VirtualMachineStatus,
	)
}

func autoConvert_v1alpha1_VirtualMachine_To_hypervisor_VirtualMachine(in *VirtualMachine, out *hypervisor.VirtualMachine, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_VirtualMachineSpec_To_hypervisor_VirtualMachineSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_VirtualMachineStatus_To_hypervisor_VirtualMachineStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_VirtualMachine_To_hypervisor_VirtualMachine is an autogenerated conversion function.
func Convert_v1alpha1_VirtualMachine_To_hypervisor_VirtualMachine(in *VirtualMachine, out *hypervisor.VirtualMachine, s conversion.Scope) error {
	return autoConvert_v1alpha1_VirtualMachine_To_hypervisor_VirtualMachine(in, out, s)
}

func autoConvert_hypervisor_VirtualMachine_To_v1alpha1_VirtualMachine(in *hypervisor.VirtualMachine, out *VirtualMachine, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_hypervisor_VirtualMachineSpec_To_v1alpha1_VirtualMachineSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_hypervisor_VirtualMachineStatus_To_v1alpha1_VirtualMachineStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_hypervisor_VirtualMachine_To_v1alpha1_VirtualMachine is an autogenerated conversion function.
func Convert_hypervisor_VirtualMachine_To_v1alpha1_VirtualMachine(in *hypervisor.VirtualMachine, out *VirtualMachine, s conversion.Scope) error {
	return autoConvert_hypervisor_VirtualMachine_To_v1alpha1_VirtualMachine(in, out, s)
}

func autoConvert_v1alpha1_VirtualMachineList_To_hypervisor_VirtualMachineList(in *VirtualMachineList, out *hypervisor.VirtualMachineList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]hypervisor.VirtualMachine)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_VirtualMachineList_To_hypervisor_VirtualMachineList is an autogenerated conversion function.
func Convert_v1alpha1_VirtualMachineList_To_hypervisor_VirtualMachineList(in *VirtualMachineList, out *hypervisor.VirtualMachineList, s conversion.Scope) error {
	return autoConvert_v1alpha1_VirtualMachineList_To_hypervisor_VirtualMachineList(in, out, s)
}

func autoConvert_hypervisor_VirtualMachineList_To_v1alpha1_VirtualMachineList(in *hypervisor.VirtualMachineList, out *VirtualMachineList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items == nil {
		out.Items = make([]VirtualMachine, 0)
	} else {
		out.Items = *(*[]VirtualMachine)(unsafe.Pointer(&in.Items))
	}
	return nil
}

// Convert_hypervisor_VirtualMachineList_To_v1alpha1_VirtualMachineList is an autogenerated conversion function.
func Convert_hypervisor_VirtualMachineList_To_v1alpha1_VirtualMachineList(in *hypervisor.VirtualMachineList, out *VirtualMachineList, s conversion.Scope) error {
	return autoConvert_hypervisor_VirtualMachineList_To_v1alpha1_VirtualMachineList(in, out, s)
}

func autoConvert_v1alpha1_VirtualMachineSpec_To_hypervisor_VirtualMachineSpec(in *VirtualMachineSpec, out *hypervisor.VirtualMachineSpec, s conversion.Scope) error {
	return nil
}

// Convert_v1alpha1_VirtualMachineSpec_To_hypervisor_VirtualMachineSpec is an autogenerated conversion function.
func Convert_v1alpha1_VirtualMachineSpec_To_hypervisor_VirtualMachineSpec(in *VirtualMachineSpec, out *hypervisor.VirtualMachineSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_VirtualMachineSpec_To_hypervisor_VirtualMachineSpec(in, out, s)
}

func autoConvert_hypervisor_VirtualMachineSpec_To_v1alpha1_VirtualMachineSpec(in *hypervisor.VirtualMachineSpec, out *VirtualMachineSpec, s conversion.Scope) error {
	return nil
}

// Convert_hypervisor_VirtualMachineSpec_To_v1alpha1_VirtualMachineSpec is an autogenerated conversion function.
func Convert_hypervisor_VirtualMachineSpec_To_v1alpha1_VirtualMachineSpec(in *hypervisor.VirtualMachineSpec, out *VirtualMachineSpec, s conversion.Scope) error {
	return autoConvert_hypervisor_VirtualMachineSpec_To_v1alpha1_VirtualMachineSpec(in, out, s)
}

func autoConvert_v1alpha1_VirtualMachineStatus_To_hypervisor_VirtualMachineStatus(in *VirtualMachineStatus, out *hypervisor.VirtualMachineStatus, s conversion.Scope) error {
	return nil
}

// Convert_v1alpha1_VirtualMachineStatus_To_hypervisor_VirtualMachineStatus is an autogenerated conversion function.
func Convert_v1alpha1_VirtualMachineStatus_To_hypervisor_VirtualMachineStatus(in *VirtualMachineStatus, out *hypervisor.VirtualMachineStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_VirtualMachineStatus_To_hypervisor_VirtualMachineStatus(in, out, s)
}

func autoConvert_hypervisor_VirtualMachineStatus_To_v1alpha1_VirtualMachineStatus(in *hypervisor.VirtualMachineStatus, out *VirtualMachineStatus, s conversion.Scope) error {
	return nil
}

// Convert_hypervisor_VirtualMachineStatus_To_v1alpha1_VirtualMachineStatus is an autogenerated conversion function.
func Convert_hypervisor_VirtualMachineStatus_To_v1alpha1_VirtualMachineStatus(in *hypervisor.VirtualMachineStatus, out *VirtualMachineStatus, s conversion.Scope) error {
	return autoConvert_hypervisor_VirtualMachineStatus_To_v1alpha1_VirtualMachineStatus(in, out, s)
}
