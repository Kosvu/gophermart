package auth_repository

import (
	"context"
	"fmt"
)

func (r *AuthRepository) Login(
	ctx context.Context,
	login string,
) ([]byte, error) {
	query := `
	SELECT password
	FROM users
	WHERE login = $1
	`
	row := r.pool.QueryRowContext(ctx, query, login)

	var hashPass []byte

	if err := row.Scan(&hashPass); err != nil {
		return nil, fmt.Errorf("row scan: %w", err)
	}

	return hashPass, nil

}
