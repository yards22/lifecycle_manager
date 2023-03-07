package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
	cors "github.com/rs/cors"
	"github.com/streadway/amqp"
	sqlc "github.com/yards22/lcmanager/db/sqlc"
	authservice "github.com/yards22/lcmanager/internal/auth_service"
	"github.com/yards22/lcmanager/internal/feedback_manager"
	"github.com/yards22/lcmanager/internal/poll_manager"
	"github.com/yards22/lcmanager/internal/r_manager"
	"github.com/yards22/lcmanager/internal/r_posts_manager"
	"github.com/yards22/lcmanager/internal/r_users_manager"
	scoremanager "github.com/yards22/lcmanager/internal/score_manager"
	"github.com/yards22/lcmanager/internal/t_posts_manager"
	"github.com/yards22/lcmanager/internal/t_users_manager"
	"github.com/yards22/lcmanager/internal/token_manager"
	"github.com/yards22/lcmanager/pkg/app_config"
	kvstore "github.com/yards22/lcmanager/pkg/kv_store"
	"github.com/yards22/lcmanager/pkg/mailer"
	objectstore "github.com/yards22/lcmanager/pkg/object_store"
)

type Author struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Summary struct {
	MatchId  string   `json:"match_id"`
	DataType string   `json:"data_type"`
	Score    []string `json:"Score"`
}

type Record struct {
	ID   string
	URLs []string
}

func initDB(app *App) {
	db, err := sql.Open(app_config.Data.MustString("DB_DRIVER_NAME"), app_config.Data.MustString("DB_DATA_SOURCE_NAME"))
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	app.db = db
	app.logger.Println("connected to db")

}

func initKVDB(app *App) {
	sess := session.Must(session.NewSession())
	db := dynamodb.New(sess, &aws.Config{
		Region:      aws.String(app_config.Data.MustString("Dynamo_Region")),
		Credentials: credentials.NewStaticCredentials("AKIAUZAIJPCMOYOR7ZEN", "HU9drLbe1E90lORcPlfDIsPlaxngAFuh+M3QbCqF", ""),
	})

	// client := asynq.NewClient(asynq.RedisClientOpt{
	// 	Addr:     "localhost:6379",
	// 	Password: "",
	// 	DB:       0,
	// })

	app.kvdb = db
	// tables, err := db.ListTables(&dynamodb.ListTablesInput{})

	// input := &dynamodb.GetItemInput{
	// 	Key: map[string]*dynamodb.AttributeValue{
	// 		"match_id": {
	// 			S: aws.String("match_2"),
	// 		},
	// 		"data_type": {
	// 			S: aws.String("match_2_raw"),
	// 		},
	// 	},
	// 	TableName: aws.String("IMatches"),
	// }

	// out, err := db.GetItem(input)

	// in, err := db.PutItem(&dynamodb.PutItemInput{
	// 	TableName: aws.String("IMatches"),
	// 	Item: map[string]*dynamodb.AttributeValue{
	// 		"match_id": {
	// 			S: aws.String("match_2"),
	// 		},
	// 		"data_type": {
	// 			S: aws.String("match_2_commentry"),
	// 		},
	// 		"Score": {
	// 			L: []*dynamodb.AttributeValue{
	// 				{
	// 					M: map[string]*dynamodb.AttributeValue{
	// 						"plate": {S: aws.String("test")},
	// 						"spoon": {
	// 							M: map[string]*dynamodb.AttributeValue{
	// 								"solid": {S: aws.String("rice")},
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// })
	// r := Summary{
	// 	MatchId:  "match_2",
	// 	DataType: "match_2_commentry",
	// 	Score: []string{
	// 		"https://example.com/first/link",
	// 		"https://example.com/second/url",
	// 	},
	// }
	// av, err := dynamodbattribute.MarshalMap(r)
	// if err != nil {
	// 	panic(fmt.Sprintf("failed to DynamoDB marshal Record, %v", err))
	// }

	// _, err = db.PutItem(&dynamodb.PutItemInput{
	// 	TableName: aws.String("IMatches"),
	// 	Item:      av,
	// })
	// if err != nil {
	// 	panic(fmt.Sprintf("failed to put Record to DynamoDB, %v", err))
	// }

	// fmt.Println("retrived data ", out)
	// fmt.Println("inserted_data ", in)

	// if err != nil {
	// 	panic(err)
	// }
	app.logger.Println("connected to kvDB")
}

func initRunnerManagers(app *App) {
	// Initialize managers and add to app

	querier := sqlc.New(app.db)

	// token manager runner
	d := time.Duration(app_config.Data.MustInt("duration_token") * int(time.Minute))
	tokenManager := token_manager.New(querier, d)
	app.managers["tokenManager"] = tokenManager

	// trending post runner
	d = time.Duration(app_config.Data.MustInt("duration_trending_post") * int(time.Minute))
	trendingPostsManager := t_posts_manager.New(querier, d)
	app.managers["trendingPostsManager"] = trendingPostsManager

	// trending user runner
	d = time.Duration(app_config.Data.MustInt("duration_trending_user") * int(time.Minute))
	trendingUserManager := t_users_manager.New(querier, d)
	app.managers["trendingUserManager"] = trendingUserManager

	// recommended user runner
	d = time.Duration(app_config.Data.MustInt("duration_recommended_user") * int(time.Minute))
	recommendedUsersManager := r_users_manager.New(querier, d)
	app.managers["recommendedUsersManager"] = recommendedUsersManager

	// recommended post runner
	d = time.Duration(app_config.Data.MustInt("duration_recommended_post") * int(time.Minute))
	recommendedPostsManager := r_posts_manager.New(querier, d)
	app.managers["recommendedPostsManager"] = recommendedPostsManager

	// rating runner
	d = time.Duration(app_config.Data.MustInt("duration_rating") * int(time.Minute))
	ratingManager := r_manager.New(querier, d)
	app.managers["ratingManager"] = ratingManager
}

func initConsumer(app *App) {
	conn, err := amqp.Dial("amqps://qwiynkfq:pEcA9NfiesS0wIbNrGewvVjIrqMmO4v4@puffin.rmq2.cloudamqp.com/qwiynkfq")
	if err != nil {
		fmt.Println("Failed Initializing Broker Connection")
		panic(err)
	}
	fmt.Println("Initializing Broker Connection")
	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
	}

	scoreManager := scoremanager.New(app.kvdb, ch)
	app.managers["score_manager"] = scoreManager

}

func initManagers(app *App) {

	// Initialize API Managers
	app.PollManager = poll_manager.New(sqlc.New(app.db))
	app.FeedbackManager = feedback_manager.New(sqlc.New(app.db))

}

func initServer(app *App) {
	r := chi.NewRouter()

	reactUri := app_config.Data.MustString("REACT_URI")

	r.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{reactUri},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "Authorization", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Authorization"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
	}).Handler)

	initHandler(app, r)
	srv := http.Server{
		Addr:    app_config.Data.MustString("SERVER_ADDR"),
		Handler: r,
	}
	app.srv = &srv
}

func initAuthService(app *App) {
	//  	mailer, err := mailer.NewGoMail("smtpout.secureserver.net", 587, "contact@22yardz.in", "JEvrW59syf5v9tc", true)
	app.authService = authservice.New(kvstore.New(), sqlc.New(app.db), app.mailer)
}

func initObjectStore(app *App) {
	accessId := app_config.Data.MustString("S3_ACCESS_ID")
	region := app_config.Data.MustString("S3_REGION")
	secret := app_config.Data.MustString("S3_SECRET")
	bucket := app_config.Data.MustString("S3_BUCKET")
	objectStore, err := objectstore.New(accessId, secret, region, bucket)
	if err != nil {
		panic(err)
	}
	app.objectStore = objectStore
}

func initMailer(app *App) {
	mailer, err := mailer.NewGoMail(app_config.Data.MustString("MAIL_HOST"), app_config.Data.MustInt("MAIL_PORT"), app_config.Data.MustString("MAIL_ID"), app_config.Data.MustString("MAIL_PASSWORD"), true)
	if err != nil {
		fmt.Println(err)
	}
	app.mailer = mailer
}
