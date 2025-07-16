# -------------------
# Stage 1: Builder
# -------------------
FROM rustlang/rust:nightly-bullseye AS builder

WORKDIR /usr/src/app

# Install build-time dependencies required for sqlx and rdkafka
RUN apt-get update && apt-get install -y --no-install-recommends \
    build-essential \
    pkg-config \
    librdkafka-dev \
    libsasl2-dev \
    libssl-dev

# Copy your manifests
COPY Cargo.toml Cargo.lock ./

# Build only the dependencies to leverage Docker layer caching.
# We create a dummy main.rs to make Cargo happy.
RUN mkdir src && echo "fn main() {}" > src/main.rs
RUN cargo build --release && rm -f src/main.rs

# Copy your actual source code
COPY ./src ./src

# Build your application
RUN cargo build --release

# -------------------------
# Stage 2: Final runtime
# -------------------------
FROM debian:bullseye-slim

# Create a non-root user for safety
RUN useradd -m appuser

# Install libpq (Postgres driver dependency)
RUN apt-get update && apt-get install -y libpq5 && rm -rf /var/lib/apt/lists/*

# Copy the compiled binary from the builder
COPY --from=builder /usr/src/app/target/release/inventory-read /usr/local/bin/inventory-read

# Switch to the non-root user
USER appuser

# Set environment to production
ENV RUST_LOG=info

# Expose the port
EXPOSE 3000

# Run the binary
CMD ["inventory-read"]
