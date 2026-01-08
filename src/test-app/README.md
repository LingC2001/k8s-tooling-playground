# test-app

A minimal test application with a simple health check endpoint, structured logging, rate limiting, and Swagger documentation.

## Features

- ✅ Simple health check endpoint
- ✅ IP-based rate limiting middleware
- ✅ Structured logging with slog
- ✅ Swagger/OpenAPI documentation
- ✅ Lightweight and easy to extend

## Prerequisites

- Go 1.25+
- Docker (for containerization)
- Make (optional, for build commands)

## Getting Started

### 1. Run Locally

```bash
# Install dependencies
go mod download

# Run the application
make run
# or
go run main.go
```

The app will start on `http://localhost:8000`.

### 2. Build Docker Image

```bash
# Build the image
docker build -t test-app:latest .

# Load into Kind for local testing
kind load docker-image test-app:latest
```

## API Endpoints

### Health Check

```bash
curl http://localhost:8000/api/v1/health
```

Response:
```json
{
  "status": "OK",
  "app": "test-app"
}
```

## Swagger Documentation

View the interactive API documentation at:

```
http://localhost:8000/swagger/index.html
```

### Regenerate Swagger Docs

If you modify endpoints or add swagger comments, regenerate the docs:

```bash
~/go/bin/swag init
# or add ~/go/bin to PATH and run:
swag init
```

## Rate Limiting

The app includes IP-based rate limiting:
- **Limit:** 10 requests/second
- **Burst:** 20 requests

Exceeding the limit returns a 429 status code:
```json
{
  "error": "rate limit exceeded"
}
```

## Development

### Code Structure

```
test-app/
├── main.go                      # Application entry point
├── go.mod / go.sum              # Dependency management
├── Dockerfile                   # Container build config
├── Makefile                     # Build commands
├── docs/                        # Swagger documentation (auto-generated)
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
└── internal/
    ├── health/
    │   ├── handler.go           # Health check handler with swagger docs
    │   └── routes.go            # Route setup
    ├── middleware/
    │   └── rate_limit.go        # IP-based rate limiting
    └── models/
        └── shared.go            # Shared response models
```

### Linting & Formatting

```bash
make lint       # Run all linters
make fmt        # Format code
make vet        # Run go vet
make test       # Run tests
```

## Deployment

See the main playground repository for Kubernetes deployment instructions using Argo CD.
