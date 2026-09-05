package accrual_repository

import "database/sql"

type AccrualRepository struct {
	pool *sql.DB
}

func NewAccrualRepository(pool *sql.DB) *AccrualRepository {
	return &AccrualRepository{
		pool: pool,
	}
}
