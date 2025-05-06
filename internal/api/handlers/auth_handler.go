package handlers

import (
	"net/http"

	"github.com/aprofessionale/crypto-alert-system/internal/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthHandler struct {
	authService service.AuthService
	logger      *zap.Logger
}

func NewAuthHandler(authSvc service.AuthService, logger *zap.Logger) *AuthHandler {
	return &AuthHandler{
		authService: authSvc,
		logger:      logger,
	}
}

func (h *AuthHandler) Subscribe(c *gin.Context) {
	var req SubscribeRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Warn("Failed to bind request for subscribe", zap.Error(err))
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
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
