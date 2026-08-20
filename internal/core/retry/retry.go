package retry

import (
	"context"
	"time"
)

func Retry(ctx context.Context, op func() error) error {
	delays := []time.Duration{1 * time.Second, 3 * time.Second, 5 * time.Second}
	var err error
	for i := 0; i < len(delays); i++ {
		if err = op(); err == nil {
			return nil
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(delays[i]):

		}
	}

	return err
}
