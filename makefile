pkgs := $(shell go list ./...)

BRANCH      ?= $(shell git rev-parse --abbrev-ref HEAD)
BUILDDATE   ?= $(shell date --iso-8601=seconds)
BUILDUSER   ?= $(shell whoami)@$(shell hostname)
REVISION    ?= $(shell git rev-parse HEAD)
TAG_VERSION ?= $(shell git describe --tags --abbrev=0)

VERSION_LDFLAGS := \
  -X github.com/prometheus/common/version.Branch=$(BRANCH) \
  -X github.com/prometheus/common/version.BuildDate=$(BUILDDATE) \
  -X github.com/prometheus/common/version.BuildUser=$(BUILDUSER) \
  -X github.com/prometheus/common/version.Revision=$(REVISION) \
  -X main.version=$(TAG_VERSION)

.PHONY: all style format vet test build smoke

all: style vet test build

style:
	@echo ">> checking code style"
	@! gofmt -d $(shell find . -name '*.go') | grep '^' || (echo "Code not properly formatted"; exit 1)

format:
	@echo ">> formatting code"
	go fmt $(pkgs)

vet:
	@echo ">> vetting code"
	go vet $(pkgs)

test:
	@echo ">> running short tests"
	go test -short $(pkgs)

build:
	@echo ">> building for amd64"
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags '$(VERSION_LDFLAGS)' -o ssh-exporter-amd64 -a -tags netgo
#	@echo ">> building for arm64"
#	GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -ldflags '$(VERSION_LDFLAGS)' -o ssh-exporter-arm64 -a -tags netgo
