#
# Makefile
#
# The objective of this is to provide a local execution environment that is as close to identical to the CI/CD platform
# as possible. It provides a number of targets that can be used locally by a developer to ensure that tests and builds
# behave consistently and gives the developer an increased amount of confidence that the their code is valid.
#
# There are elegant ways to automatically run Makefiles within subdirectories, however, for the purposes of getting
# a process into place quickly, all targets are predefined.
#
# In order to use this, one must have an Azure DevOps account, configure a Personal Access Token (PAT) with at least
# Packaging (read) permission, and set the environment variable AZURE_DEVOPS_EXT_PAT to the PAT.
#
################################################################################################
#
SHELL  := /bin/bash
.DEFAULT_GOAL := help

################################################################################################
# Environment Variables

# The help/usage information
define USAGE
The following project build targets are available.

  make build                      # Builds all the projects.

The following local deployment targets are available.

  make start
  make stop

endef
export USAGE

################################################################################################
# Helper Targets
################################################################################################



# Start the Application stack
.PHONY: start
start:
	@docker-compose down --remove-orphans
	@docker-compose up

# Stop the Application stack
.PHONY: stop
stop:
	@docker-compose down --remove-orphans

.PHONY: build
build: 
	@make -C gql build
	@make -C fibre build
