# Rate Limiter Service

A high-performance, configurable rate limiter service built in Go with support for multiple algorithms and storage backends.

## Features

- **Multiple Rate Limiting Algorithms**

  - Token Bucket
  - Fixed Window
  - Sliding Window
  - Leaky Bucket

- **Storage Backends**

  - In-memory (for single instance)
  - Redis (for distributed systems)
  - PostgreSQL (for persistent storage)

- **Monitoring & Observability**
  - Prometheus metrics
  - Structured logging
  - Health checks

## Project Structure

```
rate-limiter/
├── cmd/ratelimiter/         # Main service entry point
├── internal/                # Private application code
│   ├── api/                 # HTTP/gRPC handlers
│   ├── config/              # Configuration loading
│   ├── limiter/             # Core rate limiter logic
│   │   └── algorithms/      # Algorithm implementations
│   ├── middleware/          # HTTP middleware
│   ├── storage/             # Backend integrations
│   └── metrics/             # Metrics & tracing
├── pkg/ratelimiter/         # Public packages
├── configs/                 # Configuration files
├── test/                    # Tests
└── scripts/                 # Helper scripts
```

## Quick Start

1. Clone the repository
2. Copy and modify the configuration file:
   ```bash
   cp configs/config.yaml configs/local.yaml
   ```
3. Run the service:
   ```bash
   go run cmd/ratelimiter/main.go
   ```

## Docker

Build and run with Docker:

```bash
docker build -t rate-limiter .
docker run -p 8080:8080 -p 9090:9090 rate-limiter
```

## Configuration

See `configs/config.yaml` for available configuration options.

## Development

TODO: Add development setup instructions

## Testing

TODO: Add testing instructions

## Contributing

TODO: Add contributing guidelines

## License

TODO: Add license information
