PACKAGE=$(shell basename `pwd`)
VERSION := $(shell git describe --tags | awk -F'-' '{print $$1}')
TARGET_DIR = dist
BUILD_FLAGS = -ldflags "\
	      -X main.Version=$(VERSION) \
	      "

all: clean deps build

setup:
	go get github.com/golang/lint/golint

fmt:
	go fmt ./...

vet:
	go vet ./...

lint: setup
	golint ./...

setup_docker:
	@./test/setup_docker.sh

test: setup_docker
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
	go build $(BUILD_FLAGS) -o "$(TARGET_DIR)/$(PACKAGE)" .

.PHONY: fmt vet lint test bench doc deps clean build_cli build
