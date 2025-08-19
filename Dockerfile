# --- Build stage ---
FROM golang:1.24.3-alpine AS build
WORKDIR /app

# Cache deps
COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY . .

# Build static linux binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -ldflags "-s -w" -o app

# --- Run stage ---
FROM gcr.io/distroless/base-debian12
WORKDIR /app

COPY --from=build /app/app /app/app

# Railway detects port via PORT or EXPOSE
EXPOSE 8080
ENV PORT=8080

# Run as non-root
USER nonroot:nonroot

ENTRYPOINT ["/app/app"]
