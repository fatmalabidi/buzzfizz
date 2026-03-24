.PHONY: generate build run test test-race test-cov docker-build docker-run clean

APP_NAME=fizzbuzz-api
BINARY=bin/$(APP_NAME)

# Génération OpenAPI
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
	@echo "Building application..."
	go build -o $(BINARY) ./cmd/server

run: build
	@echo "Running application..."
	./$(BINARY)

test:
	@echo "Running tests..."
	go test ./...

test-race:
	@echo "Running tests with race detector..."
	CGO_ENABLED=1 go test -race ./...

test-cov:
	@echo "Running tests with coverage..."
	go test ./... -coverprofile=coverage.out
	go tool cover -func=coverage.out

test-cov-html: test-cov
	go tool cover -html=coverage.out -o coverage.html

docker-build:
	@echo "Building Docker image..."
	docker build -t $(APP_NAME) .

docker-run:
	@echo "Running Docker container..."
	docker run -p 8080:8080 $(APP_NAME)

clean:
	@echo "Cleaning..."
	rm -rf bin/ coverage.out coverage.html

clean-docker:
	@echo "Removing containers for $(APP_NAME)..."
	-docker ps -aq --filter "ancestor=$(APP_NAME)" | xargs -r docker rm -f

	@echo "Removing image $(APP_NAME)..."
	-docker rmi -f $(APP_NAME) 2>/dev/null || true

	@echo "Cleanup done"