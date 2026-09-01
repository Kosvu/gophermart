package auth_transport_http

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	apperrors "github.com/Kosvu/gophermart/internal/core/errors"
)

type LoginUserRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (h *AuthHTTPHandler) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var request LoginUserRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token, err := h.authService.Login(ctx, request.Login, request.Password)

	if err != nil {
		switch {
		case errors.Is(err, apperrors.ErrInvalidInput):
			w.WriteHeader(http.StatusBadRequest)
			return
		case errors.Is(err, apperrors.ErrInvalidCredentials):
			w.WriteHeader(http.StatusUnauthorized)
			return
		default:
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	tokenString := fmt.Sprintf("Bearer %s", token)

	w.Header().Set("Authorization", tokenString)
	w.WriteHeader(http.StatusOK)
}
