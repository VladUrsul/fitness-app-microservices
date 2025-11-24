# ===== Build stage =====
FROM golang:1.25.4-alpine AS builder
WORKDIR /app

RUN apk add --no-cache git ca-certificates

COPY proto ./proto

COPY api-gateway/go.mod api-gateway/go.sum ./

RUN go mod edit -replace fitness-app-microservices/proto=./proto

RUN go mod download

COPY api-gateway ./api-gateway

WORKDIR /app/api-gateway
RUN go build -o /app/main ./cmd/gateway/main.go

# ===== Run stage =====
FROM alpine:latest
WORKDIR /app

COPY --from=builder /app/main .

RUN apk add --no-cache ca-certificates

EXPOSE 8080
CMD ["./main"]
