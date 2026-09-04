package orders_service

import (
	"context"

	core_domain "github.com/Kosvu/gophermart/internal/core/domain"
)

type OrdersService struct {
	ordersRepository ordersRepository
}

type ordersRepository interface {
	Load(ctx context.Context, login, number string) error
	GetOrders(ctx context.Context, login string) ([]core_domain.Order, error)
}

func NewOrdersService(ordersRepository ordersRepository) *OrdersService {
	return &OrdersService{
		ordersRepository: ordersRepository,
	}
}
