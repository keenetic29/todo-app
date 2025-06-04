package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"todo-app/internal/api"
	"todo-app/internal/config"
	"todo-app/pkg/database"
	"todo-app/pkg/logger"
)

func main() {
	cfg, err := config.LoadConfig("conf.env")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	logger.Init(cfg.LogFile)
	logger.Info.Println("Starting application...")
	logger.Info.Printf("Logging to file: %s", cfg.LogFile)

	db, err := database.InitDB(cfg.DBURL)
	if err != nil {
		logger.Error.Fatalf("Failed to initialize database: %v", err)
	}

	router := api.SetupRouter(db, cfg)

	srv := &http.Server{
		Addr:    cfg.ServerAddress,
		Handler: router,
	}

	// Запуск сервера в горутине
	go func() {
		logger.Info.Printf("Starting server on %s", cfg.ServerAddress)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Канал для перехвата сигналов
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info.Println("Shutting down server...")

	// Создаем контекст с таймаутом для graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Error.Printf("Server forced to shutdown: %v", err)
	}

	logger.Info.Println("Server exiting")
}