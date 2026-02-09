# -------- Build stage --------
FROM golang:1.24-alpine AS builder

# Enable CA certs & git
RUN apk add --no-cache ca-certificates git

WORKDIR /app

# Cache deps first
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source
COPY . .

# Build statically linked binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -trimpath \
    -ldflags="-s -w -X main.version=${VERSION:-0.0.1}" \
    -o app cmd/api/main.go

# -------- Runtime stage --------
FROM gcr.io/distroless/base-debian12

WORKDIR /app

# Copy binary and certs
COPY --from=builder /app/app /app/app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Non-root user (recommended)
USER nonroot:nonroot

EXPOSE 8080

ENTRYPOINT ["/app/app"]