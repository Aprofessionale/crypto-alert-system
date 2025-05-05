package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Returns the status of the service
func HealthCheckHandler(c *gin.Context) {
	response := gin.H{"status": "up"}

	c.JSON(http.StatusOK, response)
}
