package main

import (
	"context"
	"log"
	"personal_knowledge_tracker/internal/application"
	"personal_knowledge_tracker/internal/controller/http"
	httpserver "personal_knowledge_tracker/pkg/server/http"
	"time"

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

	httpRouter := http.NewRouter(http.Dependencies{
		Cfg:      cfg,
		Logger:   app.GetLogger(),
		Usecases: usecases,
	})

	httpServer := httpserver.New(httpRouter,
		httpserver.Port(cfg.HTTP.Server.Port),
		httpserver.WriteTimeout(time.Duration(cfg.HTTP.Server.WriteTimeout)*time.Second),
	)

	app.RegisterHTTPServer(httpServer)

	app.Run()
}
