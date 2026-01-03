# Student CRUD API - Go Backend

A production-ready RESTful API for managing student records using Go's standard library, SQLite, and structured configuration. Demonstrates clean architecture with graceful shutdown and full CRUD operations.

## Features
- **Full CRUD**: Create, read (list/by ID), update, delete students via standard REST endpoints.
- **SQLite Persistence**: Embedded DB with custom storage layer for easy setup.
- **Structured Logging**: `log/slog` for production-grade logs with context.
- **Graceful Shutdown**: Handles SIGINT/SIGTERM with 5s timeout context.
- **Config-Driven**: YAML-based config for env-specific settings.
- **No Dependencies**: Pure stdlib `net/http` for routing and serving.

## Tech Stack
- **Runtime**: Go 1.21+ (stdlib only)
- **Database**: SQLite via `database/sql`
- **Config**: YAML parsing
- **Logging**: `log/slog`
- **Server**: `net/http` with `ServeMux`

## Architecture Flow
1. Client → `http.Server` → `ServeMux` Router
2. Route → Handler (e.g. `student.New(storage)`)
3. Handler → `SQLite Storage` → DB Query
4. Response → JSON (200/404/500)

Config/DB init → ListenAndServe → Signal shutdown.

## Installation
1. Ensure Go 1.21+ installed
2. Copy files to project dir
3. `go mod init student-api && go mod tidy`
4. Edit `local.yaml` for DB path/server addr

## Environment Variables
Uses `local.yaml`:
```yaml
env: local
addr: ":8080"
db_path: "./students.db"
Running
text
go run main.go
Logs: slog.Info("server started", slog.String("address", ":8080"))

Testing (curl/Postman)
text
# Create
curl -X POST http://localhost:8080/api/students -H "Content-Type: application/json" -d '{"name":"Alice","email":"alice@example.com"}'

# List
curl http://localhost:8080/api/students

# Get ID=1
curl http://localhost:8080/api/students/1

# Update
curl -X PUT http://localhost:8080/api/students/1 -H "Content-Type: application/json" -d '{"name":"Alice Updated"}'

# Delete
curl -X DELETE http://localhost:8080/api/students/1
Stop: Ctrl+C (graceful).

Project Structure
text
student-api/
├── main.go      # Entry: config → DB → router → server
├── conf.go      # YAML config loader
├── local.yaml   # Config file
├── student.go   # Models + handlers
├── types.go     # Shared types
├── response.go  # JSON helpers
├── sqlite.go    # DB init/queries
└── storage.go   # Interface
How It Works
main() loads config via config.MustLoad()

Inits SQLite: sqlite.New(cfg)

Sets up http.NewServeMux() with handlers

server.ListenAndServe() blocks

Signals → server.Shutdown(ctx)

Why These Technologies?
Go net/http: Fast, lightweight—no framework deps

SQLite: Zero-config, file-based persistence

YAML: Readable multi-env config

Slog: Structured prod logging

Graceful Shutdown: Prevents corruption on restarts
