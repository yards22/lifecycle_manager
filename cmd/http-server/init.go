package main

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	sqlc "github.com/yards22/lcmanager/db/sqlc"
	"github.com/yards22/lcmanager/internal/token_manager"
	"github.com/yards22/lcmanager/pkg/env"
)

func initDB() (*sql.DB, error) {
	db, err := sql.Open(env.ViperGetEnvVar("DB_DRIVER_NAME"), env.ViperGetEnvVar("DB_DATA_SOURCE_NAME"))
	if err != nil {
		return nil, err
	}
	return db, nil
}

func initManagers(app *App) {
	// Initialize managers
	tokenManager := token_manager.New(sqlc.New(app.db), time.Hour)
	// Add to app
	app.managers["tokenManager"] = tokenManager
}
