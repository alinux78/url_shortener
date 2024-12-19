# Variables
BINARY_NAME=urlshortener
BUILD_DIR=bin
SRC_DIR=cmd/server

# Default target
all: build

# Build the binary
build:
	@echo "Building the binary..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) $(SRC_DIR)/main.go

# Run the application
run: build
	@echo "Running the application..."
	@./$(BUILD_DIR)/$(BINARY_NAME)

# Clean the build directory
clean:
	@echo "Cleaning the build directory..."
	@rm -rf $(BUILD_DIR)

# Format the code
fmt:
	@echo "Formatting the code..."
	@go fmt ./...

# Run tests
test:
	@echo "Running tests..."
	@go test ./...

# Lint the code
lint:
	@echo "Linting the code..."
	@golangci-lint run

# Help message
help:
	@echo "Makefile commands:"
	@echo "  build   - Build the binary"
	@echo "  run     - Run the application"
	@echo "  clean   - Clean the build directory"
	@echo "  fmt     - Format the code"
	@echo "  test    - Run tests"
	@echo "  lint    - Lint the code"
	@echo "  help    - Show this help message"

.PHONY: all build run clean fmt test lint help
