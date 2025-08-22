package main

import (
	"context"
	"log"
	"personal_knowledge_tracker/internal/application"
	bot_v1 "personal_knowledge_tracker/internal/controller/http/v1/bot"

	"personal_knowledge_tracker/config"
)

func main() {
	ctx := context.Background()

	cfg, err := config.New()
	if err != nil {
		log.Fatal("failed to init config err:", err)
	}

	app, err := application.NewWithContext(ctx, cfg)
	if err != nil {
		log.Fatal("failed to init application err:", err)
	}

	app.InitLogger()

	usecases := app.InitUsecases()

	handler, err := bot_v1.NewHandler(&bot_v1.Dependency{
		Config:   cfg,
		Usecases: usecases,
		Logger:   app.GetLogger(),
	})
	if err != nil {
		log.Fatalf("failed to init handler: %w", err)
	}
	handler.Register()

	app.RegisterHandler(handler)

	app.Run()
}
