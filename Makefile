BINDIR ?= bin
GOPATH ?= $HOME/go
HACK_DIR ?= hack
PACKAGE_NAME := github.com/dippynark/kube-qemu

# A list of all types.go files in pkg/apis
TYPES_FILES = $(shell find pkg/apis -name types.go)

REGISTRY := dippynark
IMAGE_NAME := kube-qemu
BUILD_TAG := build
IMAGE_TAGS := canary

CMDS := controller apiserver

# Util targets
##############

generate: .generate_files

# Docker targets
################

DOCKER_BUILD_TARGETS = $(addprefix docker_build_, $(CMDS))
$(DOCKER_BUILD_TARGETS):
	$(eval DOCKER_BUILD_CMD := $(subst docker_build_,,$@))
	docker build -t $(REGISTRY)/$(IMAGE_NAME)-$(DOCKER_BUILD_CMD):$(BUILD_TAG) -f Dockerfile.$(DOCKER_BUILD_CMD) .

DOCKER_PUSH_TARGETS = $(addprefix docker_push_, $(CMDS))
$(DOCKER_PUSH_TARGETS):
	$(eval DOCKER_PUSH_CMD := $(subst docker_push_,,$@))
	set -e; \
		for tag in $(IMAGE_TAGS); do \
		docker tag $(REGISTRY)/$(IMAGE_NAME)-$(DOCKER_PUSH_CMD):$(BUILD_TAG) $(REGISTRY)/$(IMAGE_NAME)-$(DOCKER_PUSH_CMD):$${tag} ; \
		docker push $(REGISTRY)/$(IMAGE_NAME)-$(DOCKER_PUSH_CMD):$${tag}; \
	done
docker_push: $(DOCKER_PUSH_TARGETS)

# Go targets
#################

$(CMDS):
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w' -o $(BINDIR)/kube-qemu-$@_linux_amd64 ./cmd/$@

go_build: $(CMDS)

# This section contains the code generation stuff
#################################################
.generate_exes: \
	$(BINDIR)/defaulter-gen \
	$(BINDIR)/deepcopy-gen \
	$(BINDIR)/conversion-gen \
	$(BINDIR)/client-gen \
	$(BINDIR)/lister-gen \
	$(BINDIR)/informer-gen
	touch $@

$(BINDIR)/%:
	go build -o $@ ./vendor/k8s.io/code-generator/cmd/$*

# Regenerate all files if the gen exes changed or any "types.go" files changed
.generate_files: .generate_exes $(TYPES_FILES)
	# Generate defaults
	$(BINDIR)/defaulter-gen \
		--v 1 --logtostderr \
		--go-header-file "$(HACK_DIR)/boilerplate.go.txt" \
		--input-dirs "$(PACKAGE_NAME)/pkg/apis/hypervisor" \
		--input-dirs "$(PACKAGE_NAME)/pkg/apis/hypervisor/v1alpha1" \
		--extra-peer-dirs "$(PACKAGE_NAME)/pkg/apis/hypervisor" \
		--extra-peer-dirs "$(PACKAGE_NAME)/pkg/apis/hypervisor/v1alpha1" \
		--output-file-base "zz_generated.defaults"
	# Generate deep copies
	$(BINDIR)/deepcopy-gen \
		--v 1 --logtostderr \
		--go-header-file "$(HACK_DIR)/boilerplate.go.txt" \
		--input-dirs "$(PACKAGE_NAME)/pkg/apis/hypervisor" \
		--input-dirs "$(PACKAGE_NAME)/pkg/apis/hypervisor/v1alpha1" \
		--output-file-base zz_generated.deepcopy
	# Generate conversions
	$(BINDIR)/conversion-gen \
		--v 1 --logtostderr \
		--go-header-file "$(HACK_DIR)/boilerplate.go.txt" \
		--input-dirs "$(PACKAGE_NAME)/pkg/apis/hypervisor" \
		--input-dirs "$(PACKAGE_NAME)/pkg/apis/hypervisor/v1alpha1" \
		--output-file-base zz_generated.conversion
	# generate all pkg/client contents
	$(HACK_DIR)/update-client-gen.sh
	