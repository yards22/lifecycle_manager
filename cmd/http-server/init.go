package main

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
	sqlc "github.com/yards22/lcmanager/db/sqlc"
	authservice "github.com/yards22/lcmanager/internal/auth_service"
	"github.com/yards22/lcmanager/internal/feedback_manager"
	"github.com/yards22/lcmanager/internal/poll_manager"
	"github.com/yards22/lcmanager/internal/r_manager"
	"github.com/yards22/lcmanager/internal/r_posts_manager"
	"github.com/yards22/lcmanager/internal/r_users_manager"
	"github.com/yards22/lcmanager/internal/t_posts_manager"
	"github.com/yards22/lcmanager/internal/t_users_manager"
	"github.com/yards22/lcmanager/internal/token_manager"
	"github.com/yards22/lcmanager/pkg/env"
	kvstore "github.com/yards22/lcmanager/pkg/kv_store"
	objectstore "github.com/yards22/lcmanager/pkg/object_store"
	runner "github.com/yards22/lcmanager/pkg/runner"
)

type Author struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func initDB() (*sql.DB, error) {
	db, err := sql.Open(env.ViperGetEnvVar("DB_DRIVER_NAME"), env.ViperGetEnvVar("DB_DATA_SOURCE_NAME"))
	if err != nil {
		return nil, err
	}
	return db, nil
}

func initRunnerManagers(app *App) {
	// Initialize managers and add to app
	tokenManager := token_manager.New(sqlc.New(app.db), (runner.TCleanerFrequency)*(time.Hour))
	app.managers["tokenManager"] = tokenManager
	trendingPostsManager := t_posts_manager.New(sqlc.New(app.db), (runner.TPostsFrequency)*(time.Hour))
	app.managers["trendingPostsManager"] = trendingPostsManager
	trendingUserManager := t_users_manager.New(sqlc.New(app.db), (runner.TUsersFrequency)*(time.Hour))
	app.managers["trendingUserManager"] = trendingUserManager
	recommendedUsersManager := r_users_manager.New(sqlc.New(app.db), (runner.RUsersFrequency)*(time.Hour))
	app.managers["recommendedUsersManager"] = recommendedUsersManager
	recommendedPostsManager := r_posts_manager.New(sqlc.New(app.db), (runner.RPostsFrequency)*(time.Hour))
	app.managers["recommendedPostsManager"] = recommendedPostsManager
	ratingManager := r_manager.New(sqlc.New(app.db), (runner.RatingFrequency)*(time.Hour))
	app.managers["ratingManager"] = ratingManager
}

func initManagers(app *App) {

	// Initialize API Managers
	app.PollManager = poll_manager.New(sqlc.New(app.db))
	app.FeedbackManager = feedback_manager.New(sqlc.New(app.db))

}

func initServer(app *App) {
	r := chi.NewRouter()
	initHandler(app, r)
	srv := http.Server{
		Addr:    env.ViperGetEnvVar("SERVER_ADDR"),
		Handler: r,
	}
	app.srv = &srv
}

func initAuthService(app *App) {
	app.authService = authservice.New(kvstore.New(), sqlc.New(app.db))
}

func initObjectStore(app *App) {
	accessId := env.ViperGetEnvVar("S3_ACCESS_ID")
	region := env.ViperGetEnvVar("S3_REGION")
	secret := env.ViperGetEnvVar("S3_SECRET")
	bucket := env.ViperGetEnvVar("S3_BUCKET")
	objectStore, err := objectstore.New(accessId, secret, region, bucket)
	if err != nil {
		panic(err)
	}
	app.objectStore = objectStore
}
