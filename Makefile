.PHONY: clean check test build package-snapshot

export GO111MODULE=on

TAG_NAME := $(shell git tag -l --contains HEAD)
SHA := $(shell git rev-parse HEAD)
VERSION := $(if $(TAG_NAME),$(TAG_NAME),$(SHA))

default: clean check test build

clean:
	rm -rf dist/

build: clean
	@echo Version: $(VERSION)
	go build -v -ldflags '-X "main.Version=${VERSION}"' -o deptomod .

test: clean
	go test -v -cover ./...

check:
	golangci-lint run

package-snapshot:
	goreleaser --skip-publish --skip-validate --rm-dist --snapshot
