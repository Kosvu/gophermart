package accrual_repository

import (
	"context"
	"fmt"
)

func (r *AccrualRepository) UpdateOrder(ctx context.Context, number, status string, accrual *float64) error {
	tx, err := r.pool.BeginTx(ctx, nil)

	if err != nil {
		return fmt.Errorf("begin tx: %w", err)
	}

	defer tx.Rollback()

	query := `
	UPDATE orders
	SET status = $1, accrual = $2
	WHERE number_order = $3
	RETURNING user_login 
	`

	var login string

	if err := tx.QueryRowContext(ctx, query, status, accrual, number).Scan(&login); err != nil {
		return fmt.Errorf("scan login: %w", err)
	}

	if accrual != nil && *accrual > 0 {
		query := `
		UPDATE users
		SET balance = balance + $1
		WHERE login = $2
		`

		_, err := tx.ExecContext(ctx, query, accrual, login)

		if err != nil {
			return fmt.Errorf("update user: %w", err)
		}
	}

	return tx.Commit()
}
