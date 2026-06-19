.PHONY: dev test fmt build docker-dev-up docker-dev-down docker-prod-up docker-prod-down lint clean

# Default: start local hot-reload dev server
dev:
	air

# Run Go tests with writable cache location
test:
	GOCACHE=/tmp/turfbook-go-cache go test -v -race -cover ./...

# Run Go formatter
fmt:
	go fmt ./...
	goimports -w . 2>/dev/null || true

# Build local binary
build:
	go build -o tmp/main main.go

# Start Docker Dev environment (including Redis)
docker-dev-up:
	docker compose -f docker-compose.dev.yml up -d --build

# Stop Docker Dev environment
docker-dev-down:
	docker compose -f docker-compose.dev.yml down

# Start Docker Prod environment
docker-prod-up:
	docker compose -f docker-compose.prod.yml up -d --build

# Stop Docker Prod environment
docker-prod-down:
	docker compose -f docker-compose.prod.yml down

# Run golangci-lint inside a Docker container (avoids local installation)
lint:
	docker run --rm -v $(shell pwd):/app -w /app golangci/golangci-lint:latest golangci-lint run -v

# Clean temporary files and build artifacts
clean:
	rm -rf tmp/* hls/*
