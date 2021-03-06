#!/bin/bash

# The only argument this script should ever be called with is '--verify-only'

set -o errexit
set -o nounset
set -o pipefail

REPO_ROOT=$(dirname "${BASH_SOURCE}")/..
BINDIR=${REPO_ROOT}/bin

# Generate the internal clientset (pkg/client/clientset_generated/internalclientset)
${BINDIR}/client-gen "$@" \
        --input-base "github.com/dippynark/kube-qemu/pkg/apis/" \
        --input "hypervisor/" \
        --clientset-path "github.com/dippynark/kube-qemu/pkg/client/clientset_generated/" \
        --clientset-name "internalclientset" \
        --go-header-file "${GOPATH}/src/github.com/kubernetes/repo-infra/verify/boilerplate/boilerplate.go.txt"
# Generate the versioned clientset (pkg/client/clientset_generated/clientset)
${BINDIR}/client-gen "$@" \
        --input-base "github.com/dippynark/kube-qemu/pkg/apis/" \
        --input "hypervisor/v1alpha1" \
        --clientset-path "github.com/dippynark/kube-qemu/pkg/client/clientset_generated/" \
        --clientset-name "clientset" \
        --go-header-file "${GOPATH}/src/github.com/kubernetes/repo-infra/verify/boilerplate/boilerplate.go.txt"
# generate lister
${BINDIR}/lister-gen "$@" \
	--input-dirs="github.com/dippynark/kube-qemu/pkg/apis/hypervisor" \
        --input-dirs="github.com/dippynark/kube-qemu/pkg/apis/hypervisor/v1alpha1" \
        --output-package "github.com/dippynark/kube-qemu/pkg/client/listers_generated" \
        --go-header-file "${GOPATH}/src/github.com/kubernetes/repo-infra/verify/boilerplate/boilerplate.go.txt"
# generate informer
${BINDIR}/informer-gen "$@" \
        --go-header-file "${GOPATH}/src/github.com/kubernetes/repo-infra/verify/boilerplate/boilerplate.go.txt" \
        --input-dirs "github.com/dippynark/kube-qemu/pkg/apis/hypervisor" \
        --input-dirs "github.com/dippynark/kube-qemu/pkg/apis/hypervisor/v1alpha1" \
        --internal-clientset-package "github.com/dippynark/kube-qemu/pkg/client/clientset_generated/internalclientset" \
        --versioned-clientset-package "github.com/dippynark/kube-qemu/pkg/client/clientset_generated/clientset" \
        --listers-package "github.com/dippynark/kube-qemu/pkg/client/listers_generated" \
        --output-package "github.com/dippynark/kube-qemu/pkg/client/informers_generated"