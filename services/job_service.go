package services

import (
	"database/sql"
	"jobqueue/models"
	"time"
)

func CreateJob(db *sql.DB, payload string) (int64, error) {
	var id int64
	err := db.QueryRow(
		`INSERT INTO jobs (payload, status, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id`,
		payload, "pending", time.Now(), time.Now(),
	).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func GetJobByID(db *sql.DB, id int64) (*models.Job, error) {
	job := &models.Job{}
	row := db.QueryRow(`SELECT id, payload, status, result, created_at, updated_at FROM jobs WHERE id = $1`, id)
	if err := row.Scan(&job.ID, &job.Payload, &job.Status, &job.Result, &job.CreatedAt, &job.UpdatedAt); err != nil {
		return nil, err
	}
	return job, nil
}

func ListJobs(db *sql.DB, limit, offset int) ([]models.Job, error) {
	rows, err := db.Query(`SELECT id, payload, status, result, created_at, updated_at FROM jobs ORDER BY created_at DESC LIMIT $1 OFFSET $2`, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	jobs := []models.Job{}
	for rows.Next() {
		var job models.Job
		if err := rows.Scan(&job.ID, &job.Payload, &job.Status, &job.Result, &job.CreatedAt, &job.UpdatedAt); err != nil {
			return nil, err
		}
		jobs = append(jobs, job)
	}
	return jobs, nil
}
