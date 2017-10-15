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

package fake

import (
	hypervisor "github.com/dippynark/kube-qemu/pkg/apis/hypervisor"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeVirtualMachines implements VirtualMachineInterface
type FakeVirtualMachines struct {
	Fake *FakeHypervisor
	ns   string
}

var virtualmachinesResource = schema.GroupVersionResource{Group: "hypervisor.lukeaddison.co.uk", Version: "", Resource: "virtualmachines"}

var virtualmachinesKind = schema.GroupVersionKind{Group: "hypervisor.lukeaddison.co.uk", Version: "", Kind: "VirtualMachine"}

func (c *FakeVirtualMachines) Create(virtualMachine *hypervisor.VirtualMachine) (result *hypervisor.VirtualMachine, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(virtualmachinesResource, c.ns, virtualMachine), &hypervisor.VirtualMachine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*hypervisor.VirtualMachine), err
}

func (c *FakeVirtualMachines) Update(virtualMachine *hypervisor.VirtualMachine) (result *hypervisor.VirtualMachine, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(virtualmachinesResource, c.ns, virtualMachine), &hypervisor.VirtualMachine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*hypervisor.VirtualMachine), err
}

func (c *FakeVirtualMachines) UpdateStatus(virtualMachine *hypervisor.VirtualMachine) (*hypervisor.VirtualMachine, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(virtualmachinesResource, "status", c.ns, virtualMachine), &hypervisor.VirtualMachine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*hypervisor.VirtualMachine), err
}

func (c *FakeVirtualMachines) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(virtualmachinesResource, c.ns, name), &hypervisor.VirtualMachine{})

	return err
}

func (c *FakeVirtualMachines) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(virtualmachinesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &hypervisor.VirtualMachineList{})
	return err
}

func (c *FakeVirtualMachines) Get(name string, options v1.GetOptions) (result *hypervisor.VirtualMachine, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(virtualmachinesResource, c.ns, name), &hypervisor.VirtualMachine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*hypervisor.VirtualMachine), err
}

func (c *FakeVirtualMachines) List(opts v1.ListOptions) (result *hypervisor.VirtualMachineList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(virtualmachinesResource, virtualmachinesKind, c.ns, opts), &hypervisor.VirtualMachineList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &hypervisor.VirtualMachineList{}
	for _, item := range obj.(*hypervisor.VirtualMachineList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested virtualMachines.
func (c *FakeVirtualMachines) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(virtualmachinesResource, c.ns, opts))

}

// Patch applies the patch and returns the patched virtualMachine.
func (c *FakeVirtualMachines) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *hypervisor.VirtualMachine, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(virtualmachinesResource, c.ns, name, data, subresources...), &hypervisor.VirtualMachine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*hypervisor.VirtualMachine), err
}
