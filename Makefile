PACKAGE=$(shell basename `pwd`)
VERSION := $(shell cat VERSION)
DATE := $(shell date +%Y-%m-%dT%H:%M:%S%z)
BUILD_FLAGS = -ldflags "\
	      -X main.Version=$(VERSION) \
	      "

all: setup clean build test

setup:
	go get github.com/golang/lint/golint

fmt: setup
	go fmt ./...

vet: setup
	go vet ./...

lint: setup
	golint ./...

test:
	go test -v ./...

bench: setup
	go test ./... -bench=.

doc: setup
	godoc -http=:6060

deps: setup
	go get -d -v ./...

deps_local:
	rm -fr ${GOPATH}/src/github.com/hiracy/zabton
	cp -r ../$(PACKAGE) ${GOPATH}/src/github.com/hiracy

clean:
	rm -fr ${GOPATH}/src/github.com/hiracy/zabton
	go clean

build: fmt deps deps_local
	go build $(BUILD_FLAGS) -o $(PACKAGE) .

.PHONY: fmt vet lint test bench doc deps clean build_cli build
