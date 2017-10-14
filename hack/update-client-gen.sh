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
        --clientset-path "github.com/dippynark/kube-qemu/pkg/client/" \
        --clientset-name internalclientset \
        --go-header-file "${GOPATH}/src/github.com/kubernetes/repo-infra/verify/boilerplate/boilerplate.go.txt"
# Generate the versioned clientset (pkg/client/clientset_generated/clientset)
${BINDIR}/client-gen "$@" \
        --input-base "github.com/dippynark/kube-qemu/pkg/apis/" \
        --input "hypervisor/v1alpha1" \
        --clientset-path "github.com/dippynark/kube-qemu/pkg/" \
        --clientset-name "client" \
        --go-header-file "${GOPATH}/src/github.com/kubernetes/repo-infra/verify/boilerplate/boilerplate.go.txt"