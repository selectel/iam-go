default: workflow

workflow: golangci-lint unit-test

unit-test:
	@echo "--- Running unit tests ---"
	go test -v ./...

unit-test-race:
	@echo "--- Running unit tests -race ---"
	go test -v -race ./...

golangci-lint:
	@echo "--- Running golangci-lint ---"
	golangci-lint run ./...

.PHONY: workflow unit-test golangci-lint