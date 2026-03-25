.PHONY: generate lint build run test test-race test-cov clean-coverage docker-build docker-run
APP_NAME=fizzbuzz-api
BINARY=bin/$(APP_NAME)

COVERAGE_FILE := coverage.out
COVERAGE_THRESHOLD ?= 70
COVER_DIR := .coverage
COVER_RAW := $(COVER_DIR)/coverage.raw.out
COVER_REPORT := $(COVER_DIR)/coverage.txt
COVER_BADGE := coverage.svg
COVER_EXCLUDE_REGEX := '(\.gen\.go:|/api/|/cmd/)'
COVER_PKGS := $(shell go list ./... | grep -Ev '/(api|cmd|internal/mocks)($$|/)')

DOCKER_IMAGE ?= $(APP_NAME):latest
DOCKER_CONTAINER ?= $(APP_NAME)
DOCKER_PORT ?= 8080:8080

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

lint:
	@echo "Running golangci-lint..."
	golangci-lint run

build:
	go build -o $(BINARY) ./cmd/api

run: build
	@echo "Running application..."
	./$(BINARY)
	
generate-mocks:
	go generate ./internal/...

test:
	@echo "Running tests..."
	go test ./...

test-race:
	@echo "Running tests with race detector..."
	CGO_ENABLED=1 go test -race ./...

test-cov:
	@mkdir -p $(COVER_DIR)
	go test $(COVER_PKGS) -coverprofile=$(COVER_RAW) -covermode=atomic
	@grep -vE $(COVER_EXCLUDE_REGEX) $(COVER_RAW) > $(COVERAGE_FILE)
	@go tool cover -func=$(COVERAGE_FILE) | tee $(COVER_REPORT)
	@total=$$(go tool cover -func=$(COVERAGE_FILE) | awk '/total:/ {print $$3}' | sed 's/%//'); \
	echo "Total coverage: $$total%"; \
	color=$$(awk 'BEGIN {if ('"$$total"' >= $(COVERAGE_THRESHOLD)) print "brightgreen"; else if ('"$$total"' >= 60) print "yellow"; else print "red"}'); \
	curl -s "https://img.shields.io/badge/coverage-$${total}%25-$${color}.svg" -o $(COVER_BADGE); \
	echo "Generated $(COVER_BADGE)"

cover-html:
	go tool cover -html=$(COVERAGE_FILE) -o $(COVER_DIR)/coverage.html
	@echo "Generated HTML coverage report at $(COVER_DIR)/coverage.html"

clean-coverage:
	@rm -rf $(COVER_DIR) $(COVERAGE_FILE) $(COVER_BADGE)

docker-build:
	@echo "Building Docker image..."
	docker build -t $(DOCKER_IMAGE) .

docker-run: docker-build
	@echo "Running Docker container..."
	@docker rm -f $(DOCKER_CONTAINER) >/dev/null 2>&1 || true
	docker run -d --name $(DOCKER_CONTAINER) -p $(DOCKER_PORT) $(DOCKER_IMAGE)
	@echo "Container $(DOCKER_CONTAINER) is running on port $(DOCKER_PORT)"
