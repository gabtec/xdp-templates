BUILD_DIR:=build
SERVICE_NAME:=hello

lint: ## Run static analysis
	@golangci-lint run -E goimports

test: ## Run tests
	@go test -race -short ./...

build: ## Build the binary file
	@go build -o ${BUILD_DIR}/${SERVICE_NAME}

run:
	@go run main.go

.PHONY: lint test build run