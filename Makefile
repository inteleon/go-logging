all: test

test: deps
	@echo "Running tests..."
	@cd logging && go fmt
	@cd helper && go fmt
	@go test ./...

deps:
	@echo "Fetching dependencies..."
	@go get -u github.com/sirupsen/logrus
