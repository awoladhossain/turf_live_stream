# --- Stage 1: Build the Go binary ---
FROM golang:1.26.3-alpine AS builder

# Git and GCC/musl-dev installation
RUN apk add --no-cache git gcc musl-dev

WORKDIR /app

# Cache Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Statically compile the binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o turfbook-streaming main.go

# --- Stage 2: Final Run Image ---
FROM alpine:3.19

# Install ca-certificates and FFmpeg
RUN apk add --no-cache ffmpeg ca-certificates

WORKDIR /app

# Copy compiled binary
COPY --from=builder /app/turfbook-streaming .
RUN mkdir -p /app/hls

# Expose API and RTMP ports
EXPOSE 8080
EXPOSE 1935

# Command to run the application
CMD ["./turfbook-streaming"]
