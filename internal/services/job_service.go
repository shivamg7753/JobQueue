package services

import (
	"jobqueue/internal/models"
	"jobqueue/internal/repositories"
)

type JobService struct {
	Repo repositories.JobRepository
}

func (s *JobService) CreateJob(payload string) (int64, error) {
	return s.Repo.CreateJob(payload)
}

func (s *JobService) GetJobByID(id int64) (*models.Job, error) {
	return s.Repo.GetJobByID(id)
}

func (s *JobService) ListJobs(limit, offset int) ([]models.Job, error) {
	return s.Repo.ListJobs(limit, offset)
}

func (s *JobService) UpdateJobStatusAndResult(id int64, status, result string) error {
	return s.Repo.UpdateJobStatusAndResult(id, status, result)
}
