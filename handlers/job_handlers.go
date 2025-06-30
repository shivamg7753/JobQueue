package handlers

import (
	"encoding/json"
	"jobqueue/internal"
	"jobqueue/services"
	"net/http"
	"strconv"
	"strings"
)

var jobService *services.JobService

func SetJobService(s *services.JobService) {
	jobService = s
}

type submitJobRequest struct {
	Payload string `json:"payload"`
}

type submitJobResponse struct {
	ID int64 `json:"id"`
}

func writeJSONError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

func SubmitJobHandler(w http.ResponseWriter, r *http.Request) {
	var req submitJobRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSONError(w, http.StatusBadRequest, internal.ErrInvalidPayload)
		return
	}
	id, err := jobService.CreateJob(req.Payload)
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, internal.ErrFailedCreateJob+": "+err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(submitJobResponse{ID: id})
}

func GetJobHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		writeJSONError(w, http.StatusBadRequest, internal.ErrMissingJobID)
		return
	}
	id, err := strconv.ParseInt(parts[2], 10, 64)
	if err != nil {
		writeJSONError(w, http.StatusBadRequest, internal.ErrInvalidJobID)
		return
	}
	job, err := jobService.GetJobByID(id)
	if err != nil {
		writeJSONError(w, http.StatusNotFound, internal.ErrJobNotFound)
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
	jobs, err := jobService.ListJobs(limit, offset)
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, internal.ErrFailedListJobs+": "+err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jobs)
}
