package orders_transport

import (
	"errors"
	"io"
	"net/http"
	"strings"

	core_ctxvalue "github.com/Kosvu/gophermart/internal/core/ctxvalue"
	apperrors "github.com/Kosvu/gophermart/internal/core/errors"
)

func (h *OrdersHTTPHandler) Load(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	body, err := io.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) // 500 ошибка сервера
		return
	}

	login, ok := core_ctxvalue.LoginFromContext(ctx)

	if !ok {
		w.WriteHeader(http.StatusInternalServerError) // 500 ошибка сервера
		return
	}

	number := strings.TrimSpace(string(body))

	if err := h.ordersService.Load(ctx, login, number); err != nil {
		switch {
		case errors.Is(err, apperrors.ErrOrderAlreadyLoaded):
			w.WriteHeader(http.StatusOK) // 200 OK
			return
		case errors.Is(err, apperrors.ErrOrderTakenByAnother):
			w.WriteHeader(http.StatusConflict) // 409 конфликт
			return
		case errors.Is(err, apperrors.ErrInvalidNumber):
			w.WriteHeader(http.StatusUnprocessableEntity) // 422 неверный формат, неподдерживаемая сущность
			return
		default:
			w.WriteHeader(http.StatusInternalServerError) // 500 ошибка сервера
			return
		}
	}

	w.WriteHeader(http.StatusAccepted) // 202 принят в обработку
}
