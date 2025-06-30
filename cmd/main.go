package main

import (
	"database/sql"
	"jobqueue/internal"
	"jobqueue/internal/handlers"
	"jobqueue/internal/repositories"
	"jobqueue/internal/services"
	"jobqueue/internal/utils"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

var dbInstance *sql.DB

func main() {
	if errEnv := godotenv.Load(); errEnv != nil {
		utils.InitLogger()
		utils.Logger.Warn(internal.LogEnvFileNotFound + errEnv.Error())
	}
	utils.InitLogger()
	utils.Logger.Info(internal.LogStartingJobQueue)

	var err error
	dbInstance, err = utils.GetDB()
	if err != nil {
		utils.Logger.Fatalf(internal.LogFailedConnectDB, err)
	}

	// Migrate DB (auto-create table)
	if err := utils.MigrateDB(dbInstance); err != nil {
		utils.Logger.Fatalf(internal.LogDBMigrationFailed, err)
	}

	repo := &repositories.PostgresJobRepository{DB: dbInstance}
	jobService := &services.JobService{Repo: repo}
	handlers.SetJobService(jobService)

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
