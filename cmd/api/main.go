package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/aprofessionale/crypto-alert-system/internal/api"
	"github.com/aprofessionale/crypto-alert-system/internal/config"
	appLogger "github.com/aprofessionale/crypto-alert-system/internal/logger"
	"go.uber.org/zap"
)

func main() {
	logger := appLogger.NewLogger()

	defer func() {
		if err := logger.Sync(); err != nil {
			log.Printf("Failed to sync Zap logger: %v", err)
		}
	}()

	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Fatal("Failed to load configuration", zap.Error(err))
	}
	logger.Info("Configuration loaded successfully", zap.String("port", cfg.Port))

	router := api.NewStdLibRouter()

	serverWrapper := api.NewServer(cfg, logger, router)

	serverErrors := make(chan error, 1)
	var httpServer *http.Server

	go func() {
		logger.Info("Starting server goroutine . . .")
		var startErr error

		httpServer, startErr = serverWrapper.Start()
		serverErrors <- startErr
	}()

	shutdownChannel := make(chan os.Signal, 1)
	signal.Notify(shutdownChannel, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-serverErrors:
		// Server failed to start or stopped unexpectedly
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Fatal("Server error", zap.Error(err))
		} else {
			// This case handles if the server stops normally before a signal is received
			// (e.g., port already in use error during ListenAndServe)
			logger.Info("Server stopped or failed to start", zap.Error(err))
		}
	case sig := <-shutdownChannel:
		logger.Info("Shutdown signal received", zap.String("signal", sig.String()))

		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if httpServer != nil {
			logger.Info("attempting graceful shutdown of HTTP server . . .")
			if err := httpServer.Shutdown(shutdownCtx); err != nil {
				logger.Error("HTTP server graceful shutdown failed", zap.Error(err))
			} else {
				logger.Info("HTTP server shutdown complete.")
			}
		} else {
			logger.Warn("HTTP server instance was nil, could not perform graceful shutdown.")
		}

		<-shutdownCtx.Done()
		if errors.Is(shutdownCtx.Err(), context.DeadlineExceeded) {
			logger.Warn("Graceful shutdown deadline exceeded.")
		}
	}

	logger.Info("application exiting")
}
