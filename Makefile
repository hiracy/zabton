PACKAGE=$(shell basename `pwd`)
VERSION := $(shell git describe --tags | awk -F'-' '{print $$1}')
TARGET_DIR = dist
BUILD_FLAGS = -ldflags "\
	      -X main.Version=$(VERSION) \
	      "

all: clean deps build test

setup:
	go get github.com/golang/lint/golint

fmt:
	go fmt ./...

vet:
	go vet ./...

lint: setup
	golint ./...

test:
	go test -v ./...

bench:
	go test ./... -bench=.

doc:
	godoc -http=:6060

deps: deps_local
	go get -d -v ./...

deps_local:
	rm -fr ${GOPATH}/src/github.com/hiracy/zabton
	cp -r ../$(PACKAGE) ${GOPATH}/src/github.com/hiracy

clean:
	rm -fr ${GOPATH}/src/github.com/hiracy/zabton
	go clean

build: fmt
	go build $(BUILD_FLAGS) -o $(PACKAGE) .

.PHONY: fmt vet lint test bench doc deps clean build_cli build
