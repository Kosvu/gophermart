package orders_repository

import (
	"context"
	"fmt"

	core_domain "github.com/Kosvu/gophermart/internal/core/domain"
)

func (r *OrdersRepository) GetOrders(ctx context.Context, login string) ([]core_domain.Order, error) {
	query := `
	SELECT number_order,status,accrual,uploaded_at
	FROM orders
	WHERE user_login=$1
	ORDER BY uploaded_at
	`

	rows, err := r.pool.QueryContext(ctx, query, login)

	if err != nil {
		return nil, fmt.Errorf("select number_order: %w", err)
	}

	defer rows.Close()

	var ordersModel []OrderModel

	for rows.Next() {
		var orderModel OrderModel

		err := rows.Scan(
			&orderModel.Number,
			&orderModel.Status,
			&orderModel.Accrual,
			&orderModel.UploadedAt,
		)

		if err != nil {
			return nil, fmt.Errorf("scan order: %w", err)
		}

		ordersModel = append(ordersModel, orderModel)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("next row: %w", rows.Err())
	}

	ordersDomain := orderDomainsFromModels(ordersModel)

	return ordersDomain, nil
}
