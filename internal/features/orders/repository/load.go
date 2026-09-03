package orders_repository

import (
	"context"
	"errors"
	"fmt"

	apperrors "github.com/Kosvu/gophermart/internal/core/errors"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

func (r *OrdersRepository) Load(ctx context.Context, login, number string) error {

	query := `
	INSERT INTO orders (number_order, status, user_login)
	VALUES ($1, 'NEW', $2)
	`

	if _, err := r.pool.ExecContext(ctx, query, number, login); err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			ownerQuery := `
			SELECT user_login FROM orders WHERE number_order = $1
			`
			var loginRow string

			row := r.pool.QueryRowContext(ctx, ownerQuery, number)

			if err := row.Scan(&loginRow); err != nil {
				return fmt.Errorf("row scan: %w", err)
			}

			if loginRow == login {
				return apperrors.ErrOrderAlreadyLoaded
			}

			return apperrors.ErrOrderTakenByAnother
		}

		return fmt.Errorf("exec query: %w", err)
	}

	return nil
}
