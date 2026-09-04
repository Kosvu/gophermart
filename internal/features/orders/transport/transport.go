package orders_transport

import (
	"context"

	core_domain "github.com/Kosvu/gophermart/internal/core/domain"
)

type OrdersHTTPHandler struct {
	ordersService OrdersService
}

type OrdersService interface {
	Load(ctx context.Context, login string, number string) error
	GetOrders(ctx context.Context, login string) ([]core_domain.Order, error)
}

func NewOrdersHTTPHandler(ordersService OrdersService) *OrdersHTTPHandler {
	return &OrdersHTTPHandler{
		ordersService: ordersService,
	}
}
