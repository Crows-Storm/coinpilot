# CoinPilot Makefile

# Variables
BINARY_NAME=coinpilot
MAIN_PATH=cmd/coinpilot/main.go
BUILD_DIR=build

# Default target
.PHONY: all
all: build

# Build the application
.PHONY: build
build:
	@echo "Building CoinPilot..."
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PATH)
	@echo "Build complete: $(BUILD_DIR)/$(BINARY_NAME)"

# Run the application
.PHONY: run
run: build
	@echo "Running CoinPilot..."
	./$(BUILD_DIR)/$(BINARY_NAME)

# Clean build artifacts
.PHONY: clean
clean:
	@echo "Cleaning build artifacts..."
	rm -rf $(BUILD_DIR)
	go clean

# Download dependencies
.PHONY: deps
deps:
	@echo "Downloading dependencies..."
	go mod download
	go mod tidy

# Run tests
.PHONY: test
test:
	@echo "Running tests..."
	go test ./...

# Run tests with coverage
.PHONY: test-coverage
test-coverage:
	@echo "Running tests with coverage..."
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# Format code
.PHONY: fmt
fmt:
	@echo "Formatting code..."
	go fmt ./...

# Lint code
.PHONY: lint
lint:
	@echo "Linting code..."
	golangci-lint run

# Install the binary
.PHONY: install
install: build
	@echo "Installing CoinPilot..."
	cp $(BUILD_DIR)/$(BINARY_NAME) /usr/local/bin/

# Cross-platform builds
.PHONY: build-all
build-all:
	@echo "Building for all platforms..."
	@mkdir -p $(BUILD_DIR)
	# macOS builds
	GOOS=darwin GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-macos-intel $(MAIN_PATH)
	GOOS=darwin GOARCH=arm64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-macos-arm64 $(MAIN_PATH)
	# Windows builds
	GOOS=windows GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe $(MAIN_PATH)
	GOOS=windows GOARCH=386 go build -o $(BUILD_DIR)/$(BINARY_NAME)-windows-386.exe $(MAIN_PATH)
	# Linux builds
	GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 $(MAIN_PATH)
	GOOS=linux GOARCH=arm64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-linux-arm64 $(MAIN_PATH)
	@echo "Cross-platform builds complete"

# Build macOS universal binary (requires macOS)
.PHONY: build-macos-universal
build-macos-universal: build-all
	@echo "Creating macOS universal binary..."
	@if [ "$$(uname)" = "Darwin" ]; then \
		lipo -create -output $(BUILD_DIR)/$(BINARY_NAME)-macos-universal \
			$(BUILD_DIR)/$(BINARY_NAME)-macos-intel \
			$(BUILD_DIR)/$(BINARY_NAME)-macos-arm64; \
		echo "macOS universal binary created"; \
	else \
		echo "macOS universal binary can only be created on macOS"; \
	fi

# Help
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  build         - Build the application"
	@echo "  run           - Build and run the application"
	@echo "  clean         - Clean build artifacts"
	@echo "  deps          - Download and tidy dependencies"
	@echo "  test          - Run tests"
	@echo "  test-coverage - Run tests with coverage report"
	@echo "  fmt           - Format code"
	@echo "  lint          - Lint code"
	@echo "  install       - Install binary to /usr/local/bin"
	@echo "  build-all     - Cross-platform builds for Windows, macOS, and Linux
  build-macos-universal - Create macOS universal binary (macOS only)"
	@echo "  help          - Show this help message"