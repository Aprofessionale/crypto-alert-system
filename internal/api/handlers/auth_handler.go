package handlers

import (
	"net/http"

	"github.com/aprofessionale/crypto-alert-system/internal/service"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authSvc service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authSvc,
	}
}

func (h *AuthHandler) Subscribe(c *gin.Context) {
	var req SubscribeRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid request body: " + err.Error()})
		return
	}

	err := h.authService.SubscribeUser(c.Request.Context(), req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Internal server error: "})
		return
	}

	c.JSON(http.StatusOK, GeneralResponse{
		Message: "Subscription initiated. Check email for verification code.",
	})
}
