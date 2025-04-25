package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	handlers "github.com/aprofessionale/crypto-alert-system/internal/api/handlers"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHealthCheckHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/healthz", nil)
	require.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.HealthCheckHandler)

	handler.ServeHTTP(rr, req)

	// Assert status code
	assert.Equal(t, http.StatusOK, rr.Code, "handler returned wrong status code")

	// Assert response body
	expectedBody := `{"status":"UP"}`
	assert.JSONEq(t, expectedBody, rr.Body.String(), "handler returned unexpected body")
}
