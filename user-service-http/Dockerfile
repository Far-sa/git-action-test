# Stage 1: Build the Go application
FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /app/user-service

# Debugging step: List contents of /app
RUN ls -la /app

# Stage 2: Final image (uses slim alpine image)
FROM alpine:latest

WORKDIR /root/

# Copy only the user-service binary from the builder stage
COPY --from=builder /app/user-service .

# Debugging step: List contents of /root
RUN ls -la /root

EXPOSE 8080

CMD ["./user-service"]
