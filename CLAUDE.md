# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a URL shortener service built in Go that provides an API for creating and managing shortened URLs. The service uses MongoDB as its database and follows a clean architecture pattern with clear separation of concerns.

## Architecture

The codebase follows a layered architecture pattern:

- **Handler Layer** (`internal/handler/`): HTTP request handlers that validate HTTP methods and parse JSON requests/responses
- **Service Layer** (`internal/service/`): Business logic layer that orchestrates between handlers and repositories
- **Repository Layer** (`internal/database/repositories/`): Data access layer that interacts with MongoDB
- **Models** (`internal/models/`): Domain entities (e.g., `Url` struct with MongoDB BSON tags)
- **DTOs** (`internal/dto/`): Data transfer objects for API requests (e.g., `CreateUrlDto`)
- **Config** (`config/`): Configuration management using Viper (note: this package exists but is not currently used in main.go)

### Dependency Flow

```
Handler -> Service -> Repository -> MongoDB
```

Dependency injection is used throughout: repositories are injected into services, and services are injected into handlers. All dependencies are initialized in `cmd/main/main.go`.

### Application Entry Point

The application starts in `cmd/main/main.go:16` which:
1. Loads `.env` configuration using Viper
2. Connects to MongoDB using the `DB_URL` environment variable
3. Creates repository, service, and handler instances
4. Registers the HTTP route `/` (POST only) for URL creation
5. Starts HTTP server on port 8000

## Development Commands

### Running the Application

```bash
# Start MongoDB using Docker Compose
docker-compose up -d

# Run the application
go run cmd/main/main.go
```

The server will start on `http://localhost:8000`.

### MongoDB

MongoDB runs in a Docker container and stores data in `docker/data/db/`. To start/stop MongoDB:

```bash
docker-compose up -d    # Start MongoDB
docker-compose down     # Stop MongoDB
```

### Dependencies

```bash
# Install/update dependencies
go mod download

# Tidy dependencies
go mod tidy
```

### Testing

There are currently no tests in the codebase. When adding tests, follow Go conventions:
- Test files should be named `*_test.go`
- Run tests with `go test ./...`
- Run tests for a specific package: `go test ./internal/service`

## Configuration

The application uses Viper for configuration management. Environment variables are loaded from a `.env` file in the project root:

- `DB_URL`: MongoDB connection string (example: `mongodb://localhost:27017`)

Note: There's a typo in `.env.example` - it says `DB_URl` instead of `DB_URL`. The application expects `DB_URL`.

## Key Implementation Details

- **Database**: MongoDB is accessed via the official `go.mongodb.org/mongo-driver/v2` driver
- **Database Name**: `url-shortener`
- **Collection Name**: `urls`
- **HTTP Framework**: Standard library `net/http` (no external router/framework)
- **Repository Context**: The repository stores a `context.Context` instance created in main.go and reuses it for all database operations
- **ID Handling**: MongoDB ObjectIDs are converted to strings by stripping the `ObjectID("...")` wrapper in `internal/database/repositories/url_repository.go:44-46`

## Current API Endpoints

- `POST /`: Create a shortened URL
  - Request body: `{"longUrl": "string", "shortCode": "string", "accessCount": number}`
  - Returns: `Url` model as JSON

## Known Issues/TODOs

- The `Update()` method in `UrlRepository` is not implemented (panics with "implement me")
- The `UrlRepositoryInterface` in `internal/database/repositories/interface/url_repository_interface.go` has a signature mismatch with the actual `UrlRepository` implementation
- No URL retrieval/redirect functionality is implemented yet
- Error handling is minimal (mostly just printing errors)
- No validation of URL format or shortCode uniqueness
