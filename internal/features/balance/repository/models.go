package balance_repository

import core_domain "github.com/Kosvu/gophermart/internal/core/domain"

type BalanceModel struct {
	Current   float64
	Withdrawn float64
}

func BalanceDomainFromModel(balance BalanceModel) core_domain.Balance {
	return core_domain.Balance{
		Current:   balance.Current,
		Withdrawn: balance.Withdrawn,
	}
}
