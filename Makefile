all: test

test: deps
	@echo "Running tests..."
	@cd logging && go fmt
	@cd helper && go fmt
	@go test ./...

clean:
	@rm -rf vendor/

deps: clean
	@echo "Fetching dependencies..."
	@go mod tidy
