package balance_repository

import (
	"context"
	"fmt"
)

func (r *BalanceRepository) GetBalance(ctx context.Context, login string) (BalanceModel, error) {
	queryBalance := `
	SELECT balance
	FROM users
	WHERE login=$1
	`

	var balance float64

	if err := r.pool.QueryRowContext(ctx, queryBalance, login).Scan(&balance); err != nil {
		return BalanceModel{}, fmt.Errorf("select balance: %w", err)
	}

	queryWithdrawn := `
	SELECT COALESCE(SUM(amount), 0)
	FROM withdrawals
	WHERE user_login=$1
	`

	var withdrawn float64
	if err := r.pool.QueryRowContext(ctx, queryWithdrawn, login).Scan(&withdrawn); err != nil {
		return BalanceModel{}, fmt.Errorf("select withdrawn: %w", err)
	}

	return BalanceModel{Current: balance, Withdrawn: withdrawn}, nil
}
