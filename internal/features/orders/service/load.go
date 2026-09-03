package orders_service

import (
	"context"
	"fmt"

	apperrors "github.com/Kosvu/gophermart/internal/core/errors"
)

func isValidLuhn(number string) bool {
	sum := 0
	alt := false

	if len(number) == 0 {
		return false
	}

	for i := len(number) - 1; i >= 0; i-- {
		ch := number[i]

		if ch < '0' || ch > '9' {
			return false
		}

		val := int(ch - '0')

		if alt {
			val *= 2
			if val > 9 {
				val -= 9
			}
		}

		sum += val
		alt = !alt
	}

	return sum%10 == 0
}

func (s *OrdersService) Load(ctx context.Context, login string, number string) error {
	if !isValidLuhn(number) {
		return fmt.Errorf("invalid number: %w", apperrors.ErrInvalidNumber)
	}

	if err := s.ordersRepository.Load(ctx, login, number); err != nil {
		return fmt.Errorf("order repository: %w", err)
	}

	return nil
}
