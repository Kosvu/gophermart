package auth_transport_http

import "context"

type AuthHTTPHandler struct {
	authService AuthService
}

type AuthService interface {
	Register(
		ctx context.Context,
		login string,
		password string,
	) (string, error)
}

func NewAuthHTTPHandler(authService AuthService) *AuthHTTPHandler {
	return &AuthHTTPHandler{
		authService: authService,
	}
}
