# BuzzFizz

![CI/CD](https://github.com/fatmalabidi/buzzfizz/actions/workflows/cicd.yml/badge.svg)
![Coverage](./coverage.svg)

A small configurable FizzBuzz HTTP API documented by an OpenAPI 3.0 spec.
 
**Requirements**
- Go 1.24
- Docker (optional)
- `oapi-codegen`  
- `golangci-lint`
- `mockgen` (only if you will be updating services)

to install oapi-codegen:

```bash
go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest

export PATH=$PATH:$(go env GOPATH)/bin

oapi-codegen --version 

```

Install golangci-lint:
```bash
go install github.com/golangci-lint/golangci-lint/cmd/golangci-lint@latest
```

Install codegen:
```bash
go install go.uber.org/mock/mockgen@latest
```
  
**Code Generation**
After updating `api/openapi.yml`, run:
```bash
make generate
```

DO NOT edit generated files!


**Lint**
```bash
make lint
```

**Tests**  
You need to generate mocks to be able to test handlers:

```bash
make generate-mocks
```


Run unit tests:
```bash
make test
```

Test with race
```bash
make test-race
```

Test with coverage and generate coverage badge
```bash
make test-cov
```

To see the coverage report
```bash
make cover-html
```

**Run**
To run the application on docker
```bash
make docker-run
```
To run the application locally
```bash
make run
```