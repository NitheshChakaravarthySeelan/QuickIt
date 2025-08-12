# Stage 1: Builder/Test stage
# This stage has the Go toolchain, dependencies, and source code.
# It's used to build the application and to run tests.
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy module files and download dependencies first to leverage Docker cache
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go app. This binary is for the final stage.
# We build it here to ensure the code is buildable before proceeding.
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main ./cmd/main.go

# Stage 2: Final production stage
# This stage contains only the compiled binary for a small image size.
FROM alpine:latest

WORKDIR /app

# Copy the pre-built binary from the builder stage
COPY --from=builder /app/main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]