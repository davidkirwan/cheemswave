##############################
##@ Testing

.PHONY: test-unit
test-unit: ## Execute the unit tests
	go get -u github.com/rakyll/gotest
	gotest -v -covermode=count -coverprofile=coverage.out ./internal/pkg/resources/...