FROM golang:1.24 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o api ./cmd/api

FROM alpine:3.20

WORKDIR /app
COPY --from=builder /app/api .

EXPOSE 8080

ENTRYPOINT ["./api"]
