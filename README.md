# Turfbook Streaming

Turfbook Streaming is a Go-based backend service for live turf streaming features. The project is structured for HTTP health checks, stream handling, WebSocket-based live communication, JWT authentication, Redis support, and FFmpeg/HLS based video processing.

> Note: This repository is currently an early scaffold. Some packages are placeholders, but the config loader, health handler, and message models are already present.

## Features

- Go backend service for live streaming workflows
- Environment based configuration using `.env`
- Health check endpoint support
- WebSocket message model for chat, reactions, viewer count, stream status, ping, and pong
- JWT authentication package placeholder
- Redis configuration support
- FFmpeg and HLS output configuration support
- NestJS backend integration URL support

## Tech Stack

- Go
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
├── auth/                 # JWT/auth related logic
├── config/               # Application configuration loader
├── handlers/             # HTTP/WebSocket/stream handlers
├── hub/                  # WebSocket client and hub management
├── middleware/           # Auth and rate limit middleware
├── models/               # Shared request/response/message models
├── Dockerfile            # Docker build file, currently empty
├── go.mod                # Go module definition
├── go.sum                # Go dependency lock file
└── main.go               # Application entry point scaffold
```

## Requirements

- Go version defined in `go.mod`
- Redis, if Redis-backed features are enabled
- FFmpeg, if live stream to HLS conversion is enabled
- A NestJS backend service, if this service needs to sync stream metadata with the main API

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

## Development

Run tests:

```bash
go test ./...
```

If your environment has a read-only Go build cache, use a writable cache directory:

```bash
GOCACHE=/tmp/turfbook-go-cache go test ./...
```

Run formatting:

```bash
go fmt ./...
```

## Running The Server

The repository currently contains the service scaffold, but `main.go` is not yet wired as a runnable `package main` HTTP server. After the server bootstrap is implemented, the expected local command will be:

```bash
go run .
```

Expected base URL:

```text
http://localhost:8080
```

## Current Endpoint

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
  "timestamp": "2026-05-25T00:00:00Z"
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

## Docker

`Dockerfile` exists but is currently empty. After Docker setup is added, the expected workflow will be:

```bash
docker build -t turfbook-streaming .
docker run --env-file .env -p 8080:8080 -p 1935:1935 turfbook-streaming
```

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
- `handlers.HealthCheck()` is implemented.
- `models.IncomingMessage` and message type constants are implemented.
- `auth`, `hub`, `stream`, `websocket`, and middleware files are currently placeholders.
- `Dockerfile` is currently empty.

## License

No license file has been added yet.
