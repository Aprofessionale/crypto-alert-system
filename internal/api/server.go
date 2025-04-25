package api

import (
	"fmt"
	"net/http"

	"github.com/aprofessionale/crypto-alert-system/internal/config"
	"go.uber.org/zap"
)

type Server struct {
	config *config.Config
	logger *zap.Logger
	router http.Handler
}

func NewServer(cfg *config.Config, logger *zap.Logger, router http.Handler) *Server {
	return &Server{
		config: cfg,
		logger: logger,
		router: router,
	}
}

func (s *Server) Start() (*http.Server, error) {
	addr := fmt.Sprintf(":%s", s.config.Port)
	s.logger.Info("Attempting to start HTTP Server", zap.String("address", addr))

	httpServer := &http.Server{
		Addr:    addr,
		Handler: s.router,
		// TODO: Add recommended timeouts later for production readiness
		// ReadTimeout:  5 * time.Second,
		// WriteTimeout: 10 * time.Second,
		// IdleTimeout:  120 * time.Second,
	}

	s.logger.Info("HTTP Server starting to listen ", zap.String("address", addr))
	err := httpServer.ListenAndServe()

	if err != nil && err != http.ErrServerClosed {
		s.logger.Error("HTTP server ListenAndServe error", zap.Error(err))
		return httpServer, fmt.Errorf("server failed to start or encountered error: %w", err)
	}

	s.logger.Info("HTTP server stopped listening")
	return httpServer, err
}
