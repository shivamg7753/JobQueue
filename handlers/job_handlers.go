package handlers

import (
	"database/sql"
	"encoding/json"
	"jobqueue/services"
	"jobqueue/utils"
	"net/http"
	"strconv"
	"strings"
)

var dbInstance *sql.DB

func SetDBInstance(db *sql.DB) {
	dbInstance = db
}

type submitJobRequest struct {
	Payload string `json:"payload"`
}

type submitJobResponse struct {
	ID int64 `json:"id"`
}

func SubmitJobHandler(w http.ResponseWriter, r *http.Request) {
	var req submitJobRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid request payload"))
		return
	}
	id, err := services.CreateJob(dbInstance, req.Payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to create job"))
		return
	}
	services.EnqueueJob(id, req.Payload)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(submitJobResponse{ID: id})
}

func GetJobHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Missing job ID"))
		return
	}
	id, err := strconv.ParseInt(parts[2], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid job ID"))
		return
	}
	db, err := utils.GetDB()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Database connection error"))
		return
	}
	defer db.Close()
	job, err := services.GetJobByID(db, id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Job not found"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(job)
}

func ListJobsHandler(w http.ResponseWriter, r *http.Request) {
	limit := 10
	offset := 0
	q := r.URL.Query()
	if l := q.Get("limit"); l != "" {
		if v, err := strconv.Atoi(l); err == nil {
			limit = v
		}
	}
	if o := q.Get("offset"); o != "" {
		if v, err := strconv.Atoi(o); err == nil {
			offset = v
		}
	}
	db, err := utils.GetDB()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Database connection error"))
		return
	}
	defer db.Close()
	jobs, err := services.ListJobs(db, limit, offset)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to list jobs"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jobs)
}
