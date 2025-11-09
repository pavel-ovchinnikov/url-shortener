# ENV
BINARY_NAME := app
SRC := ./...
BIN_DIR := .bin
NAMESPACE := url-shortener
APP_NAME := url-shortener

.PHONY: generate build run test lint clean

generate:
	@echo "Nothing to generate yet"

build:
	@echo "🔨 Building..."
	go build -o ${BIN_DIR}/$(BINARY_NAME) cmd/main.go

run: build
	@echo "🚀 Running..."
	./${BIN_DIR}/$(BINARY_NAME)

test:
	@echo "🧪 Running tests..."
	go test -v $(SRC)

lint:
	@echo "🔍 Linting..."
	@if ! command -v golangci-lint >/dev/null 2>&1; then \
		echo "⚠️  golangci-lint not found. Installing..."; \
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest; \
	fi
	golangci-lint run ./...

clean:
	@echo "🧹 Cleaning..."
	rm -rf ${BIN_DIR}

deploy-docker:
	@echo "Nothing to deploy yet"


docker-build:
	docker-compose -f deploy/docker/docker-compose.yml build

docker-up: docker-build
	docker-compose -f deploy/docker/docker-compose.yml up -d

docker-down:
	docker-compose -f deploy/docker/docker-compose.yml down

k8s-deploy:
	docker save backend:latest | nerdctl --namespace k8s.io load
	kubectl create namespace ${NAMESPACE} || true
	kubectl apply -k deploy/k8s --namespace=${NAMESPACE} || true
	kubectl get svc ${APP_NAME} --namespace=${NAMESPACE}

k8s-delete:
	@echo "Nothing to delete yet"
	kubectl delete namespace ${NAMESPACE} || true
