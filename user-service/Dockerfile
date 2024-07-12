FROM golang:1.22-alpine AS builder
WORKDIR /app

# Copy go.mod and go.sum files of user-service
COPY ./user-service/go.mod ./user-service/go.sum ./

# Copy common go.mod and go.sum files
COPY ./common/go.mod ./common/go.sum ./common/

# Download dependencies
RUN go mod download

# Copy the common module and the source code of user-service
COPY ./common ./common
COPY ./user-service .

# Build the service
RUN go build -o user_service ./user_service.go

# Final stage
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/user_service .
CMD ["./user_service"]


# # Stage 1: Build the binary
# FROM golang:1.22-alpine AS builder

# # Set the working directory inside the container
# WORKDIR /app

# # Copy go.mod and go.sum files
# COPY go.mod go.sum ./

# # Copy the common module
# COPY common ./common

# # Download all dependencies
# RUN go mod download

# # Copy the rest of the application code
# COPY . .

# # Build the service
# RUN go build -o user_service ./user_service.go

# # Final stage
# FROM alpine:latest

# WORKDIR /root/

# COPY --from=builder /app/user_service .

# CMD ["./user_service"]
