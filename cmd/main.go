package main

import (
	"log"
	"os"

	"github.com/mrtuuro/auto-messager/internal/application"
	"github.com/mrtuuro/auto-messager/internal/autosend"
	"github.com/mrtuuro/auto-messager/internal/config"
	"github.com/mrtuuro/auto-messager/internal/db"
	"github.com/mrtuuro/auto-messager/internal/dispatcher"
	"github.com/mrtuuro/auto-messager/internal/repository"
	"github.com/mrtuuro/auto-messager/internal/router"
	"github.com/mrtuuro/auto-messager/internal/service"
)

func main() {
	cfg := config.NewConfig()

	mongoClient, err := db.Connect(cfg.MongoURI)
	if err != nil {
		log.Fatalf("err connecting db: %v", err)
		os.Exit(1)
	}

	messageColl := db.GetCollection(mongoClient, cfg.DatabaseName, cfg.CollectionName)
	send := dispatcher.New(cfg.WebhookURL, cfg.WebhookKey)
	repo := repository.NewMongoMessageRepository(messageColl)

	messageSvc := service.NewMessageService(repo, send)
	sched := autosend.NewScheduler(messageSvc)
	sched.Start()

	app := application.NewApp(cfg, messageSvc, sched)

	router.Register(app)
	router.PrintRoutes(app)

	app.Run(cfg.Port)
}
