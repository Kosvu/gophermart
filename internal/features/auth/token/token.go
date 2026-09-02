package token

import (
	"fmt"
	"time"

	apperrors "github.com/Kosvu/gophermart/internal/core/errors"
	"github.com/golang-jwt/jwt/v5"
)

const DefaultTTL = 24 * time.Hour

type Token struct {
	secret string
	ttl    time.Duration
}

func NewToken(secret string, ttl time.Duration) *Token {
	return &Token{
		secret: secret,
		ttl:    ttl,
	}
}

type Claims struct {
	Login string
	jwt.RegisteredClaims
}

func (t *Token) Validate(tokenString string) (string, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		func(tn *jwt.Token) (any, error) { return []byte(t.secret), nil },
		jwt.WithValidMethods([]string{"HS256"}),
	)

	if err != nil {
		return "", apperrors.ErrInvalidToken
	}

	if !token.Valid {
		return "", apperrors.ErrInvalidToken
	}

	return claims.Login, nil
}

func (t *Token) Issue(login string) (string, error) {
	claims := Claims{
		Login: login,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(t.ttl)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(t.secret))
	if err != nil {
		return "", fmt.Errorf("signed token: %w", err)
	}

	return signed, nil
}
