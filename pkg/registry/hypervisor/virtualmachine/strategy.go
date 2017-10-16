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

package virtualmachine

import (
	"fmt"

	"github.com/dippynark/kube-qemu/pkg/apis/hypervisor"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/storage"
	"k8s.io/apiserver/pkg/storage/names"

	genericapirequest "k8s.io/apiserver/pkg/endpoints/request"
)

func NewStrategy(typer runtime.ObjectTyper) virtualMachineStrategy {
	return virtualMachineStrategy{typer, names.SimpleNameGenerator}
}

func GetAttrs(obj runtime.Object) (labels.Set, fields.Set, bool, error) {
	apiserver, ok := obj.(*hypervisor.VirtualMachine)
	if !ok {
		return nil, nil, false, fmt.Errorf("given object is not a VirtualMachine.")
	}
	return labels.Set(apiserver.ObjectMeta.Labels), virtualMachineToSelectableFields(apiserver), apiserver.Initializers != nil, nil
}

// MatchFischer is the filter used by the generic etcd backend to watch events
// from etcd to clients of the apiserver only interested in specific labels/fields.
func MatchVirtualMachine(label labels.Selector, field fields.Selector) storage.SelectionPredicate {
	return storage.SelectionPredicate{
		Label:    label,
		Field:    field,
		GetAttrs: GetAttrs,
	}
}

// virtualMachineToSelectableFields returns a field set that represents the object.
func virtualMachineToSelectableFields(obj *hypervisor.VirtualMachine) fields.Set {
	return generic.ObjectMetaFieldsSet(&obj.ObjectMeta, true)
}

type virtualMachineStrategy struct {
	runtime.ObjectTyper
	names.NameGenerator
}

func (virtualMachineStrategy) NamespaceScoped() bool {
	return true
}

func (virtualMachineStrategy) PrepareForCreate(ctx genericapirequest.Context, obj runtime.Object) {
}

func (virtualMachineStrategy) PrepareForUpdate(ctx genericapirequest.Context, obj, old runtime.Object) {
}

func (virtualMachineStrategy) Validate(ctx genericapirequest.Context, obj runtime.Object) field.ErrorList {
	return field.ErrorList{}
}

func (virtualMachineStrategy) AllowCreateOnUpdate() bool {
	return false
}

func (virtualMachineStrategy) AllowUnconditionalUpdate() bool {
	return false
}

func (virtualMachineStrategy) Canonicalize(obj runtime.Object) {
}

func (virtualMachineStrategy) ValidateUpdate(ctx genericapirequest.Context, obj, old runtime.Object) field.ErrorList {
	return field.ErrorList{}
}
