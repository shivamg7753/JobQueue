package models

import (
	"time"
)

type Job struct {
	ID        int64     `db:"id" json:"id"`
	Payload   string    `db:"payload" json:"payload"`
	Status    string    `db:"status" json:"status"`
	Result    string    `db:"result" json:"result"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
