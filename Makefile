.PHONY: generate

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
