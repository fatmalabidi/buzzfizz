# Buzzfizz (Fizzbuzz API)

A small configurable FizzBuzz HTTP API documented by an OpenAPI 3.0 spec.
 
**Requirements**
- Go 1.24
- Docker (optional)
- `oapi-codegen`  

to install oapi-codegen:

```bash
go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest

export PATH=$PATH:$(go env GOPATH)/bin

oapi-codegen --version 

```
  
**Code Generation**
After updating `api/openapi.yml`, run:
```bash
make generate
```

DO NOT edit generated files!

**Tests**  
Run unit tests:
```bash
make test
```

Test with race
```bash
make test-race
```

**Docker**
```bash
make docker-run
```

**Cleanup**
```bash
make clean
```


