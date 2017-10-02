all: test

test: deps
	@echo "Running tests..."
	@cd logging && go fmt
	@cd helper && go fmt
	@go test ./...

clean:
	@glide cc
	@rm -rf .glide/
	@rm -rf vendor/

deps: clean
	@echo "Fetching dependencies..."
	@glide install
