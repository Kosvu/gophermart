package orders_service

import "context"

type OrdersService struct {
	ordersRepository ordersRepository
}

type ordersRepository interface {
	Load(ctx context.Context, login, number string) error
}

func NewOrdersRepository(ordersRepository ordersRepository) *OrdersService {
	return &OrdersService{
		ordersRepository: ordersRepository,
	}
}
