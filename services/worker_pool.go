package services

import (
	"database/sql"
	"jobqueue/utils"
	"sync"
)

type JobPayload struct {
	ID      int64
	Payload string
}

var (
	jobQueue chan JobPayload
	once     sync.Once
)

func StartWorkerPool(db *sql.DB, numWorkers int) {
	once.Do(func() {
		jobQueue = make(chan JobPayload, 100)
		for i := 0; i < numWorkers; i++ {
			utils.Logger.WithField("worker", i).Info("Worker started")
			go worker(db, i)
		}
	})
}

func EnqueueJob(id int64, payload string) {
	jobQueue <- JobPayload{ID: id, Payload: payload}
}

func worker(db *sql.DB, workerID int) {
	for job := range jobQueue {
		utils.Logger.WithFields(map[string]interface{}{
			"worker": workerID,
			"job_id": job.ID,
		}).Info("Processing job")
		// Simulate job processing (replace with real logic)
		result := "processed: " + job.Payload
		_, err := db.Exec(`UPDATE jobs SET status = $1, result = $2, updated_at = NOW() WHERE id = $3`, "completed", result, job.ID)
		if err != nil {
			utils.Logger.WithFields(map[string]interface{}{
				"worker": workerID,
				"job_id": job.ID,
				"error":  err.Error(),
			}).Error("Failed to update job")
			continue
		}
		utils.Logger.WithFields(map[string]interface{}{
			"worker": workerID,
			"job_id": job.ID,
		}).Info("Completed job")
	}
}
