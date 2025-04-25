package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

// Returns the status of the service
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// Ensure we only handle GET requests
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]string{"status": "UP"}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding health check response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
