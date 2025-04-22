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

all: style forat vet build

style:
	@echo ">> checking code style"
	@! gofmt -d $(shell find . -name '*.go') | grep '^' || (echo "Code not properly formatted"; exit 1)

format:
	@echo ">> formatting code"
	go fmt $(pkgs)

vet:
	@echo ">> vetting code"
	go vet $(pkgs)

build:
	@echo ">> building code"
	CGO_ENABLED=0 go build -ldflags "$(VERSION_LDFLAGS)" -o ../../process-exporter -a -tags netgo
