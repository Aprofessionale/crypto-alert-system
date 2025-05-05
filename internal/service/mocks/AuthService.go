package mocks

import (
	"context"

	"github.com/aprofessionale/crypto-alert-system/internal/service"

	"github.com/stretchr/testify/mock"
)

type MockAuthService struct {
	mock.Mock
}

func (_m *MockAuthService) SubscribeUser(ctx context.Context, email string) error {
	ret := _m.Called(ctx, email)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, email)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

var _ service.AuthService = (*MockAuthService)(nil)
