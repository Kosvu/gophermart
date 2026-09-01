package auth_service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	apperrors "github.com/Kosvu/gophermart/internal/core/errors"
	"golang.org/x/crypto/bcrypt"
)

func (s *AuthService) Login(
	ctx context.Context,
	login string,
	password string,
) (string, error) {
	if login == "" || password == "" {
		return "", fmt.Errorf("account data: %w", apperrors.ErrInvalidInput)
	}

	hashPass, err := s.authRepository.Login(ctx, login)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", fmt.Errorf("user not found: %w", apperrors.ErrInvalidCredentials)
		}
		return "", fmt.Errorf("auth repository: %w", err)
	}

	if err := bcrypt.CompareHashAndPassword(hashPass, []byte(password)); err != nil {
		return "", fmt.Errorf("compare hash and password: %w", apperrors.ErrInvalidCredentials)
	}

	token, err := s.token.Issue(login)

	if err != nil {
		return "", fmt.Errorf("issue token: %w", err)
	}

	return token, err
}
