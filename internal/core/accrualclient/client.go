package core_accrual

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	apperrors "github.com/Kosvu/gophermart/internal/core/errors"
)

type AccrualResponse struct {
	Order   string   `json:"order"`
	Accrual *float64 `json:"accrual"`
	Status  string   `json:"status"`
}

type Client struct {
	client *http.Client
	addr   string
}

func NewClient(addr string) *Client {

	if !strings.HasPrefix(addr, "http://") && !strings.HasPrefix(addr, "https://") {
		addr = "http://" + addr
	}

	return &Client{
		client: &http.Client{Timeout: 10 * time.Second},
		addr:   addr,
	}
}

type RateLimitError struct {
	RetryAfter time.Duration
}

func (e *RateLimitError) Error() string {
	return fmt.Sprintf("rate limited, retry after %s", e.RetryAfter)
}

func (c *Client) OrderInfo(ctx context.Context, number string) (*AccrualResponse, error) {

	url := fmt.Sprintf("%s/api/orders/%s", c.addr, number)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	resp, err := c.client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}

	defer resp.Body.Close()

	var accrual AccrualResponse

	switch resp.StatusCode {

	case http.StatusNoContent:
		return nil, apperrors.ErrNoContent
	case http.StatusTooManyRequests: // 429 (при этом статус коде в header приходит поле "Retry-After")
		retryAfter := resp.Header.Get("Retry-After")
		retryAfterInt, err := strconv.Atoi(retryAfter)

		if err != nil {
			return nil, &RateLimitError{RetryAfter: 60 * time.Second}
		}

		return nil, &RateLimitError{RetryAfter: time.Duration(retryAfterInt) * time.Second}
	case http.StatusOK:
		if err := json.NewDecoder(resp.Body).Decode(&accrual); err != nil {
			return nil, fmt.Errorf("json decode %w", err)
		}

		return &accrual, nil
	default:
		return nil, apperrors.ErrInternalServer
	}
}
