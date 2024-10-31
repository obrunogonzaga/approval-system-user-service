# Build stage
FROM golang:1.21-alpine AS builder

# Install git and essential build tools
RUN apk add --no-cache git make build-base

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /go/bin/app ./cmd/api

# Final stage
FROM alpine:3.19

# Add necessary runtime dependencies
RUN apk --no-cache add ca-certificates tzdata

# Create non-root user
RUN adduser -D -g '' appuser

# Set working directory
WORKDIR /app

# Copy binary from builder
COPY --from=builder /go/bin/app .

# Copy config files if needed
COPY --from=builder /app/configs ./configs

# Use non-root user
USER appuser

# Expose port
EXPOSE 8080

# Command to run
ENTRYPOINT ["./app"]