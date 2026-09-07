package balance_service

import (
	"context"
	"fmt"

	core_domain "github.com/Kosvu/gophermart/internal/core/domain"
)

func (s *BalanceService) GetBalance(ctx context.Context, login string) (core_domain.Balance, error) {
	balance, err := s.balanceRepository.GetBalance(ctx, login)

	if err != nil {
		return core_domain.Balance{}, fmt.Errorf("get balance: %w", err)
	}

	return balance, nil
}
