SHELL:=$(shell /usr/bin/env which bash)

# project details
APPNAME = json-search-cli

run:
	go run search.go ${type} ${key} ${value}

format:
	go fmt .
test:
	go test -cover -v ./cmd/... ./reader/...