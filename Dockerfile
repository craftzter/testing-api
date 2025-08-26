
# syntax=docker/dockerfile:1
# === Stage 1: Build ===
FROM golang:1.24.6-alpine AS builder

# Install git dan ca-certificates (optional tapi sering dibutuhkan)
RUN apk add --no-cache git ca-certificates

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum first (for caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy seluruh source code
COPY . .

# Build binary
RUN go build -o main ./cmd/server/main.go

# === Stage 2: Run ===
FROM alpine:latest

# Install ca-certificates
RUN apk add --no-cache ca-certificates

WORKDIR /app

# Copy binary dari stage build
COPY --from=builder /app/main .

# Copy .env (optional, tapi kalau pakai env_file di docker-compose bisa skip ini)
COPY .env .

# Expose port sesuai APP_PORT
EXPOSE 8080

# Command to run
CMD ["./main"]
