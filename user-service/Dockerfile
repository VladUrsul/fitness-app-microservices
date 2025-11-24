# Build stage
FROM golang:1.25.4-alpine AS builder
WORKDIR /app

# Copy proto module
COPY ../proto ./proto

# Copy this service's go.mod/go.sum
COPY user-service/go.mod user-service/go.sum ./

# Download dependencies
RUN go mod download

# Copy service source code
COPY . .

# Build binary
RUN go build -o main ./cmd/server

# Run stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/internal/docs ./docs

EXPOSE 50051  
CMD ["./main"]
