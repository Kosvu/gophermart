package orders_repository

import "database/sql"

type OrdersRepository struct {
	pool *sql.DB
}

func NewOrdersRepository(pool *sql.DB) *OrdersRepository {
	return &OrdersRepository{
		pool: pool,
	}
}
