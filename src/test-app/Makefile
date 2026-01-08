.PHONY: lint vet staticcheck fmt test run

# Run all linters
lint: vet staticcheck

# Run go vet
vet:
	@echo "Running go vet..."
	@go vet ./...

# Run staticcheck
staticcheck:
	@echo "Running staticcheck..."
	@staticcheck ./...

# Format code
fmt:
	@echo "Formatting code..."
	@go fmt ./...

# Run tests
test:
	@echo "Running tests..."
	@go test -v ./...

# Run the application
run:
	@go run main.go
