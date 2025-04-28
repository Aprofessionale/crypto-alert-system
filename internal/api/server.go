package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/aprofessionale/crypto-alert-system/internal/config"
	"go.uber.org/zap"
)

type Server struct {
	config     *config.Config
	logger     *zap.Logger
	router     http.Handler
	httpServer *http.Server
}

func NewServer(cfg *config.Config, logger *zap.Logger, router http.Handler) *Server {
	return &Server{
		config: cfg,
		logger: logger,
		router: router,
	}
}

func (s *Server) Start() error {
	addr := fmt.Sprintf(":%s", s.config.Port)
	// s.logger.Info("Attempting to start HTTP Server", zap.String("address", addr))

	s.httpServer = &http.Server{
		Addr:    addr,
		Handler: s.router,
		// TODO: Add recommended timeouts later for production readiness
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	s.logger.Info("HTTP Server starting to listen ", zap.String("address", addr))
	err := s.httpServer.ListenAndServe()

	if err != nil && err != http.ErrServerClosed {
		s.logger.Error("HTTP server ListenAndServe error", zap.Error(err))
		return fmt.Errorf("server failed to start or encountered error: %w", err)
	}

	s.logger.Info("HTTP server stopped listening")
	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	if s.httpServer == nil {
		s.logger.Warn("Shutdown called, but HTTP server was not running")
		return nil
	}

	s.logger.Info("Attempting graceful shutdown of HTTP server")
	err := s.httpServer.Shutdown(ctx)
	if err != nil {
		s.logger.Error("HTTP server graceful shutdown failed", zap.Error(err))
		return err
	}

	s.logger.Info("HTTP server shutdown complete.")
	return nil
}
