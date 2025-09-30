.PHONY: all build run test clean docker-build docker-run docker-stop

# Variables
BINARY_NAME=relay
DOCKER_IMAGE=p2p-relay:latest
GO=go
GOFLAGS=-v

# Default target
all: test build

# Build the binary
build:
	$(GO) build $(GOFLAGS) -o $(BINARY_NAME) ./cmd/relay

# Run the application
run: build
	./$(BINARY_NAME)

# Run with custom config
run-local: build
	./$(BINARY_NAME) -config config.local.yaml

# Run tests
test:
	$(GO) test $(GOFLAGS) ./...

# Run tests with coverage
test-coverage:
	$(GO) test -cover -coverprofile=coverage.out ./...
	$(GO) tool cover -html=coverage.out -o coverage.html

# Run tests with race detection
test-race:
	$(GO) test -race $(GOFLAGS) ./...

# Format code
fmt:
	$(GO) fmt ./...

# Run linter
lint:
	golangci-lint run

# Download dependencies
deps:
	$(GO) mod download
	$(GO) mod tidy

# Clean build artifacts
clean:
	rm -f $(BINARY_NAME)
	rm -f coverage.out coverage.html

# Docker commands
docker-build:
	docker build -t $(DOCKER_IMAGE) .

docker-run:
	docker-compose up -d

docker-stop:
	docker-compose down

docker-logs:
	docker-compose logs -f

# Development setup
dev-setup: deps
	@echo "Installing development tools..."
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@echo "Development environment ready!"

# Help target
help:
	@echo "Available targets:"
	@echo "  make build         - Build the binary"
	@echo "  make run           - Build and run the application"
	@echo "  make test          - Run tests"
	@echo "  make test-coverage - Run tests with coverage report"
	@echo "  make test-race     - Run tests with race detection"
	@echo "  make fmt           - Format code"
	@echo "  make lint          - Run linter"
	@echo "  make clean         - Clean build artifacts"
	@echo "  make docker-build  - Build Docker image"
	@echo "  make docker-run    - Run with Docker Compose"
	@echo "  make docker-stop   - Stop Docker containers"
	@echo "  make dev-setup     - Setup development environment"