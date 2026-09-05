package accrual_repository

import (
	"context"
	"fmt"
)

const BatchLimit = 10

func (r *AccrualRepository) GetOrders(ctx context.Context, limit int) ([]string, error) {
	query := `
	SELECT number_order FROM orders
	WHERE status IN ('NEW', 'PROCESSING')
	LIMIT $1
	`

	rows, err := r.pool.QueryContext(ctx, query, limit)

	if err != nil {
		return nil, fmt.Errorf("select rows: %w", err)
	}
	defer rows.Close()

	var numbers []string

	for rows.Next() {
		var number string

		if err := rows.Scan(&number); err != nil {
			return nil, fmt.Errorf("row scan: %w", err)
		}

		numbers = append(numbers, number)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows next: %w", err)
	}

	return numbers, nil
}
