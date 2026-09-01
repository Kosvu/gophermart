package auth_repository

import (
	"database/sql"
)

type AuthRepository struct {
	pool *sql.DB
}

func NewAuth(pool *sql.DB) *AuthRepository {
	return &AuthRepository{
		pool: pool,
	}
}
