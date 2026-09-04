package orders_transport

import (
	"bytes"
	"encoding/json"
	"net/http"

	core_ctxvalue "github.com/Kosvu/gophermart/internal/core/ctxvalue"
)

func (h *OrdersHTTPHandler) GetOrders(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	login, ok := core_ctxvalue.LoginFromContext(ctx)

	if !ok {
		w.WriteHeader(http.StatusInternalServerError) // 500 ошибка сервера
		return
	}

	ordersDomain, err := h.ordersService.GetOrders(ctx, login)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if len(ordersDomain) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	ordersDTO := orderDTOsFromDomain(ordersDomain)

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(ordersDTO); err != nil {
		w.WriteHeader(http.StatusInternalServerError) // 500 ошибка сервера
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(buf.Bytes())
}
