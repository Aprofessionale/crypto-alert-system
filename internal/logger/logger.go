package logger

import (
	"log"

	"go.uber.org/zap"
)

func NewLogger() *zap.Logger {
	// Development stage logger
	logger, err := zap.NewDevelopment()

	if err != nil {
		log.Fatalf("Cannot initialize Zap logger: %v", err)
	}

	logger.Info("Logger initialized successfully")
	return logger
}
