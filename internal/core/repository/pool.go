package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Kosvu/gophermart/internal/core/retry"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func NewPool(ctx context.Context, dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)

	if err != nil {
		return nil, fmt.Errorf("open pool: %w", err)
	}

	db.SetMaxOpenConns(20) //сколько всего соединений
	db.SetMaxIdleConns(20) //сколько готов держать открытым
	db.SetConnMaxLifetime(5 * time.Minute)

	err = retry.Retry(ctx, func() error {
		return db.PingContext(ctx)
	})

	if err != nil {
		db.Close()
		return nil, fmt.Errorf("ping db: %w", err)
	}

	return db, nil
}
