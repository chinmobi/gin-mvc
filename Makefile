
include ./BUILD.variables


PROJECT_NAME ?= $(shell basename "$(CURDIR)")
BINARY_NAME ?= $(PROJECT_NAME)

GOPATH ?= $(CURDIR)/..

OUTPUT_DIR := $(GOPATH)/bin
OUTPUT_BIN := $(OUTPUT_DIR)/$(BINARY_NAME)


# Basic shell commands
RM := rm -f

# Basic go commands
GOCMD := go
GOBUILD := $(GOCMD) build
GOCLEAN := $(GOCMD) clean
GOTEST := $(GOCMD) test


.PHONY: all build test clean help

.DEFAULT_GOAL := help

all: test build


build:
	$(GOBUILD) -o $(OUTPUT_BIN) -v $(BUILDFLAGS)

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	$(RM) $(OUTPUT_BIN)


help:
	@echo
	@echo '  Usage:'
	@echo '    make <target>'
	@echo
	@echo '  Targets:'
	@echo '    all           Test and build'
	@echo '    build         Compile packages and dependencies'
	@echo '    test          Test packages'
	@echo '    clean         Remove object files and cached files'
	@echo
	@echo '    help          Show this help message'
	@echo
