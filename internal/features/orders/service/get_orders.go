package orders_service

import (
	"context"
	"fmt"

	core_domain "github.com/Kosvu/gophermart/internal/core/domain"
)

func (s *OrdersService) GetOrders(ctx context.Context, login string) ([]core_domain.Order, error) {
	orders, err := s.ordersRepository.GetOrders(ctx, login)

	if err != nil {
		return nil, fmt.Errorf("orders repository: %w", err)
	}

	return orders, nil
}
