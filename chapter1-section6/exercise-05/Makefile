# Makefile for Go project

.PHONY: help build run test clean lint fmt vet docker

# Variables
APP_NAME=myapp
VERSION=1.0.0
BUILD_DIR=bin
DOCKER_IMAGE=$(APP_NAME):$(VERSION)

# Default target
help:
	@echo "Available commands:"
	@echo "  make build         - Build the application"
	@echo "  make run           - Run the application"
	@echo "  make test          - Run tests"
	@echo "  make test-coverage - Run tests with coverage"
	@echo "  make clean         - Remove build artifacts"
	@echo "  make lint          - Run linter"
	@echo "  make fmt           - Format code"
	@echo "  make vet           - Run go vet"
	@echo "  make all           - Run fmt, vet, lint, test, build"
	@echo "  make docker        - Build Docker image"

# Build binary
build:
	@echo "Building $(APP_NAME) v$(VERSION)..."
	@mkdir -p $(BUILD_DIR)
	go build -ldflags "-X main.version=$(VERSION)" -o $(BUILD_DIR)/$(APP_NAME) .

# Run application
run:
	go run main.go

# Run tests
test:
	@echo "Running tests..."
	go test -v ./...

# Run tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report: coverage.html"

# Clean build artifacts
clean:
	@echo "Cleaning..."
	rm -rf $(BUILD_DIR)
	rm -f coverage.out coverage.html
	go clean

# Run linter
lint:
	@echo "Running linter..."
	golangci-lint run

# Format code
fmt:
	@echo "Formatting code..."
	go fmt ./...
	goimports -w .

# Run go vet
vet:
	@echo "Running go vet..."
	go vet ./...

# Run all checks
all: fmt vet lint test build
	@echo "✅ All checks passed!"

# Build Docker image
docker:
	@echo "Building Docker image..."
	docker build -t $(DOCKER_IMAGE) .

# Install dependencies
deps:
	@echo "Installing dependencies..."
	go mod download
	go mod tidy