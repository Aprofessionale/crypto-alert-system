package api

import (
	handlers "github.com/aprofessionale/crypto-alert-system/internal/api/handlers"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	// Health check
	router.GET("/healthz", handlers.HealthCheckHandler)

	return router
}
