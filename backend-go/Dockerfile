# --- Builder stage ---
FROM golang:1.23 AS builder

# Set working directory
WORKDIR /app

# Install make
RUN apt-get update && apt-get install -y make

# Copy everything and build the binary
COPY . .
RUN make build

# --- Final stage ---
FROM debian:bookworm-slim

# Install CA certificates required for Pinata API Calls
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Copy only the compiled binary
COPY --from=builder /app/bin/koala-world-backend .

# Copy the config file
COPY ./config/config.dev.yaml ./config/config.dev.yaml

# Set executable permission just in case
RUN chmod +x ./koala-world-backend

# Expose backend port
EXPOSE 8085

# Run the binary
CMD ["./koala-world-backend"]
	