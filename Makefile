SHELL=/bin/bash

default:
	make fmt lint test

test:
	go test ./cmd ./app_config ./crypto ./file_helpers --cover

lint:
	staticcheck ./...

fmt:
	gofmt -l -s -w ./

build:
	go build -o dist/secrets-manager .

install:
	go mod download

version:
	source scripts/make-helpers.bash && publish_new_version $(level)
