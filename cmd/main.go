package main

import (
	"database/sql"
	"jobqueue/handlers"
	"jobqueue/services"
	"jobqueue/utils"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

var dbInstance *sql.DB

func main() {
	_ = godotenv.Load() // Load .env file if present
	utils.InitLogger()
	utils.Logger.Info("Starting Job Queue System...")

	var err error
	dbInstance, err = utils.GetDB()
	if err != nil {
		utils.Logger.Fatalf("Failed to connect to DB: %v", err)
	}

	// Migrate DB (auto-create table)
	if err := utils.MigrateDB(dbInstance); err != nil {
		utils.Logger.Fatalf("DB migration failed: %v", err)
	}

	handlers.SetDBInstance(dbInstance)

	services.StartWorkerPool(dbInstance, 5)

	http.HandleFunc("/jobs", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handlers.SubmitJobHandler(w, r)
		case http.MethodGet:
			handlers.ListJobsHandler(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/jobs/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handlers.GetJobHandler(w, r)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
