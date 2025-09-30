# Build stage
FROM golang:1.21-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git make gcc musl-dev

# Set working directory
WORKDIR /build

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the binary
RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o relay ./cmd/relay

# Runtime stage
FROM alpine:latest

# Install runtime dependencies
RUN apk --no-cache add ca-certificates

# Create non-root user
RUN addgroup -g 1000 relay && \
    adduser -D -u 1000 -G relay relay

# Set working directory
WORKDIR /app

# Copy binary from builder
COPY --from=builder /build/relay /app/relay
COPY --from=builder /build/config.yaml /app/config.yaml

# Change ownership
RUN chown -R relay:relay /app

# Switch to non-root user
USER relay

# Expose ports
# Private DHT
EXPOSE 4001/tcp
# Public DHT
EXPOSE 4002/tcp
# Metrics
EXPOSE 8080/tcp
# Health
EXPOSE 8081/tcp

# Health check
HEALTHCHECK --interval=30s --timeout=10s --start-period=10s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:8081/health || exit 1

# Run the binary
ENTRYPOINT ["/app/relay"]
CMD ["-config", "/app/config.yaml"]