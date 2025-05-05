package service

import (
	"context"
	"errors"
)

type AuthService interface {
	// handle initial subscription request
	SubscribeUser(ctx context.Context, email string) error
}

var (
	ErrUserAlreadyExists  = errors.New("user with this email already exists")
	ErrNotificationFailed = errors.New("failed to send verification notification")
	ErrInvalidInput       = errors.New("invalid input provided")
)
