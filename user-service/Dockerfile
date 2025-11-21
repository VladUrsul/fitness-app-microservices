# Build stage
FROM golang:1.25.4-alpine AS builder
WORKDIR /app

# Download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the binary
RUN go build -o main ./cmd/server

# Run stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/internal/docs ./docs

# Expose service port
EXPOSE 8081

# Start the service
CMD ["./main"]
