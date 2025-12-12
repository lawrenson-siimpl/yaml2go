# Toolbox

TOP = $(shell dirname $(abspath $(lastword $(MAKEFILE_LIST))))

TARGET := yaml2go
COMMIT_SHA = "$(shell git rev-parse HEAD)"
GO_VERSION = "$(shell go env GOVERSION)"
BUILD_DATE = "$(shell date -u +%Y-%m-%dT%H:%M:%SZ)"
GOLDFLAGS = "\
	-X github.com/lawrenson-siimpl/yaml2go/version.Commit=${COMMIT_SHA} \
	-X github.com/lawrenson-siimpl/yaml2go/version.GoVersion=${GO_VERSION} \
	-X github.com/lawrenson-siimpl/yaml2go/version.BuildDate=${BUILD_DATE} \
	"

.PHONY: all clean build install test

all: build

clean:
	# rm -rf '$(TOP)/build'
	@echo "Not Implemented"

build: clean
	go build -ldflags $(GOLDFLAGS) -o $(TOP)/build/$(TARGET) ./...

install: build
	install -m 0755 $(TOP)/build/$(TARGET) $(shell go env GOBIN)/$(TARGET)

test:
	# go test -cover ./...
	@echo "Not Implemented"
