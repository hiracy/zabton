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
	go test ./...

bench:
	go test ./... -bench=.

doc:
	godoc -http=:6060

deps:
	go get -d -v ./...

clean:
	go clean

build: deps
	go build $(BUILD_FLAGS) -o $(PACKAGE) ./cmd

.PHONY: fmt vet lint test bench doc deps clean build
