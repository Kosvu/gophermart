package balance_transport

import (
	"bytes"
	"encoding/json"
	"net/http"

	core_ctxvalue "github.com/Kosvu/gophermart/internal/core/ctxvalue"
)

type GetBalanceResponse BalanceDTO

func (h *BalanceHTTPHandler) GetBalance(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	login, ok := core_ctxvalue.LoginFromContext(ctx)

	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	balanceDomain, err := h.balanceService.GetBalance(ctx, login)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	getBalanceDTO := balanceDTOFromDomain(balanceDomain)
	getBalanceResponse := GetBalanceResponse(getBalanceDTO)

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(getBalanceResponse); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(buf.Bytes())
}
