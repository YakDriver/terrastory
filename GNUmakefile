TEST?=./...
GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)

default: build

build: check
	go install

test: check
	go test $(TEST) -timeout=30s -parallel=4

check: lint
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"

fmt:
	@echo "==> Fixing source code with gofmt..."
	gofmt -s -w $(GOFMT_FILES)

lint:
	@echo "==> Checking source code against linters..."
	@GOGC=30 golangci-lint run $(TEST)

.PHONY: build test check fmt lint

