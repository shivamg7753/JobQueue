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
- `handlers/` - HTTP handlers
- `services/` - Business logic (job queue, workers)
- `models/` - Database models
- `utils/` - Utilities (DB, logging, config)

## Getting Started (Local)

1. Clone the repo
2. Set environment variables (or create a `.env` file):
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
5. API available at `http://localhost:8081`

## API Endpoints

- `POST /jobs` — Submit a new job
- `GET /jobs/{id}` — Get job status and result
- `GET /jobs` — List all jobs with pagination


## License

MIT
