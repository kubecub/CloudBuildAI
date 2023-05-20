# Copyright 2023 KubeCub. All rights reserved.
# Use of this source code is governed by a MIT style
# license that can be found in the LICENSE file.

################################################################################
# ========================== Capture Environment ===============================
# get the repo root and output path
ROOT_PACKAGE=github.com/kubecub/feishu-sheet-parser
OUT_DIR=$(REPO_ROOT)/_output
# ==============================================================================
# define the default goal
#

SHELL := /bin/bash
DIRS=$(shell ls)
GO=go

.DEFAULT_GOAL := help

# include the common makefile
COMMON_SELF_DIR := $(dir $(lastword $(MAKEFILE_LIST)))
# ROOT_DIR: root directory of the code base
ifeq ($(origin ROOT_DIR),undefined)
ROOT_DIR := $(abspath $(shell cd $(COMMON_SELF_DIR)/. && pwd -P))
endif
# OUTPUT_DIR: The directory where the build output is stored.
ifeq ($(origin OUTPUT_DIR),undefined)
OUTPUT_DIR := $(ROOT_DIR)/bin
$(shell mkdir -p $(OUTPUT_DIR))
endif

ifeq ($(origin VERSION), undefined)
VERSION := $(shell git describe --abbrev=0 --dirty --always --tags | sed 's/-/./g')
endif

# Check if the tree is dirty. default to dirty(maybe u should commit?)
GIT_TREE_STATE:="dirty"
ifeq (, $(shell git status --porcelain 2>/dev/null))
	GIT_TREE_STATE="clean"
endif
GIT_COMMIT:=$(shell git rev-parse HEAD)

IMG ?= ghcr.io/kubecub/feishu-sheet-parser:latest

BUILDFILE = "./main.go"
BUILDAPP = "$(OUTPUT_DIR)/"

# Define the directory you want to copyright
CODE_DIRS := $(ROOT_DIR)/ #$(ROOT_DIR)/pkg $(ROOT_DIR)/core $(ROOT_DIR)/integrationtest $(ROOT_DIR)/lib $(ROOT_DIR)/mock $(ROOT_DIR)/db $(ROOT_DIR)/openapi
FINDS := find $(CODE_DIRS)

ifndef V
MAKEFLAGS += --no-print-directory
endif

# Linux command settings
FIND := find . ! -path './image/*' ! -path './vendor/*' ! -path './bin/*'
XARGS := xargs -r
LICENSE_TEMPLATE ?= $(ROOT_DIR)/scripts/LICENSE_TEMPLATES

# ==============================================================================
# Build definition

GO_SUPPORTED_VERSIONS ?= 1.18|1.19|1.20
GO_LDFLAGS += -X $(VERSION_PACKAGE).GitVersion=$(VERSION) \
	-X $(VERSION_PACKAGE).GitCommit=$(GIT_COMMIT) \
	-X $(VERSION_PACKAGE).GitTreeState=$(GIT_TREE_STATE) \
	-X $(VERSION_PACKAGE).BuildDate=$(shell date -u +'%Y-%m-%dT%H:%M:%SZ')
ifneq ($(DLV),)
	GO_BUILD_FLAGS += -gcflags "all=-N -l"
	LDFLAGS = ""
endif
GO_BUILD_FLAGS += -ldflags "$(GO_LDFLAGS)"

ifeq ($(GOOS),windows)
	GO_OUT_EXT := .exe
endif

ifeq ($(ROOT_PACKAGE),)
	$(error the variable ROOT_PACKAGE must be set prior to including golang.mk)
endif

GOPATH := $(shell go env GOPATH)
ifeq ($(origin GOBIN), undefined)
	GOBIN := $(GOPATH)/bin
endif

COMMANDS ?= $(filter-out %.md, $(wildcard ${ROOT_DIR}/cmd/*))
BINS ?= $(foreach cmd,${COMMANDS},$(notdir ${cmd}))

ifeq (${COMMANDS},)
  $(error Could not determine COMMANDS, set ROOT_DIR or run in source dir)
endif
ifeq (${BINS},)
  $(error Could not determine BINS, set ROOT_DIR or run in source dir)
endif

EXCLUDE_TESTS=github.com/kubecub/CloudBuildAI/test

# ==============================================================================
# Build

## all: Build all the necessary targets.
.PHONY: all
all: tidy add-copyright lint cover build

## build: Build binaries by default.
.PHONY: build
build: go.build.verify $(addprefix go.build., $(addprefix $(PLATFORM)., $(BINS)))

## build.%: Builds a binary of the specified directory.
.PHONY: build.%
build.%:
	@echo "$(shell go version)"
	@echo "===========> Building binary $(BUILDAPP) *[Git Info]: $(VERSION)-$(GIT_COMMIT)"
	@export CGO_ENABLED=0 && GOOS=linux go build -o $(BUILDAPP)/$*/ -ldflags '-s -w' $*/example/$(BUILDFILE)

.PHONY: go.build.verify
go.build.verify:
ifneq ($(shell $(GO) version | grep -q -E '\bgo($(GO_SUPPORTED_VERSIONS))\b' && echo 0 || echo 1), 0)
	$(error unsupported go version. Please make install one of the following supported version: '$(GO_SUPPORTED_VERSIONS)')
endif

.PHONY: go.build.%
go.build.%:
	$(eval COMMAND := $(word 2,$(subst ., ,$*)))
	$(eval PLATFORM := $(word 1,$(subst ., ,$*)))
	$(eval OS := $(word 1,$(subst _, ,$(PLATFORM))))
	$(eval ARCH := $(word 2,$(subst _, ,$(PLATFORM))))
	@echo "=====> COMMAND=$(COMMAND)"
	@echo "=====> PLATFORM=$(PLATFORM)"
	@echo "=====> BIN_DIR=$(BIN_DIR)"
	@echo "===========> Building binary $(COMMAND) $(VERSION) for $(OS)_$(ARCH)"
	@mkdir -p $(OUTPUT_DIR)/platforms/$(OS)/$(ARCH)
	@CGO_ENABLED=0 GOOS=$(OS) GOARCH=$(ARCH) $(GO) build $(GO_BUILD_FLAGS) -o $(OUTPUT_DIR)/platforms/$(OS)/$(ARCH)/$(COMMAND)$(GO_OUT_EXT) $(ROOT_PACKAGE)/cmd/$(COMMAND)

.PHONY: go.build.multiarch
go.build.multiarch: go.build.verify $(foreach p,$(PLATFORMS),$(addprefix go.build., $(addprefix $(p)., $(BINS))))

# ==============================================================================
# Targets

## tidy: tidy go.mod
.PHONY: tidy
tidy:
	@$(GO) mod tidy

## fmt: Run go fmt against code.
.PHONY: fmt
fmt:
	@$(GO) fmt ./...

## vet: Run go vet against code.
.PHONY: vet
vet:
	@$(GO) vet ./...

## generate: Run go generate against code.
.PHONY: generate
generate:
	@$(GO) generate ./...

## lint: Run go lint against code.
.PHONY: lint
lint:
	@echo "===========> Run golangci to lint source codes"
	@golangci-lint run -c $(ROOT_DIR)/.golangci.yml $(ROOT_DIR)/...

## style: Code style -> fmt,vet,lint
.PHONY: style
style: fmt vet lint

## test: Run unit test
.PHONY: test
test: 
	@$(GO) test ./... 

## cover: Run unit test with coverage.
.PHONY: cover
cover: test
	@$(GO) test -cover

## copyright.verify: Validate boilerplate headers for assign files.
.PHONY: copyright-verify
copyright-verify: 
	@echo "===========> Validate boilerplate headers for assign files starting in the $(ROOT_DIR) directory"
	@addlicense -v -check -ignore **/test/** -f $(LICENSE_TEMPLATE) $(CODE_DIRS)
	@echo "===========> End of boilerplate headers check..."

## copyright-add: Add the boilerplate headers for all files.
.PHONY: copyright-add
copyright-add: 
	@echo "===========> Adding $(LICENSE_TEMPLATE) the boilerplate headers for all files"
	@addlicense -y $(shell date +"%Y") -v -c "KubeCub & Xinwei Xiong(cubxxw)." -f $(LICENSE_TEMPLATE) $(CODE_DIRS)
	@echo "===========> End the copyright is added..."

## go.clean: Clean all builds.
.PHONY: clean
clean:
	@echo "===========> Cleaning all builds OUTPUT_DIR($(OUTPUT_DIR))"
	@-rm -vrf $(OUTPUT_DIR)
	@echo "===========> End clean..."

## help: Show this help info.
.PHONY: help
help: Makefile
	@printf "\n\033[1mUsage: make <TARGETS> ...\033[0m\n\n\\033[1mTargets:\\033[0m\n\n"
	@sed -n 's/^##//p' $< | awk -F':' '{printf "\033[36m%-28s\033[0m %s\n", $$1, $$2}' | sed -e 's/^/ /'

################################################################################
