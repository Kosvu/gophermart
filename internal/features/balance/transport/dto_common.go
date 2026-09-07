package balance_transport

import core_domain "github.com/Kosvu/gophermart/internal/core/domain"

type BalanceDTO struct {
	Current   float64 `json:"balance"`
	Withdrawn float64 `json:"withdrawn"`
}

func balanceDTOFromDomain(balanceDomain core_domain.Balance) BalanceDTO {
	return BalanceDTO{
		Current:   balanceDomain.Current,
		Withdrawn: balanceDomain.Withdrawn,
	}
}
