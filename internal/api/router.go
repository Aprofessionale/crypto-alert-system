package api

import (
	"net/http"

	handlers "github.com/aprofessionale/crypto-alert-system/internal/api/handlers"
)

func NewStdLibRouter() *http.ServeMux {
	mux := http.NewServeMux()

	// Health check
	mux.HandleFunc("GET /healthz", handlers.HealthCheckHandler)

	return mux
}
