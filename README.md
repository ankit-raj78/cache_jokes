# Cache Jokes - SaaS Cache Service

A Go-based HTTP server providing cache recommendations and provisioning services.

## Features

- **Request-driven recommendations**: Fast `/v1/recommend` endpoint for immediate cache strategy suggestions
- **Event-driven provisioning**: Async `/v1/provision` endpoint that returns job IDs for background cache provisioning
- **Health checks**: `/v1/health` endpoint for monitoring

## Quick Start

1. **Run the server:**
   ```bash
   go run ./cmd/server
   ```

2. **Test the endpoints:**
   ```bash
   # Health check
   curl http://localhost:8080/v1/health
   
   # Get recommendations
   curl -X POST http://localhost:8080/v1/recommend \
     -H "Content-Type: application/json" \
     -d '{}'
   
   # Provision cache (async)
   curl -X POST http://localhost:8080/v1/provision \
     -H "Content-Type: application/json" \
     -d '{"engine":"redis","pattern":"cache-aside","nodeCount":1}'
   ```

## Project Structure

```
├── cmd/server/          # Main executable
├── internal/
│   ├── api/             # HTTP router & handlers
│   ├── config/          # Configuration loader
│   ├── recommend/       # Recommendation engine
│   ├── provision/       # Provisioning interfaces & stubs
│   └── integrate/       # Integration formatter stub
├── db/migrations/       # SQL migration files
└── pkg/                 # Reusable libraries
```

## Configuration

Set the `PORT` environment variable to change the server port (default: 8080).

## Docker

Build and run with Docker:

```bash
docker build -t cache-jokes .
docker run -p 8080:8080 cache-jokes
```
