##############################
##@ Build binary

.PHONY: clean
clean: ## Clean up compiled/generated files
	-rm ./build/parallax_scrolling
	-rm ./build/parallax_scrolling.exe
	-rm -rf ./build/assets

.PHONY: build
build: ## Compile and build the binary
	@GOOS=linux GOARCH=amd64 CGO_ENABLED=1 GO111MODULE=on go build -mod=vendor -a -v -o=$(COMPILE_TARGET) ./cmd/parallax_scrolling/main.go

.PHONY: build-windows
build-windows: ## Compile and build the binary for MS Windows
	@CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=amd64 GO111MODULE=on go build -mod=vendor -a -v -o=$(COMPILE_TARGET).exe ./cmd/parallax_scrolling/main.go

.PHONY: build-copy-assets
build-copy-assets: ## Copy the assets to the build directory
	cp -R ./assets ./build
