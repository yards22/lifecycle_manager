package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	authservice "github.com/yards22/lcmanager/internal/auth_service"
	"github.com/yards22/lcmanager/internal/feedback_manager"
	"github.com/yards22/lcmanager/internal/manager"
	"github.com/yards22/lcmanager/internal/poll_manager"
	"github.com/yards22/lcmanager/pkg/mailer"
	objectstore "github.com/yards22/lcmanager/pkg/object_store"
)

type App struct {
	db   *sql.DB
	kvdb *dynamodb.DynamoDB

	// redis           *redis.Client
	logger          *log.Logger
	PollManager     *poll_manager.PollManager
	FeedbackManager *feedback_manager.FeedbackManager
	managers        map[string]manager.Manager
	srv             *http.Server
	authService     *authservice.AuthService
	objectStore     objectstore.ObjectStore
	mailer          *mailer.GoMail
}

var (
	l = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
)

func main() {

	// redis := initRedis()
	// if err != nil {
	// 	l.Fatal("Error: Cannot initialize cache", err)
	// 	return
	// }

	app := &App{
		logger:   l,
		managers: make(map[string]manager.Manager),
	}

	initDB(app)
	initKVDB(app)
	initRunnerManagers(app)
	for _, v := range app.managers {
		go v.Run()
	}
	initManagers(app)
	initServer(app)
	initMailer(app)
	initAuthService(app)
	initObjectStore(app)
	app.logger.Println("app running on", app.srv.Addr)
	app.logger.Fatalln(app.srv.ListenAndServe())

}
