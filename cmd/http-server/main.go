package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/yards22/lcmanager/internal/feedback_manager"
	"github.com/yards22/lcmanager/internal/manager"
	"github.com/yards22/lcmanager/internal/poll_manager"
)

type App struct {
	db              *sql.DB
	logger          *log.Logger
	PollManager     *poll_manager.PollManager
	FeedbackManager *feedback_manager.FeedbackManager
	managers        map[string]manager.Manager
	srv             *http.Server
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
		logger:   l,
		managers: make(map[string]manager.Manager),
	}

	initRunnerManagers(app)
	for _, v := range app.managers {
		go v.Run()
	}

	initManagers(app)
	initServer(app)
	app.logger.Fatalln(app.srv.ListenAndServe())

}
