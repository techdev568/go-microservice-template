# ===== Build Stage =====
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build binary
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server

# ===== Run Stage =====
FROM alpine:latest

WORKDIR /root/

# Copy binary from builder
COPY --from=builder /app/server .

# Copy example env (optional)
COPY .env.example .env

# Expose port
EXPOSE 8080

# Run binary
CMD ["./server"]
