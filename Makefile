# Generate SHA256 checksums for all built binaries
sha256:
	shasum -a 256 bin/lazyalias-* > SHA256SUMS
BINARY_NAME=lazyalias
MAIN_PATH=./cmd/lazyalias/main.go
COVERAGE_FILE=coverage.out

GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOMOD=$(GOCMD) mod
GOGET=$(GOCMD) get
GORUN=$(GOCMD) run

UNAME_S := $(shell uname -s)

.PHONY: all build test coverage lint clean deps tidy install-lint

install-lint:
ifeq ($(UNAME_S),Darwin)
	@if ! command -v golangci-lint >/dev/null 2>&1; then \
		brew install golangci-lint; \
	fi
else ifeq ($(UNAME_S),Linux)
	@if ! command -v golangci-lint >/dev/null 2>&1; then \
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin v1.56.2; \
	fi
else
	@echo "Por favor, instala golangci-lint manualmente desde https://golangci-lint.run/usage/install/"
endif


# Versioning
VERSION := $(shell cat VERSION)
LDFLAGS=-ldflags="-w -s -X 'main.version=$(VERSION)'"

# Cross-compile for Linux (amd64)
build-linux-amd64:
	GOOS=linux GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o bin/$(BINARY_NAME)-linux-amd64 $(MAIN_PATH)

# Cross-compile for Linux (arm64)
build-linux-arm64:
	GOOS=linux GOARCH=arm64 $(GOBUILD) $(LDFLAGS) -o bin/$(BINARY_NAME)-linux-arm64 $(MAIN_PATH)

# Cross-compile for macOS (amd64)
build-darwin-amd64:
	GOOS=darwin GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o bin/$(BINARY_NAME)-darwin-amd64 $(MAIN_PATH)

# Cross-compile for macOS (arm64)
build-darwin-arm64:
	GOOS=darwin GOARCH=arm64 $(GOBUILD) $(LDFLAGS) -o bin/$(BINARY_NAME)-darwin-arm64 $(MAIN_PATH)

# Cross-compile for Windows (amd64)
build-windows-amd64:
	GOOS=windows GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o bin/$(BINARY_NAME)-windows-amd64.exe $(MAIN_PATH)

# Cross-compile for Windows (arm64)
build-windows-arm64:
	GOOS=windows GOARCH=arm64 $(GOBUILD) $(LDFLAGS) -o bin/$(BINARY_NAME)-windows-arm64.exe $(MAIN_PATH)

# Build all targets
build: build-linux-amd64 build-linux-arm64 build-darwin-amd64 build-darwin-arm64 build-windows-amd64 build-windows-arm64

run:
	$(GORUN) $(MAIN_PATH)

test:
	$(GOTEST) -v ./...

coverage:
	$(GOTEST) -coverprofile=$(COVERAGE_FILE) ./...
	$(GOCMD) tool cover -html=$(COVERAGE_FILE)

lint: install-lint
	golangci-lint run

clean:
	rm -f $(BINARY_NAME)
	rm -f $(COVERAGE_FILE)

deps:
	$(GOGET) -v -t -d ./...

tidy:
	$(GOMOD) tidy

all: lint test build
