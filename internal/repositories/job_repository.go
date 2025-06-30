package repositories

import "jobqueue/internal/models"

type JobRepository interface {
	CreateJob(payload string) (int64, error)
	GetJobByID(id int64) (*models.Job, error)
	ListJobs(limit, offset int) ([]models.Job, error)
	UpdateJobStatusAndResult(id int64, status, result string) error
}
