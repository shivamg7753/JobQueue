package services

import (
	"database/sql"
	"jobqueue/internal"
	"jobqueue/internal/utils"
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
			utils.Logger.WithField("worker", i).Info(internal.LogWorkerStarted)
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
		}).Info(internal.LogProcessingJob)
		// Simulate job processing (replace with real logic)
		result := internal.JobResultPrefix + job.Payload
		_, err := db.Exec(`UPDATE jobs SET status = $1, result = $2, updated_at = NOW() WHERE id = $3`, internal.StatusCompleted, result, job.ID)
		if err != nil {
			utils.Logger.WithFields(map[string]interface{}{
				"worker": workerID,
				"job_id": job.ID,
				"error":  err.Error(),
			}).Error(internal.LogFailedUpdateJob)
			continue
		}
		utils.Logger.WithFields(map[string]interface{}{
			"worker": workerID,
			"job_id": job.ID,
		}).Info(internal.LogCompletedJob)
	}
}
