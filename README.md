# jobqueue


A high-performance asynchronous job queue system in Go, with RESTful APIs, worker pool, structured logging, and easy deployment to Render or Docker.

## Features

- Submit, retrieve, and list jobs via API
- Asynchronous job processing with worker pool
- Structured logging (logrus)
- PostgreSQL support
- Docker and Render deployment ready

## Project Structure

- `cmd/` - Main application entrypoint
- `internal/` - All application code (handlers, services, repositories, models, utils)
  - `internal/handlers/` - HTTP handlers
  - `internal/services/` - Business logic (job queue, workers)
  - `internal/repositories/` - Database repository layer
  - `internal/models/` - Database models
  - `internal/utils/` - Utilities (DB, logging, config)
  - `internal/constants.go` - Status and system constants
  - `internal/log_messages.go` - All log and error message constants (centralized)

## Getting Started (Local)

1. Clone the repo
2. Set environment variables (or create a `.env` file):
   - **Preferred for deployment:**
     ```
     DATABASE_URL=postgresql://<user>:<password>@<host>/<dbname>
     ```
   - **Or, for local development:**
     ```
     DB_HOST=localhost
     DB_PORT=5432
     DB_USER=Admin
     DB_PASSWORD=Admin123
     DB_NAME=JobQueue
     ```
3. Start PostgreSQL (Docker recommended):
   ```sh
   docker-compose up -d
   ```
4. Build and run:
   ```sh
   go build -o jobqueue ./cmd
   ./jobqueue
   ```
5. API available at `https://jobqueue-2.onrender.com`

## API Endpoints

- `POST https://jobqueue-2.onrender.com/jobs` — Submit a new job
- `GET https://jobqueue-2.onrender.com/jobs/{id}` — Get job status and result
- `GET https://jobqueue-2.onrender.com/jobs` — List all jobs with pagination



## Constants & Logging

- All error, log, and status message constants are centralized in:
  - `internal/log_messages.go` (log and error messages)
  - `internal/constants.go` (status and system constants)
- Update these files to change any user-facing or log messages project-wide.

## License

MIT
