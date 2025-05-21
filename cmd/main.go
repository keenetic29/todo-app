package main

import (
	"log"
	"todo-app/internal/api"
	"todo-app/internal/config"
	"todo-app/pkg/database"
)

func main() {
	cfg, err := config.LoadConfig("conf.env")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	db, err := database.InitDB(cfg.DBURL)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	router := api.SetupRouter(db, cfg)

	log.Printf("Starting server on %s", cfg.ServerAddress)
	if err := router.Run(cfg.ServerAddress); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}