# Start from the official Golang base image
FROM golang:1.20.5 AS development

# Add a work directory
WORKDIR /app

# Copy go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Clear the Go module cache
RUN go clean -modcache

# Download Go module dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Expose the port that the application listens on
EXPOSE 8080

# Set the entry point
CMD go run ./cmd/main.go



