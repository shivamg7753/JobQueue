package utils

import (
	"database/sql"
)

func MigrateDB(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS jobs (
		id SERIAL PRIMARY KEY,
		payload TEXT NOT NULL,
		status VARCHAR(32) NOT NULL,
		result TEXT,
		created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
	);
	`
	_, err := db.Exec(query)
	return err
}
