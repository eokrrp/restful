# Use a base image with Go and Alpine Linux pre-installed
FROM golang:1.21-alpine AS builder

# Set the current working directory inside the container
WORKDIR /app

# Copy the Go modules and build files
COPY go.mod go.sum ./

# Download and install Go dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# Start a new stage from Alpine Linux base image
FROM alpine:latest

# Set the working directory in the container
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/app .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./app"]
