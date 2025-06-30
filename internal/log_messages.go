package internal

const (
	LogEnvFileNotFound   = ".env file not found or failed to load: "
	LogStartingJobQueue  = "Starting Job Queue System..."
	LogFailedConnectDB   = "Failed to connect to DB: %v"
	LogDBMigrationFailed = "DB migration failed: %v"

	// Worker pool and service status/log messages
	LogWorkerStarted   = "Worker started"
	LogProcessingJob   = "Processing job"
	LogFailedUpdateJob = "Failed to update job status in DB"
	LogCompletedJob    = "Completed job"

	// Handler error messages
	ErrInvalidPayload  = "Invalid request payload"
	ErrFailedCreateJob = "Failed to create job"
	ErrMissingJobID    = "Missing job ID"
	ErrInvalidJobID    = "Invalid job ID"
	ErrJobNotFound     = "Job not found"
	ErrFailedListJobs  = "Failed to list jobs"
)
