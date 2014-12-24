NAME := michel

VERSION := 0.1.0-dev
GIT_COMMIT := $(shell git rev-parse HEAD)

GOFLAGS := -ldflags "-X main.Version ${VERSION} -X main.GitCommit ${GIT_COMMIT}"

all: format test

format:
	@go fmt ./...

deps:
	@go get -d ./...

test: deps
	@go test -v ./...

dev: test
	@go run $(GOFLAGS) main.go

build:
	@go build $(GOFLAGS) -o bin/${NAME}

clean:
	@go clean ./...

.PHONY: format deps test dev build clean