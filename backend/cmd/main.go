package main

import (
	"log"
	"os"

	"github.com/404th/Ink/api"
	"github.com/404th/Ink/api/handler"
	"github.com/404th/Ink/config"
	"github.com/404th/Ink/internal/postgres"
	"github.com/404th/Ink/internal/server"
	"github.com/404th/Ink/internal/service"
	"github.com/404th/Ink/internal/storage"
	"github.com/404th/Ink/internal/tigres"
	_ "github.com/lib/pq"

	"go.uber.org/zap"
)

func main() {
	// 0. Load configuration
	cfg := config.Load()

	// 1. Initialize logger
	var logger *zap.Logger
	switch os.Getenv(cfg.ProjectMode) {
	case config.ProjectModeDevelopment:
		logger = zap.Must(zap.NewDevelopment())
	case config.ProjectModeProduction:
		logger = zap.Must(zap.NewProduction())
	default:
		logger = zap.Must(zap.NewDevelopment())
	}
	defer logger.Sync()
	sugar := logger.Sugar()

	// 2. Initialize database
	db, err := postgres.NewPostgres(cfg)
	if err != nil {
		sugar.Fatalf("Failed to connect to databases: %v", err)
	}
	sugar.Infoln("Successfully connected to databases!")

	// 3. Initialize file database
	tg := tigres.NewTigres(cfg, sugar)

	// 4. Initialize storage
	storage := storage.NewStorage(db.Pool)

	// 5. Initialize service
	service := service.NewService(cfg, sugar, storage, tg)

	// 6. Initialize handler
	handler := handler.NewHandler(cfg, sugar, service, tg)

	// 7. running api router
	r := api.Run(cfg, sugar, handler)

	// printing api's in log
	for _, v := range r.Routes() {
		log.Println(v.Method, v.Path)
	}

	// 8. Start the server
	server.Run(cfg, sugar, r)
}
