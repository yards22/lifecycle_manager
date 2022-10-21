package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	sqlc "github.com/yards22/lcmanager/db/sqlc"
	"github.com/yards22/lcmanager/internal/manager"
)

type App struct {
	db       *sql.DB
	store    *sqlc.Queries
	logger   *log.Logger
	managers map[string]manager.Manager
	srv      *http.Server
}

var (
	l = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
)

func main() {
	db, err := initDB()
	if err != nil {
		l.Fatal("Error: Cannot initialize database", err)
		return
	}

	app := &App{
		db:       db,
		store:    sqlc.New(db),
		logger:   l,
		managers: make(map[string]manager.Manager),
	}

	initManagers(app)
	for _, v := range app.managers {
		go v.Run()
	}

	initServer(app)
	app.logger.Println("starting server...")
	app.logger.Fatalln(app.srv.ListenAndServe())
}
