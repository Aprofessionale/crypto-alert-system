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
		// Attempt to sync the logger, log error if it fails
		if err := logger.Sync(); err != nil && errors.Is(err, syscall.EBADF) && !errors.Is(err, os.ErrInvalid) {
			// Ignore common errors during shutdown like "invalid argument" or "bad file descriptor"
			log.Printf("Failed to sync Zap logger: %v", err)
		}
	}()

	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Fatal("Failed to load configuration", zap.Error(err))
	}
	logger.Info("Configuration loaded successfully")

	router := api.NewStdLibRouter()

	serverWrapper := api.NewServer(cfg, logger, router)

	serverErrors := make(chan error, 1)

	go func() {
		logger.Info("Starting server goroutine . . .")

		serverErrors <- serverWrapper.Start()
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

		if err := serverWrapper.Shutdown(shutdownCtx); err != nil {
			// Log error, but continue shutdown process
			logger.Error("HTTP server shutdown error", zap.Error(err))
		}
	}

	logger.Info("application exiting")
}
