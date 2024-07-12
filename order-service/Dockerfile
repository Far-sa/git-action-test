# /order/Dockerfile
# Stage 1: Build the proto files
FROM golang:1.20-alpine AS proto-builder

WORKDIR /proto

# Copy the proto files
COPY ../proto /proto

# Install dependencies
RUN apk add --no-cache git \
    && go install google.golang.org/protobuf/cmd/protoc-gen-go@latest \
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest \
    && go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest

# Generate the proto files
RUN mkdir -p /generated/proto \
    && protoc -I /proto --go_out=/generated/proto --go-grpc_out=/generated/proto /proto/*.proto \
    && protoc -I /proto --grpc-gateway_out=/generated/proto /proto/*.proto

# Initialize the Go module in the generated directory
COPY ../generated/go.mod /generated/
RUN cd /generated && go mod tidy

# Stage 2: Build the service
FROM golang:1.20-alpine

WORKDIR /app

# Copy the generated proto files and go.mod
COPY --from=proto-builder /generated /generated

# Copy the source code
COPY . .

# Install dependencies
RUN apk add --no-cache git \
    && go mod tidy

# Build the service
RUN go build -o order ./main.go

# Run the service
CMD ["./order"]


#!
# FROM golang:1.22-alpine AS builder
# WORKDIR /app

# # Copy go.mod and go.sum files of order-service
# COPY ./order-service/go.mod ./order-service/go.sum ./

# # Copy common go.mod and go.sum files
# COPY ./common/go.mod ./common/go.sum ./common/

# # Download dependencies
# RUN go mod download

# # Copy the common module and the source code of order-service
# COPY ./common ./common
# COPY ./order-service .

# # Build the service
# RUN go build -o order_service ./order_service.go

# # Final stage
# FROM alpine:latest
# WORKDIR /root/
# COPY --from=builder /app/order_service .
# CMD ["./order_service"]



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
# RUN go build -o order_service ./order_service.go

# # Final stage
# FROM alpine:latest

# WORKDIR /root/

# COPY --from=builder /app/order_service .

# CMD ["./order_service"]
