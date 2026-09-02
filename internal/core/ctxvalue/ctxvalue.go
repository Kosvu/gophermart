package core_ctxvalue

import "context"

type ctxKey struct{}

func WithLogin(ctx context.Context, login string) context.Context {
	return context.WithValue(ctx, ctxKey{}, login)
}

func LoginFromContext(ctx context.Context) (string, bool) {
	login := ctx.Value(ctxKey{})

	loginString, ok := login.(string)

	if !ok {
		return "", false
	}

	return loginString, true
}
