package auth_middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/Kosvu/gophermart/internal/features/auth/token"
)

type Middleware func(http.Handler) http.Handler

type ctxKey struct{}

func LoginFromContext(ctx context.Context) (string, bool) {
	login := ctx.Value(ctxKey{})

	loginString, ok := login.(string)

	if !ok {
		return "", false
	}

	return loginString, true
}

func Auth(t *token.Token) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenBearer := r.Header.Get("Authorization")
			token, ok := strings.CutPrefix(tokenBearer, "Bearer ")

			if !ok {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			login, err := t.Validate(token)

			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), ctxKey{}, login)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
