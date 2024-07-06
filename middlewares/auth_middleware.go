package middlewares

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

type AuthMiddleware struct {
	ja *jwtauth.JWTAuth
}

func Authenticator(ja *jwtauth.JWTAuth) *AuthMiddleware {
	return &AuthMiddleware{ja: ja}
}

func (auth *AuthMiddleware) Handler(next http.Handler) http.Handler {
	middlewares := []func(http.Handler) http.Handler{
		jwtauth.Verify(auth.ja, jwtauth.TokenFromHeader, jwtauth.TokenFromQuery),
		jwtauth.Authenticator(auth.ja),
	}

	return chi.Chain(middlewares...).Handler(next)
}
