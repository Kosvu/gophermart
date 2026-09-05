package accrual_worker

import (
	"context"
	"errors"
	"log"
	"time"

	core_accrual "github.com/Kosvu/gophermart/internal/core/accrualclient"
	apperrors "github.com/Kosvu/gophermart/internal/core/errors"
	accrual_repository "github.com/Kosvu/gophermart/internal/features/accrual/repository"
)

const batchLimit = 10
const DefaultInterval = 3 * time.Second

type Worker struct {
	Client            *core_accrual.Client
	AccrualRepository *accrual_repository.AccrualRepository
	Interval          time.Duration
}

func NewWorker(client *core_accrual.Client, accrualRepository *accrual_repository.AccrualRepository, interval time.Duration) *Worker {
	return &Worker{
		Client:            client,
		AccrualRepository: accrualRepository,
		Interval:          interval,
	}
}

func (w *Worker) ProcessOnce(ctx context.Context) {
	numbers, err := w.AccrualRepository.GetOrders(ctx, batchLimit)

	if err != nil {
		log.Printf("get orders: %v", err)
		return
	}

	for _, number := range numbers {

		accrual, err := w.Client.OrderInfo(ctx, number)

		if err != nil {
			if errors.Is(err, apperrors.ErrNoContent) {
				continue
			}

			var rateErr *core_accrual.RateLimitError
			if errors.As(err, &rateErr) {
				// time.Sleep(rateErr.RetryAfter) даже после SIGTERM будет спать и программа закончится позже

				select {
				case <-ctx.Done():
					return
				case <-time.After(rateErr.RetryAfter):
					return
				}
			}

			log.Printf("order info: %v", err)
			continue
		}

		status := accrual.Status
		if status == "REGISTERED" {
			status = "PROCESSING"
		}

		if err := w.AccrualRepository.UpdateOrder(ctx, number, status, accrual.Accrual); err != nil {
			log.Printf("update order: %v", err)
		}
	}
}

func (w *Worker) Run(ctx context.Context) {
	ticker := time.NewTicker(w.Interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			w.ProcessOnce(ctx)
		}
	}

}
