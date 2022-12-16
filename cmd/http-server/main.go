package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/yards22/lcmanager/internal/manager"
)

type App struct {
	db       *sql.DB
	logger   *log.Logger
	managers map[string]manager.Manager
	srv      *http.Server
}

var (
	l = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
)

func main() {
	fmt.Println("main.go, 1")
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

	initManagers(app)
	for _, v := range app.managers {
		go v.Run()
	}

	initServer(app)
	app.logger.Fatalln(app.srv.ListenAndServe())
}
