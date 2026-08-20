package auth_transport_http

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	apperrors "github.com/Kosvu/gophermart/internal/core/errors"
)

type RegisterUserRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (h *AuthHTTPHandler) Register(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var request RegisterUserRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Authorization: Bearer <токен>
	token, err := h.authService.Register(ctx, request.Login, request.Password)

	if err != nil {
		switch {
		case errors.Is(err, apperrors.ErrInvalidInput):
			http.Error(w, apperrors.ErrInvalidInput.Error(), http.StatusBadRequest)
			return
		case errors.Is(err, apperrors.ErrLoginTaken):
			http.Error(w, apperrors.ErrLoginTaken.Error(), http.StatusConflict)
			return
		default:
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
	}

	tokenString := fmt.Sprintf("Bearer %s", token)

	w.Header().Set("Authorization", tokenString)
	w.WriteHeader(http.StatusOK)
}
