# ENV
BINARY_NAME := app
SRC := ./...
BIN_DIR := .bin

.PHONY: generate build run test lint clean

generate:
	@echo "Nothing to generate yet"

build:
	@echo "🔨 Building..."
	go build -o ${BIN_DIR}/$(BINARY_NAME) cmd/main.go
	@echo "🔨 Done"

run: build
	@echo "🚀 Running..."
	./${BIN_DIR}/$(BINARY_NAME)
	@echo "🚀 Done"

test:
	@echo "🧪 Running tests..."
	go test -v $(SRC)
	@echo "🧪 Done"

lint:
	@echo "🔍 Linting..."
	@if ! command -v golangci-lint >/dev/null 2>&1; then \
		echo "⚠️  golangci-lint not found. Installing..."; \
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest; \
	fi
	golangci-lint run ./...
	@echo "🔍 Done"

clean:
	@echo "🧹 Cleaning..."
	rm -rf ${BIN_DIR}
	@echo "🧹 Done"

deploy-docker:
	@echo "Nothing to deploy yet"

# IMAGE_NAME=go-server

docker-build:
	docker-compose -f deploy/docker/docker-compose.yml build

docker-up: docker-build
	docker-compose -f deploy/docker/docker-compose.yml up -d

docker-down:
	docker-compose -f deploy/docker/docker-compose.yml down