all: build
.PHONY: all

# Ensure update-scripts are run before crd-gen so updates to Godoc are included in CRDs.
update-codegen-crds: update-scripts

# Include the library makefile
include $(addprefix ./vendor/github.com/openshift/build-machinery-go/make/, \
	golang.mk \
	targets/openshift/deps.mk \
	targets/openshift/crd-schema-gen.mk \
)

GO_PACKAGES :=$(addsuffix ...,$(addprefix ./,$(filter-out vendor/,$(filter-out hack/,$(wildcard */)))))
GO_BUILD_PACKAGES :=$(GO_PACKAGES)
GO_BUILD_PACKAGES_EXPANDED :=$(GO_BUILD_PACKAGES)
# LDFLAGS are not needed for dummy builds (saving time on calling git commands)
GO_LD_FLAGS:=
CONTROLLER_GEN_VERSION :=v0.7.0

# $1 - target name
# $2 - apis
# $3 - manifests
# $4 - output
# $(call add-crd-gen,release,./pkg/apis/release/v1alpha1/,./pkg/apis/release/v1alpha1/,./pkg/apis/release/v1alpha1/)

update-scripts:
	hack/update-codegen.sh
.PHONY: update-scripts

verify-scripts:
	hack/verify-codegen.sh
.PHONY: verify-scripts

generate-release-crd: ensure-controller-gen
	'$(CONTROLLER_GEN)' crd \
		paths=./pkg/apis/release/v1alpha1 \
		output:dir=./artifacts
.PHONY: generate-crds

clean:
	rm release-payload-controller
.PHONY: clean
