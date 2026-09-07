package balance_transport

import (
	"context"

	core_domain "github.com/Kosvu/gophermart/internal/core/domain"
)

type BalanceHTTPHandler struct {
	balanceService BalanceService
}

type BalanceService interface {
	GetBalance(ctx context.Context, login string) (core_domain.Balance, error)
}

func NewBalanceHTTPHandler(balanceService BalanceService) *BalanceHTTPHandler {
	return &BalanceHTTPHandler{
		balanceService: balanceService,
	}
}
