GOBIN ?= $$(go env GOPATH)/bin

.PHONY: prepare test test-coverage html-coverage

prepare: install-tools

install-tools:
	@go install github.com/vladopajic/go-test-coverage/v2@latest

test:
	@go test

test-coverage:
	@go test ./... -coverprofile=./cover.out -covermode=atomic -coverpkg=./...

html-coverage: test-coverage
	@go tool cover -html=./cover.out -o ./cover.html
	@open ./cover.html
