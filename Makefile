# Go and compilation related variables
BUILD_DIR ?= dist
BINARY_NAME := goodhosts
GOLANGCI_LINT_VERSION ?= v1.54.2

$(BUILD_DIR):
	mkdir -p $(BUILD_DIR)

clean:
	rm -rf $(BUILD_DIR)

golangci-lint:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s $(GOLANGCI_LINT_VERSIONM)

goimports:
	go install golang.org/x/tools/cmd/goimports@latest

goreleaser:
	go install github.com/goreleaser/goreleaser@latest

release:
	goreleaser release

ci: goimports golangci-lint
	goimports -d .
	golangci-lint run
	go test -v ./...

build:
	go build -o $(BINARY_NAME) ./main.go

