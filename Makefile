# Project variables
APP_NAME=GoPost
CMD_DIR=./cmd/api
BIN_DIR=./bin
DOCKER_COMPOSE_FILE=docker-compose.yaml

-include .env
export

LOCAL_POSTGRES_URL=postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:5432/${POSTGRES_DB}?sslmode=disable

# Go build flags
GO_BUILD=cd $(CMD_DIR) && go build -o ../../$(BIN_DIR)/$(APP_NAME)

# Default target
.PHONY: all
all: build run

# Build the Go binary
.PHONY: build
build: wire
	@echo "🚧 Building $(APP_NAME)..."
	@mkdir -p $(BIN_DIR)
	@rm -f $(BIN_DIR)/$(APP_NAME)
	@cd $(CMD_DIR) && go build -o ../../$(BIN_DIR)/$(APP_NAME) .
	@echo "✅ Build complete: $(BIN_DIR)/$(APP_NAME)"

# Run the application with Docker
.PHONY: run
run: build docker-up
	@echo "🚀 Running $(APP_NAME) on host..."
	@trap 'echo "🛑 Caught signal, stopping..."; kill $$pid 2>/dev/null; make docker-down; exit 0' INT TERM EXIT; \
	export POSTGRES_URL=$(LOCAL_POSTGRES_URL); \
	$(BIN_DIR)/$(APP_NAME) & pid=$$!; \
	wait $$pid || true


.PHONY: air
air: build docker-up
	@echo "🚀 Running  $(APP_NAME) with Air..."
	@trap 'make docker-down' INT TERM EXIT; \
	export POSTGRES_URL=$(LOCAL_POSTGRES_URL); \
	air & \
	wait $$! || true


# Run using go run (skips build)
.PHONY: dev
dev: wire docker-up
	@echo "💻 Running $(APP_NAME) in dev mode..."
	export POSTGRES_URL=$(LOCAL_POSTGRES_URL); \
	@go run $(CMD_DIR)/main.go

# Bring up Docker containers
.PHONY: docker-up
docker-up:
	@echo "🐳 Starting Docker containers..."
	@docker compose -f $(DOCKER_COMPOSE_FILE) up -d

# Stop Docker containers
.PHONY: docker-down
docker-down:
	@echo "🛑 Stopping Docker containers..."
	@docker compose -f $(DOCKER_COMPOSE_FILE) down

# Run tests
.PHONY: test
test:
	@echo "🧪 Running tests..."
	export POSTGRES_URL=$(LOCAL_POSTGRES_URL); \
	@go test ./... -v

# Run lint checks (if golangci-lint installed)
.PHONY: lint
lint:
	@echo "🔍 Running linter..."
	@golangci-lint run || true

# Clean up binaries and temp files
.PHONY: clean
clean:
	@echo "🧹 Cleaning up..."
	@rm -rf $(BIN_DIR)
	@go clean

.PHONY: wire
wire:
	@wire $(CMD_DIR)