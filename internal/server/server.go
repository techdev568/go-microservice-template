package server

import (
	"context"
	"fmt"

	"github.com/techdev568/go-microservice-template/internal/api"
	"github.com/techdev568/go-microservice-template/internal/config"

	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Server struct {
	cfg    *config.Config
	log    *zap.SugaredLogger
	engine *gin.Engine
	http   *http.Server
}

func New(cfg *config.Config, log *zap.SugaredLogger) *Server {
	router := gin.New()
	router.Use(gin.Recovery())

	// Register routes
	api.RegisterHealthRoutes(router)

	s := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Port),
		Handler: router,
	}

	return &Server{
		cfg:    cfg,
		log:    log,
		engine: router,
		http:   s,
	}
}

func (s *Server) Start() error {
	s.log.Infof("starting %s on port %s", s.cfg.AppName, s.cfg.Port)
	return s.http.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	s.log.Info("stopping server...")
	return s.http.Shutdown(ctx)
}
