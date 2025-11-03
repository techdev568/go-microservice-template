package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/techdev568/go-microservice-template/internal/config"
	"github.com/techdev568/go-microservice-template/internal/logger"
	"github.com/techdev568/go-microservice-template/internal/server"
)

func main() {
	// Initialize logger
	log := logger.New()
	defer func() {
		_ = log.Sync()
	}()

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// Start server
	srv := server.New(cfg, log)

	go func() {
		if err := srv.Start(); err != nil {
			log.Fatalf("server error: %v", err)
		}
	}()

	// Graceful shutdown
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
