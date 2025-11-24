# ===== Build stage =====
FROM golang:1.25.4-alpine AS builder
WORKDIR /app

RUN apk add --no-cache git ca-certificates

COPY proto ./proto

COPY user-service/go.mod user-service/go.sum ./

RUN go mod edit -replace fitness-app-microservices/proto=./proto

RUN go mod download

COPY user-service ./user-service

WORKDIR /app/user-service
RUN go build -o /app/main ./cmd/server/main.go

# ===== Run stage =====
FROM alpine:latest
WORKDIR /app

COPY --from=builder /app/main .

RUN apk add --no-cache ca-certificates

EXPOSE 50051
CMD ["./main"]
