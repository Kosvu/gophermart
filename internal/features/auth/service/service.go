package auth_service

import (
	"context"

	"github.com/Kosvu/gophermart/internal/features/auth/token"
)

type AuthService struct {
	authRepository AuthRepository
	token          token.Token
}

type AuthRepository interface {
	Register(
		ctx context.Context,
		login string,
		password []byte,
	) error
}

func NewAuthService(authRepository AuthRepository, token token.Token) *AuthService {
	return &AuthService{
		authRepository: authRepository,
		token:          token,
	}
}
