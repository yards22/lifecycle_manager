package main

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
	sqlc "github.com/yards22/lcmanager/db/sqlc"
	"github.com/yards22/lcmanager/internal/token_manager"
	"github.com/yards22/lcmanager/pkg/env"
	"github.com/yards22/lcmanager/pkg/env/json"
)

func initDB() (*sql.DB, error) {
	db, err := sql.Open(env.ViperGetEnvVar("DB_DRIVER_NAME"), env.ViperGetEnvVar("DB_DATA_SOURCE_NAME"))
	if err != nil {
		return nil, err
	}
	return db, nil
}

func getDuration(i json.RunnerInterval) time.Duration {
	return i.Seconds*time.Second + i.Minutes*time.Minute + i.Hours*time.Hour
}

func initManagers(app *App) {
	managerConfigs, err := json.New("manspec.json")
	if err != nil {
		app.logger.Fatalln(err.Error())
	}
	app.managers["tokenManager"] = token_manager.New(sqlc.New(app.db), getDuration(managerConfigs.Get("tokenManager").Interval))
}

func initServer(app *App) {
	r := chi.NewRouter()
	handler(app, r)
	srv := http.Server{
		Addr:    env.ViperGetEnvVar("SERVER_ADDR"),
		Handler: r,
	}
	app.srv = &srv
}
