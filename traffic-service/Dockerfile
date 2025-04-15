FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./
# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /traffic-service ./cmd/server

# Use a smaller image for the final container
FROM alpine:3.18

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

# Copy the binary from the build stage
COPY --from=builder /traffic-service .

# Create a non-root user and change ownership
RUN adduser -D appuser && \
    chown -R appuser:appuser /app

# Switch to non-root user
USER appuser

# Expose the service port
EXPOSE 8083

# Command to run
CMD ["./traffic-service"]