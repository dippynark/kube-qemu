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

package internalversion

import (
	hypervisor "github.com/dippynark/kube-qemu/pkg/apis/hypervisor"
	scheme "github.com/dippynark/kube-qemu/pkg/client/internalclientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// VirtualMachinesGetter has a method to return a VirtualMachineInterface.
// A group's client should implement this interface.
type VirtualMachinesGetter interface {
	VirtualMachines(namespace string) VirtualMachineInterface
}

// VirtualMachineInterface has methods to work with VirtualMachine resources.
type VirtualMachineInterface interface {
	Create(*hypervisor.VirtualMachine) (*hypervisor.VirtualMachine, error)
	Update(*hypervisor.VirtualMachine) (*hypervisor.VirtualMachine, error)
	UpdateStatus(*hypervisor.VirtualMachine) (*hypervisor.VirtualMachine, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*hypervisor.VirtualMachine, error)
	List(opts v1.ListOptions) (*hypervisor.VirtualMachineList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *hypervisor.VirtualMachine, err error)
	VirtualMachineExpansion
}

// virtualMachines implements VirtualMachineInterface
type virtualMachines struct {
	client rest.Interface
	ns     string
}

// newVirtualMachines returns a VirtualMachines
func newVirtualMachines(c *HypervisorClient, namespace string) *virtualMachines {
	return &virtualMachines{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Create takes the representation of a virtualMachine and creates it.  Returns the server's representation of the virtualMachine, and an error, if there is any.
func (c *virtualMachines) Create(virtualMachine *hypervisor.VirtualMachine) (result *hypervisor.VirtualMachine, err error) {
	result = &hypervisor.VirtualMachine{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("virtualmachines").
		Body(virtualMachine).
		Do().
		Into(result)
	return
}

// Update takes the representation of a virtualMachine and updates it. Returns the server's representation of the virtualMachine, and an error, if there is any.
func (c *virtualMachines) Update(virtualMachine *hypervisor.VirtualMachine) (result *hypervisor.VirtualMachine, err error) {
	result = &hypervisor.VirtualMachine{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("virtualmachines").
		Name(virtualMachine.Name).
		Body(virtualMachine).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclientstatus=false comment above the type to avoid generating UpdateStatus().

func (c *virtualMachines) UpdateStatus(virtualMachine *hypervisor.VirtualMachine) (result *hypervisor.VirtualMachine, err error) {
	result = &hypervisor.VirtualMachine{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("virtualmachines").
		Name(virtualMachine.Name).
		SubResource("status").
		Body(virtualMachine).
		Do().
		Into(result)
	return
}

// Delete takes name of the virtualMachine and deletes it. Returns an error if one occurs.
func (c *virtualMachines) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("virtualmachines").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *virtualMachines) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("virtualmachines").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Get takes name of the virtualMachine, and returns the corresponding virtualMachine object, and an error if there is any.
func (c *virtualMachines) Get(name string, options v1.GetOptions) (result *hypervisor.VirtualMachine, err error) {
	result = &hypervisor.VirtualMachine{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("virtualmachines").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of VirtualMachines that match those selectors.
func (c *virtualMachines) List(opts v1.ListOptions) (result *hypervisor.VirtualMachineList, err error) {
	result = &hypervisor.VirtualMachineList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("virtualmachines").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested virtualMachines.
func (c *virtualMachines) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("virtualmachines").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Patch applies the patch and returns the patched virtualMachine.
func (c *virtualMachines) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *hypervisor.VirtualMachine, err error) {
	result = &hypervisor.VirtualMachine{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("virtualmachines").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
