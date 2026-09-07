package balance_service

import (
	"context"

	core_domain "github.com/Kosvu/gophermart/internal/core/domain"
)

type BalanceService struct {
	balanceRepository BalanceRepository
}

type BalanceRepository interface {
	GetBalance(ctx context.Context, login string) (core_domain.Balance, error)
}

func NewBalanseService(balanseRepository BalanceRepository) *BalanceService {
	return &BalanceService{
		balanceRepository: balanseRepository,
	}
}
