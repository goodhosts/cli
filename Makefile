# Go and compilation related variables
BUILD_DIR ?= out

BINARY_NAME := goodhosts
RELEASE_DIR ?= release

# Add default target
.PHONY: all
all: build

vendor:
	go mod vendor

$(BUILD_DIR):
	mkdir -p $(BUILD_DIR)

.PHONY: clean
clean:
	rm -rf $(BUILD_DIR)
	rm -rf vendor
	rm -fr release

$(BUILD_DIR)/macos-amd64/$(BINARY_NAME):
	GOARCH=amd64 GOOS=darwin go build -o $(BUILD_DIR)/macos-amd64/$(BINARY_NAME) ./main.go

$(BUILD_DIR)/linux-amd64/$(BINARY_NAME):
	GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/linux-amd64/$(BINARY_NAME) ./main.go

$(BUILD_DIR)/windows-amd64/$(BINARY_NAME).exe:
	GOARCH=amd64 GOOS=windows go build -o $(BUILD_DIR)/windows-amd64/$(BINARY_NAME).exe ./main.go

.PHONY: cross ## Cross compiles all binaries
cross: $(BUILD_DIR)/macos-amd64/$(BINARY_NAME) $(BUILD_DIR)/linux-amd64/$(BINARY_NAME) $(BUILD_DIR)/windows-amd64/$(BINARY_NAME).exe

.PHONY: release
release: clean cross
	mkdir $(RELEASE_DIR)
	tar cJSf $(RELEASE_DIR)/goodhosts-cli-macos-amd64.tar.xz -C $(BUILD_DIR)/macos-amd64 $(BINARY_NAME)
	tar cJSf $(RELEASE_DIR)/goodhosts-cli-linux-amd64.tar.xz -C $(BUILD_DIR)/linux-amd64 $(BINARY_NAME)
	tar cJSf $(RELEASE_DIR)/goodhosts-cli-windows-amd64.tar.xz -C $(BUILD_DIR)/windows-amd64 $(BINARY_NAME).exe

	pushd $(RELEASE_DIR) && sha256sum * > sha256sum.txt && popd

.PHONY: build
build:
	go build -o $(BINARY_NAME) ./main.go

