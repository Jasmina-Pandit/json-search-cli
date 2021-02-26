SHELL:=$(shell /usr/bin/env which bash)

# project details
APPNAME = json-search-cli

run:
	go run .

format:
	go fmt .
test:
	go test -cover -v ./cmd/... ./reader/...