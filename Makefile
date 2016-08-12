PACKAGE=$(shell basename `pwd`)
VERSION := $(shell cat VERSION)
DATE := $(shell date +%Y-%m-%dT%H:%M:%S%z)
BUILD_FLAGS = -ldflags "\
	      -X main.Version=$(VERSION) \
              -X main.buildDate=$(DATE) \
	      "

all: clean build test

fmt:
	go fmt ./...

vet:
	go vet ./...

lint:
	golint ./...

test:
	go test -v ./...

bench:
	go test ./... -bench=.

doc:
	godoc -http=:6060

deps:
	go get -d -v ./...

clean:
	go clean

build_cli: deps
	go build $(BUILD_FLAGS) -o $(PACKAGE) ./cmd/...

build: build_cli

.PHONY: fmt vet lint test bench doc deps clean build_cli build
