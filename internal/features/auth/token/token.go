package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

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
