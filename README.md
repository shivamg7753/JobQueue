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
5. API available at `http://localhost:8081`

## API Endpoints

- `POST /jobs` — Submit a new job
- `GET /jobs/{id}` — Get job status and result
- `GET /jobs` — List all jobs with pagination

## Deployment on Render

1. Push your code to GitHub (repo name: `jobqueue` recommended)
2. Go to [Render Dashboard](https://dashboard.render.com/)
3. Click **New > Web Service** and connect your repo
4. Set:
   - **Build Command:** `go build -tags netgo -ldflags '-s -w' -o app`
   - **Start Command:** `./app`
5. Add environment variable `DATABASE_URL` (from your Render Postgres dashboard)
6. (Optional) Add a managed Postgres instance from Render and use its credentials
7. Click **Deploy Web Service**

## License

MIT
