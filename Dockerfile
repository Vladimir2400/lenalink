# Multi-stage build for LenaLink (server + seed tools)
FROM golang:1.22-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git ca-certificates tzdata

WORKDIR /app

# Copy dependency files first (for Docker layer caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build both binaries with optimizations
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /bin/server ./cmd/server
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /bin/seed ./cmd/seed

# Final stage - minimal Alpine image
FROM alpine:latest

# Install runtime dependencies
RUN apk add --no-cache ca-certificates tzdata wget

# Copy binaries from builder
COPY --from=builder /bin/server /usr/local/bin/server
COPY --from=builder /bin/seed /usr/local/bin/seed

# Set working directory
WORKDIR /app

# Health check for server (only works when running server command)
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD wget --no-verbose --tries=1 --spider http://localhost:8080/health || exit 1

# Default command: run server
CMD ["server"]
