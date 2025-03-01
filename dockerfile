# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum to download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go binary with static linking
RUN CGO_ENABLED=0 go build -o nlp-task-runner

# Runtime stage
FROM alpine:3.18

# Install common utilities
RUN apk add --no-cache bash findutils tar gzip unzip curl procps

WORKDIR /app

# Copy the binary and .env file from the build stage
COPY --from=builder /app/nlp-task-runner .
COPY .env .

# Create a non-root user
RUN adduser -D appuser

# Switch to the non-root user
USER appuser

# Set the entrypoint to run the binary
ENTRYPOINT ["./nlp-task-runner"]
