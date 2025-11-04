package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/techdev568/go-microservice-template/internal/config"
	"github.com/techdev568/go-microservice-template/internal/database"
	"github.com/techdev568/go-microservice-template/internal/logger"
	"github.com/techdev568/go-microservice-template/internal/server"
)

func main() {
	log := logger.New()
	defer func() { _ = log.Sync() }()

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	db, err := database.Connect(cfg, log)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	srv := server.New(cfg, log, db)

	go func() {
		if err := srv.Start(); err != nil {
			log.Fatalf("server error: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop
	log.Info("shutting down gracefully...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Stop(ctx); err != nil {
		log.Errorf("graceful shutdown failed: %v", err)
	}

	log.Info("server stopped")
}
