# Build stage
FROM golang:1.25.4-alpine AS builder
WORKDIR /app

# Copy proto module
COPY ../proto ./proto

# Copy go.mod/go.sum
COPY session-service/go.mod session-service/go.sum ./

# Download dependencies
RUN go mod download

# Copy gateway source
COPY . .

# Build binary
RUN go build -o main ./cmd/gateway/main.go

# Run stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 8080
CMD ["./main"]
