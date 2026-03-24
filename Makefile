.PHONY: generate build run lint test test-race test-cov docker-build docker-run , docker-up , clean

APP_NAME=fizzbuzz-api
BINARY=bin/$(APP_NAME)

generate:
	@echo "Generating models..."
	oapi-codegen \
		-generate models \
		-package api \
		-o ./internal/api/types.gen.go \
		./api/openapi.yml

	@echo "Generating server..."
	oapi-codegen \
		-generate std-http-server \
		-package api \
		-o ./internal/api/server.gen.go \
		./api/openapi.yml

build:
	go build -o $(BINARY) ./cmd/api

run: build
	@echo "Running application..."
	./$(BINARY)

test:
	@echo "Running tests..."
	go test ./...

test-race:
	@echo "Running tests with race detector..."
	CGO_ENABLED=1 go test -race ./...

lint:
	@echo "Running golangci-lint..."
	golangci-lint run

docker-build:
	@echo "Building Docker image..."
	docker build -t $(APP_NAME) .

docker-run:
	@echo "Running Docker container..."
	docker run -p 8080:8080 $(APP_NAME)

docker-up: docker-build docker-run

clean:
	@echo "Cleaning build artifacts..."
	rm -f $(BINARY)
