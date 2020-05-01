##############################
##@ Fetch dependencies

.PHONY: dependencies
dependencies: ## Pull down the dependencies
	GO111MODULE=on go mod vendor && go mod tidy
	modvendor -copy="**/*.c **/*.h **/*.proto" -v
