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

package banvirtualmachine

import (
	"fmt"
	"io"

	"github.com/dippynark/kube-qemu/pkg/admission/hypervisorinitializer"
	informers "github.com/dippynark/kube-qemu/pkg/informers/internalversion"
	listers "github.com/dippynark/kube-qemu/pkg/listers/hypervisor/internalversion"
	"k8s.io/apiserver/pkg/admission"
)

// Register registers a plugin
func Register(plugins *admission.Plugins) {
	plugins.Register("BanVirtualMachine", func(config io.Reader) (admission.Interface, error) {
		return New()
	})
}

type disallowVirtualMachine struct {
	*admission.Handler
	lister listers.VirtualMachineLister
}

var _ = hypervisorinitializer.WantsInternalHypervisorInformerFactory(&disallowVirtualMachine{})

// Admit ensures that the object in-flight is of kind virtualmachine.
// In addition checks that the Name is not on the banned list.
// The list is stored in Fischers API objects.
func (d *disallowVirtualMachine) Admit(a admission.Attributes) error {
	// we are only interested in virtualmachines

	return nil
}

// SetInternalHypervisorInformerFactory gets Lister from SharedInformerFactory.
// The lister knows how to lists Fischers.
func (d *disallowVirtualMachine) SetInternalHypervisorInformerFactory(f informers.SharedInformerFactory) {
	d.lister = f.Hypervisor.InternalVersion().VirtualMachines().Lister()
}

// Validate checks whether the plugin was correctly initialized.
func (d *disallowVirtualMachine) Validate() error {
	if d.lister == nil {
		return fmt.Errorf("missing fischer lister")
	}
	return nil
}

// New creates a new ban virtualmachine admission plugin
func New() (admission.Interface, error) {
	return &disallowVirtualMachine{
		Handler: admission.NewHandler(admission.Create),
	}, nil
}
