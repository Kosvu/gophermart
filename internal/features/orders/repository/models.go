package orders_repository

import (
	"time"

	core_domain "github.com/Kosvu/gophermart/internal/core/domain"
)

type OrderModel struct {
	Number     string
	Status     string
	Accrual    *float64
	UploadedAt time.Time
}

func orderDomainFromModel(order OrderModel) core_domain.Order {
	return core_domain.Order{
		Number:     order.Number,
		Status:     order.Status,
		Accrual:    order.Accrual,
		UploadedAt: order.UploadedAt,
	}
}

func orderDomainsFromModels(orders []OrderModel) []core_domain.Order {
	domains := make([]core_domain.Order, len(orders))

	for i, el := range orders {
		domains[i] = orderDomainFromModel(el)
	}

	return domains
}
