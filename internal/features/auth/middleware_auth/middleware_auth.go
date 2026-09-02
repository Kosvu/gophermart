package auth_middleware

import (
	"net/http"
	"strings"

	core_ctxvalue "github.com/Kosvu/gophermart/internal/core/ctxvalue"
	"github.com/Kosvu/gophermart/internal/features/auth/token"
)

type Middleware func(http.Handler) http.Handler

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

			ctx := core_ctxvalue.WithLogin(r.Context(), login)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
