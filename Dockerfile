# Stage 1: Build the application
FROM golang:1.18-alpine AS build

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download necessary Go modules
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN go build -o chatservice ./cmd/chatservice

# Stage 2: Run the application
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the built binary from the build stage
COPY --from=build /app/chatservice .

# Expose the service port
EXPOSE 8080

# Run the binary
CMD ["./chatservice"]