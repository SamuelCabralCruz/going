.PHONY: build
build:
	@go build -tags=test ./...

.PHONY: test-local
test-local:
	@mkdir -p coverage
	@ginkgo -tags=test -r --coverpkg=./... --coverprofile=coverage.out.tmp --output-dir=coverage
	@cat coverage/coverage.out.tmp | grep -vE "/testing/|_test\.go|_?fixture\.go|_?mock\.go" > coverage/coverage.out
	@go tool cover -html=coverage/coverage.out -o coverage/coverage.html
	@go tool cover -func=coverage/coverage.out | perl -n -e '/total:\D*(\d+\.\d+)%/ && print $$1' > coverage/metric.txt

.PHONY: test-local-verbose
test-local-verbose:
	@ginkgo -tags=test -r -v ./...
