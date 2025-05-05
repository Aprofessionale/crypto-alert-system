package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	handlers "github.com/aprofessionale/crypto-alert-system/internal/api/handlers"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHealthCheckHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.New()
	router.GET("/healthz", handlers.HealthCheckHandler)

	req, err := http.NewRequest(http.MethodGet, "/healthz", nil)
	require.NoError(t, err, "Failed to create request")

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code, "Handler returned wrong status")

	expectedBody := `{"status":"up"}`
	assert.JSONEq(t, expectedBody, rr.Body.String(), "Handler returned wrong response body")

	assert.Equal(t, "application/json; charset=utf-8", rr.Header().Get("Content-Type"), "Handler returned unexpected content type header")
}
