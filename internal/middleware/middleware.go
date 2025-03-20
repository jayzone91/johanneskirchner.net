package middleware

import (
	"net/http"
	"slices"
)

type Middleware func(next http.Handler) http.Handler

func Chain(middlewares ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for _, x := range slices.Backward(middlewares) {
			next = x(next)
		}
		return next
	}
}
