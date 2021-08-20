####################################################################################################
# @author David Kirwan https://github.com/davidkirwan/parallax_scrolling
# @description Makefile for building/executing the parallax scrolling demo
#
# @usage make
#
# @date 2020-05-01
####################################################################################################
include ./scripts/*.mk
.DEFAULT_GOAL:=help
PROJECT=parallax_scrolling
COMPILE_TARGET=./build/$(PROJECT)
##############################

.PHONY: help
help: ## Show this help screen
	@echo 'Usage: make <OPTIONS> ... <TARGETS>'
	@echo ''
	@echo 'Available targets are:'
	@echo ''
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z0-9_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##############################
##@ All

.PHONY: all
all: dependencies clean build build-copy-assets run ## Pull dependencies, builds binary, executes binary all in one
