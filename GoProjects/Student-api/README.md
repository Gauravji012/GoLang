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
- **Runtime**: Go 1.21+ (stdlib only).
- **Database**: SQLite via `database/sql`.
- **Config**: YAML parsing.
- **Logging**: `log/slog`.
- **Server**: `net/http` with `ServeMux`.

## Architecture
Client Request → http.Server → ServeMux Router ↓
CRUD Handlers → SQLite Storage → DB Operations ↓
JSON Response (200/404/500)

text
Config loads first → DB init → Router setup → ListenAndServe with signal handling.

## Installation
1. Ensure Go 1.21+ installed.
2. Copy files to project dir.
3. Init modules: `go mod init student-api && go mod tidy`.
4. Edit `local.yaml` for DB path/server addr.

## Environment Variables
Uses `local.yaml`:
env: local
addr: ":8080"
db_path: "./students.db"

text
No external env vars needed.

## Running
go run main.go

text
Server starts: `slog.Info("server started", slog.String("address", ":8080"))`

## Testing
Create student
curl -X POST http://localhost:8080/api/students
-H "Content-Type: application/json"
-d '{"name":"Alice","email":"alice@example.com"}'

List all
curl http://localhost:8080/api/students

Get by ID (replace :id)
curl http://localhost:8080/api/students/1

Update
curl -X PUT http://localhost:8080/api/students/1
-H "Content-Type: application/json"
-d '{"name":"Alice Updated"}'

Delete
curl -X DELETE http://localhost:8080/api/students/1

text
Graceful stop: Ctrl+C.

## Project Structure
student-api/
├── main.go # Entry: config → DB → router → server
├── conf.go # YAML config loader
├── local.yaml # Config file
├── student.go # Models + handlers
├── types.go # Shared types
├── response.go # JSON helpers
├── sqlite.go # DB init/queries
└── storage.go # Interface

text

## How It Works
1. `main()` loads config via `config.MustLoad()`
2. Inits SQLite: `sqlite.New(cfg)`
3. Sets up `http.NewServeMux()` with handlers like `student.New(storage)`
4. `server.ListenAndServe()` blocks; signals trigger `server.Shutdown(ctx)`
5. Handlers use storage interface for DB ops.

## Why These Technologies?
**Go Stdlib `net/http`**: Lightweight, fast, no framework bloat—perfect for microservices.  
**SQLite**: Zero-config DB, file-based, ideal for APIs without external services.  
**YAML Config**: Human-readable, flexible for dev/prod/env switching.  
**Slog**: Built-in structured logging beats fmt.Printf for observability.  
**Graceful Shutdown**: Production essential—prevents data corruption on restarts.
