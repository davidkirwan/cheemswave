##############################
##@ Build binary

.PHONY: clean
clean: ## Clean up compiled/generated files
	-rm ./bin/parallax_scrolling

.PHONY: build
build: ## Compile and build the binary
	@GOOS=linux GOARCH=amd64 CGO_ENABLED=1 GO111MODULE=on go build -mod=vendor -a -v -o=$(COMPILE_TARGET) ./cmd/parallax_scrolling/main.go
