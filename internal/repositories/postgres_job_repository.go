package repositories

import (
	"database/sql"
	"fmt"
	"jobqueue/internal"
	"jobqueue/internal/models"
	"time"
)

type PostgresJobRepository struct {
	DB *sql.DB
}

func (r *PostgresJobRepository) CreateJob(payload string) (int64, error) {
	var id int64
	err := r.DB.QueryRow(
		`INSERT INTO jobs (payload, status, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id`,
		payload, internal.StatusPending, time.Now(), time.Now(),
	).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("insert job: %w", err)
	}
	return id, nil
}

func (r *PostgresJobRepository) GetJobByID(id int64) (*models.Job, error) {
	job := &models.Job{}
	row := r.DB.QueryRow(`SELECT id, payload, status, result, created_at, updated_at FROM jobs WHERE id = $1`, id)
	if err := row.Scan(&job.ID, &job.Payload, &job.Status, &job.Result, &job.CreatedAt, &job.UpdatedAt); err != nil {
		return nil, fmt.Errorf("get job by id: %w", err)
	}
	return job, nil
}

func (r *PostgresJobRepository) ListJobs(limit, offset int) ([]models.Job, error) {
	rows, err := r.DB.Query(`SELECT id, payload, status, result, created_at, updated_at FROM jobs ORDER BY created_at DESC LIMIT $1 OFFSET $2`, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("list jobs: %w", err)
	}
	defer rows.Close()
	jobs := []models.Job{}
	for rows.Next() {
		var job models.Job
		if err := rows.Scan(&job.ID, &job.Payload, &job.Status, &job.Result, &job.CreatedAt, &job.UpdatedAt); err != nil {
			return nil, fmt.Errorf("scan job row: %w", err)
		}
		jobs = append(jobs, job)
	}
	return jobs, nil
}

func (r *PostgresJobRepository) UpdateJobStatusAndResult(id int64, status, result string) error {
	_, err := r.DB.Exec(`UPDATE jobs SET status = $1, result = $2, updated_at = NOW() WHERE id = $3`, status, result, id)
	if err != nil {
		return fmt.Errorf("update job: %w", err)
	}
	return nil
}
