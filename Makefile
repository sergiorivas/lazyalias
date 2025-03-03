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

build:
	$(GOBUILD) -o bin/$(BINARY_NAME) $(MAIN_PATH)

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
