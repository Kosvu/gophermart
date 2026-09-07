package balance_repository

import "database/sql"

type BalanceRepository struct {
	pool *sql.DB
}

func NewBalanceRepository(pool *sql.DB) *BalanceRepository {
	return &BalanceRepository{
		pool: pool,
	}
}
