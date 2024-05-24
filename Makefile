.PHONY: help
help:
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-21s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Set up

.PHONY: install
install: ## Set up local.
	@mise install
	@go mod tidy

##@ Build

.PHONY: build
build: ## Build project.
	@go build -tags=test ./...

##@ Test

.PHONY: test-local
test-local: ## Run tests with coverage.
	@mkdir -p coverage
	@ginkgo -tags=test -r --coverpkg=./... --coverprofile=coverage.out.tmp --output-dir=coverage
	@cat coverage/coverage.out.tmp | grep -vE "/testing/|_test\.go|_?fixture\.go|_?mock\.go" > coverage/coverage.out
	@go tool cover -html=coverage/coverage.out -o coverage/coverage.html
	@go tool cover -func=coverage/coverage.out | perl -n -e '/total:\D*(\d+\.\d+)%/ && print $$1' > coverage/metric.txt

.PHONY: test-local-verbose
test-local-verbose: ## Run tests with verbose.
	@ginkgo -tags=test -r -v ./...
