package auth

import "context"

type Token string

type contextKey struct {
	name string
}

var tokenKey = &contextKey{"authTokenKey"}

// WithToken returns a context with the given token.
func WithToken(ctx context.Context, token Token) context.Context {
	return context.WithValue(ctx, tokenKey, token)
}

// GetToken returns the token from the context
// and a boolean indicating whether the token was present in the context.
func GetToken(ctx context.Context) (Token, bool) {
	t, ok := ctx.Value(tokenKey).(Token)
	return t, ok
}