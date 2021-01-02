
include ./BUILD.variables
include ./_deploy/BUILD.variables
include ./_deploy/.env


PROJECT_NAME ?= $(shell basename "$(CURDIR)")
BINARY_NAME ?= $(PROJECT_NAME)

GOPATH ?= $(CURDIR)/..

OUTPUT_DIR := $(GOPATH)/bin
OUTPUT_BIN := $(OUTPUT_DIR)/$(BINARY_NAME)


GO_TEST_DIRS = $(shell ./go-test-dirs.sh)
GO_TEST_SUITES = $(shell ./go-test-suite-dirs.sh)


# Deployment variables
APP_NAME ?= $(BINARY_NAME)

DEPLOY_OS_ARCH = $(shell echo $(IMAGE_OS_ARCH))

DEPLOY_OUTPUT_DIR := ./_deploy/bin
DEPLOY_OUTPUT_BIN := $(DEPLOY_OUTPUT_DIR)/$(APP_NAME)

DOCKER_SCRIPTS_DIR := ./_deploy/docker


# Basic shell commands
RM := rm -f

# Basic go commands
GOCMD := go
GOBUILD := $(GOCMD) build
GOCLEAN := $(GOCMD) clean
GOTEST := $(GOCMD) test


.PHONY: all build test test-suites clean help

.DEFAULT_GOAL := help

all: test build


build:
	$(GOBUILD) -o $(OUTPUT_BIN) -v $(BUILDFLAGS)

test: $(GO_TEST_DIRS)
	@for dir in $^; do \
		go test -v ./$$dir ; \
	done;

test-suites: $(GO_TEST_SUITES)
	@for dir in $^; do \
		go test -v ./$$dir ; \
	done;

clean:
	$(GOCLEAN)
	$(RM) $(OUTPUT_BIN)
	$(RM) $(DEPLOY_OUTPUT_BIN)


.PHONY: deploy-build docker-image docker-run docker-stop docker-clean

deploy-build:
	CGO_ENABLED=0 $(DEPLOY_OS_ARCH) $(GOBUILD) -o $(DEPLOY_OUTPUT_BIN) -v $(BUILDFLAGS)

docker-image: deploy-build
	@ $(DOCKER_SCRIPTS_DIR)/build-image.sh

docker-run: docker-image
	@ $(DOCKER_SCRIPTS_DIR)/run.sh

docker-stop:
	@ $(DOCKER_SCRIPTS_DIR)/stop.sh

docker-clean: docker-stop
	@ $(DOCKER_SCRIPTS_DIR)/clean-image.sh

help:
	@echo
	@echo '  Usage:'
	@echo '    make <target>'
	@echo
	@echo '  Targets:'
	@echo '    all           Test and build'
	@echo '    build         Compile packages and dependencies'
	@echo '    test          Test packages'
	@echo '    test-suites   Run integration test suites'
	@echo '    clean         Remove object files and cached files'
	@echo
	@echo '    deploy-build  Compile packages for deployment'
	@echo '    docker-image  Build docker image'
	@echo '    docker-run    Run docker container'
	@echo '    docker-stop   Stop and clean docker container'
	@echo '    docker-clean  Clean docker image'
	@echo
	@echo '    help          Show this help message'
	@echo
