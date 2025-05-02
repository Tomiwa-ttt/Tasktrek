# Use Go 1.24.2 as base image
FROM golang:1.24.2

# Set working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files first for dependency caching
COPY go.mod go.sum ./

# Download Go modules
RUN go mod download

# Copy the rest of your source code
COPY . .

# Build the Go app
RUN go build -tags netgo -ldflags '-s -w' -o app

# Expose port (change if your app uses a different port)
EXPOSE 8080

# Command to run the executable
CMD ["./app"]
