# Turfbook Streaming

Turfbook Streaming is a Go-based backend service for live turf streaming features. The project is structured for HTTP health checks, stream handling, WebSocket-based live communication, JWT authentication, Redis support, and FFmpeg/HLS based video processing.

## Features

- Go backend service for live streaming workflows
- Environment based configuration using `.env`
- Health check endpoint support with JSON response
- WebSocket message model for chat, reactions, viewer count, stream status, ping, and pong
- JWT authentication package and verification claims
- Redis configuration support
- FFmpeg and HLS output configuration support
- NestJS backend integration URL support
- **Production-Ready Enhancements:**
  - **Graceful Shutdown:** Handles OS signals (`SIGINT`, `SIGTERM`) to cleanly shut down HTTP server connections.
  - **Structured Logging (`slog`):** Outputs text logs in development and structured JSON logs in production.
  - **Task Automation (`Makefile`):** Easy commands for development, testing, formatting, and dockerized linting.
  - **CI/CD Pipeline:** Fully configured GitHub Actions workflow that compiles and runs `golangci-lint` v2 and executes tests on Go 1.26.3.

## Tech Stack

- Go (1.26.3)
- Gorilla Mux
- Gorilla WebSocket
- Go Redis
- JWT
- FFmpeg
- HLS
- godotenv

## Project Structure

```text
.
├── .github/workflows/    # CI/CD pipelines (GitHub Actions)
├── auth/                 # JWT/auth related logic
├── config/               # Application configuration loader
├── handlers/             # HTTP/WebSocket/stream handlers
├── hub/                  # WebSocket client and hub management
├── middleware/           # Auth and rate limit middleware
├── models/               # Shared request/response/message models
├── Dockerfile            # Docker build file (Multi-stage build)
├── docker-compose.dev.yml# Docker compose for local development
├── docker-compose.prod.yml# Docker compose for production
├── Makefile              # Automation script for local commands
├── go.mod                # Go module definition
├── go.sum                # Go dependency lock file
└── main.go               # Application entry point with graceful shutdown & slog
```

## Requirements

- Go version defined in `go.mod` (Go 1.26.3)
- Redis, if Redis-backed features are enabled
- FFmpeg, if live stream to HLS conversion is enabled
- A NestJS backend service, if this service needs to sync stream metadata with the main API
- Docker & Docker Compose (for running containers)

## Environment Variables

Create a `.env` file in the project root. The repository already ignores `.env`, so secrets will not be committed.

```env
PORT=8080
RTMP_PORT=1935
REDIS_URL=redis://localhost:6379
JWT_SECRET=change_this_secret
NESTJS_URL=http://localhost:3000
HLS_OUTPUT_DIR=./hls
FFMPEG_PATH=ffmpeg
```

### Config Details

| Variable | Default | Description |
| --- | --- | --- |
| `PORT` | `8080` | HTTP server port |
| `RTMP_PORT` | `1935` | RTMP ingest port |
| `REDIS_URL` | `redis://localhost:6379` | Redis connection URL |
| `JWT_SECRET` | `your_jwt_secret` | Secret used for JWT verification/signing |
| `NESTJS_URL` | `http://localhost:3000` | Main NestJS backend URL |
| `HLS_OUTPUT_DIR` | `./hls` | Directory where generated HLS files will be stored |
| `FFMPEG_PATH` | `ffmpeg` | FFmpeg binary path |

## Installation

```bash
go mod download
```

## Development and Automation

All tasks are automated using the `Makefile` in the root directory.

### Common Local Commands

- **Run Hot-Reload Server (Air):**
  ```bash
  make dev
  ```
- **Run Tests:**
  ```bash
  make test
  ```
- **Format Code:**
  ```bash
  make fmt
  ```
- **Compile Application:**
  ```bash
  make build
  ```
- **Clean Build Artifacts and HLS output:**
  ```bash
  make clean
  ```

### Running the Linter (golangci-lint v2)

Since local installation of `golangci-lint` v2 may vary, you can run the exact configuration inside a Docker container:
```bash
make lint
```

---

## Docker Environments

### Local Development (with Hot-Reload & Redis Container)

To spin up the application with a hot-reload container volume-mounted to your local files along with a local Redis container:

```bash
make docker-dev-up
```

To stop the dev containers:
```bash
make docker-dev-down
```

### Production Release (Multi-Stage Build)

To build and run the production-optimized image:

```bash
make docker-prod-up
```

To stop production containers:
```bash
make docker-prod-down
```

---

## Endpoints

### Health Check

Expected route:

```http
GET /health
```

Response shape:

```json
{
  "status": "ok",
  "service": "turfbook-streaming",
  "timestamp": "2026-06-19T11:28:47Z"
}
```

The handler for this response is available in `handlers/health.go`.

## WebSocket Message Format

Client messages use this structure:

```json
{
  "type": "CHAT",
  "content": "Hello from live stream",
  "reaction": "LIKE"
}
```

Supported message types:

| Type | Purpose |
| --- | --- |
| `CHAT` | Live chat message |
| `REACTION` | Viewer reaction such as like/love/fire |
| `VIEWER_COUNT` | Viewer count update |
| `STREAM_STATUS` | Stream status update |
| `ERROR` | Error message |
| `PING` | Client/server ping |
| `PONG` | Client/server pong |

## Suggested API Design

These routes are not fully wired yet, but they match the current package layout and intended service responsibility:

```http
GET  /health
GET  /ws/streams/{streamId}
POST /streams/{streamId}/start
POST /streams/{streamId}/stop
GET  /streams/{streamId}/status
GET  /hls/{streamId}/index.m3u8
```

## Streaming Flow

1. Streamer starts a live session from the Turfbook app.
2. Main backend creates or validates the stream session.
3. Streaming service receives RTMP input.
4. FFmpeg converts RTMP input into HLS segments.
5. HLS files are written to `HLS_OUTPUT_DIR`.
6. Viewers watch the HLS stream from the app.
7. WebSocket channel handles chat, reactions, viewer count, and stream status updates.

## Production Notes

- Use a strong `JWT_SECRET`.
- Do not commit `.env`.
- Run Redis as a managed service or stable container.
- Mount `HLS_OUTPUT_DIR` as persistent/shared storage if HLS files need to survive container restarts.
- Put HLS output behind CDN or object storage for high viewer traffic.
- Restrict CORS and WebSocket origins before production release.
- Add rate limiting for chat and reaction events.
- Add structured logs for stream start, stop, FFmpeg errors, and WebSocket disconnects.

## Development Status

- `config.Load()` is implemented.
- `handlers.HealthCheck()` is implemented with error checking.
- `models.IncomingMessage` and message type constants are fully documented and implemented.
- `Dockerfile`, `docker-compose.dev.yml`, `docker-compose.prod.yml`, and `Makefile` are implemented.
- `auth`, `hub`, `stream`, `websocket`, and middleware files are currently placeholders.

## License

No license file has been added yet.
