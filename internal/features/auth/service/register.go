package auth_service

import (
	"context"
	"fmt"

	apperrors "github.com/Kosvu/gophermart/internal/core/errors"
	"golang.org/x/crypto/bcrypt"
)

func (s *AuthService) Register(ctx context.Context, login string, password string) (string, error) {
	if login == "" || password == "" {
		return "", fmt.Errorf("account data: %w", apperrors.ErrInvalidInput)
	}

	hashPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", fmt.Errorf("hash password: %w", err)
	}

	if err = s.authRepository.Register(ctx, login, hashPass); err != nil {
		return "", fmt.Errorf("auth repository: %w", err)
	}

	token, err := s.token.Issue(login)

	if err != nil {
		return "", fmt.Errorf("token issue: %w", err)
	}

	return token, nil
}
