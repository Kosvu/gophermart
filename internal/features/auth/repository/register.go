package auth_repository

import (
	"context"
	"errors"
	"fmt"

	apperrors "github.com/Kosvu/gophermart/internal/core/errors"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

func (r *AuthRepository) Register(
	ctx context.Context,
	login string,
	password []byte,
) error {
	query := `
	INSERT INTO users (login,password)
	VALUES($1, $2)
	`

	_, err := r.pool.ExecContext(ctx, query, login, password)

	var pgErr *pgconn.PgError

	if err != nil {
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			return apperrors.ErrLoginTaken
		}

		return fmt.Errorf("auth repository: %w", err)
	}

	return nil

}
