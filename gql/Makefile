#
# Makefile
#
# The objective of this is to provide a local execution environment that is as close to identical to the CI/CD platform
# as possible. It provides a number of targets that can be used locally by a developer to ensure that tests and builds
# behave consistently and gives the developer an increased amount of confidence that the their code is valid.
#
################################################################################################
#
SHELL  := /bin/bash
.DEFAULT_GOAL := help

################################################################################################
# Environment Variables
################################################################################################

# AWS settings
IMAGE_NAME ?= gql

define USAGE
The following targets are available:

  make build                # Build the image.

endef
export USAGE

################################################################################################
# Common Targets
################################################################################################

# Build the image
.PHONY: build
build:
	@echo "Building $(CURDIR)"
	@docker build . --tag $(IMAGE_NAME) 
