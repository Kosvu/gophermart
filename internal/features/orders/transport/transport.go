package orders_transport

import "context"

type OrdersHTTPHandler struct {
	ordersService OrdersService
}

type OrdersService interface {
	Load(ctx context.Context, login string, number string) error
}

func NewOrdersHTTPHandler(ordersService OrdersService) *OrdersHTTPHandler {
	return &OrdersHTTPHandler{
		ordersService: ordersService,
	}
}
