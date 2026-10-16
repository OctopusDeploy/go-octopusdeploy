.PHONY: test lint fmt

test:
	go test -v ./...

.PHONY: check-golangci-lint
check-golangci-lint:
	@command -v golangci-lint >/dev/null 2>&1 || { \
		echo "golangci-lint is not installed. See https://golangci-lint.run for instructions."; \
		exit 1; \
	}

.PHONY: lint
lint: check-golangci-lint
	golangci-lint run ./...

.PHONY: fmt
fmt: check-golangci-lint
	golangci-lint fmt ./...