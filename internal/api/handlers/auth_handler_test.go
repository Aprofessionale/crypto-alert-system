package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aprofessionale/crypto-alert-system/internal/api/handlers"
	"github.com/aprofessionale/crypto-alert-system/internal/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// TestSubscribe_Success tests the successful subscription scenario.
func TestSubscribe_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockAuthService := new(mocks.MockAuthService)

	authHandler := handlers.NewAuthHandler(mockAuthService) // needs to be created

	email := "test@example.com"
	requestBody := map[string]interface{}{"email": email}
	bodyBytes, err := json.Marshal(requestBody)
	require.NoError(t, err) // for setup steps
	bodyReader := bytes.NewBuffer(bodyBytes)

	mockAuthService.On("SubscribeUser", mock.Anything, email).Return(nil).Once()

	router := gin.New()
	router.POST("/auth/subscribe", authHandler.Subscribe)

	req, err := http.NewRequest(http.MethodPost, "/auth/subscribe", bodyReader)
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	expectedBody := `{"message":"Subscription initiated. Check email for verification code."}`
	assert.JSONEq(t, expectedBody, rr.Body.String())

	mockAuthService.AssertExpectations(t)
}
