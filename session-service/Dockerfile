# ===== Build stage =====
FROM golang:1.25.4-alpine AS builder
WORKDIR /app

RUN apk add --no-cache git ca-certificates

COPY proto ./proto

COPY session-service/go.mod session-service/go.sum ./

RUN go mod edit -replace fitness-app-microservices/proto=./proto

RUN go mod download

COPY session-service ./session-service

WORKDIR /app/session-service
RUN go build -o /app/main ./cmd/server/main.go

# ===== Run stage =====
FROM alpine:latest
WORKDIR /app

COPY --from=builder /app/main .

RUN apk add --no-cache ca-certificates

EXPOSE 50053
CMD ["./main"]
