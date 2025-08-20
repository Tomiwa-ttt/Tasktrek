# Use official Golang image
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY . .

# Build the binary
RUN go build -o main .

# Run stage
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .

# Set PORT env (Railway/Heroku will override automatically)
ENV PORT=8080

EXPOSE 8080

CMD ["./main"]
