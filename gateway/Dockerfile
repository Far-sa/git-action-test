# Stage 1: Build the binary
FROM golang:1.22-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go app
RUN go build -o gateway .

# Stage 2: Run the binary
FROM alpine:latest

# Copy the binary from the builder stage
COPY --from=builder /app/gateway /gateway

# Command to run the executable
CMD ["/gateway"]
