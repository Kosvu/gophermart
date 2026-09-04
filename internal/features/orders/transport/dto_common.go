package orders_transport

import (
	"time"

	core_domain "github.com/Kosvu/gophermart/internal/core/domain"
)

type OrderDTOResponse struct {
	Number     string    `json:"number"`
	Status     string    `json:"status"`
	Accrual    *float64  `json:"accrual,omitempty"`
	UploadedAt time.Time `json:"uploaded_at"`
}

func orderDTOFromDomain(order core_domain.Order) OrderDTOResponse {
	return OrderDTOResponse{
		Number:     order.Number,
		Status:     order.Status,
		Accrual:    order.Accrual,
		UploadedAt: order.UploadedAt,
	}
}

func orderDTOsFromDomain(orders []core_domain.Order) []OrderDTOResponse {
	dtos := make([]OrderDTOResponse, len(orders))

	for i, el := range orders {
		dtos[i] = orderDTOFromDomain(el)
	}

	return dtos
}
