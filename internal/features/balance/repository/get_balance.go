package balance_repository

import (
	"context"
	"fmt"

	core_domain "github.com/Kosvu/gophermart/internal/core/domain"
)

func (r *BalanceRepository) GetBalance(ctx context.Context, login string) (core_domain.Balance, error) {
	queryBalance := `
	SELECT balance
	FROM users
	WHERE login=$1
	`

	var balance float64

	if err := r.pool.QueryRowContext(ctx, queryBalance, login).Scan(&balance); err != nil {
		return core_domain.Balance{}, fmt.Errorf("select balance: %w", err)
	}

	queryWithdrawn := `
	SELECT COALESCE(SUM(amount), 0)
	FROM withdrawals
	WHERE user_login=$1
	`

	var withdrawn float64
	if err := r.pool.QueryRowContext(ctx, queryWithdrawn, login).Scan(&withdrawn); err != nil {
		return core_domain.Balance{}, fmt.Errorf("select withdrawn: %w", err)
	}

	balanceModel := BalanceModel{
		Current:   balance,
		Withdrawn: withdrawn,
	}

	balanceDomain := BalanceDomainFromModel(balanceModel)

	return balanceDomain, nil
}
