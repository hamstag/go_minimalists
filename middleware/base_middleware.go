package middleware

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

func BaseMiddleware(next http.Handler) http.Handler {
	middlewares := []func(http.Handler) http.Handler{
		render.SetContentType(render.ContentTypeJSON),
		cors.AllowAll().Handler,
		middleware.RequestID,
		middleware.RealIP,
		middleware.Logger,
		middleware.Compress(5),
		middleware.Heartbeat("/ping"),
		middleware.Recoverer,
	}

	return chi.Chain(middlewares...).Handler(next)
}
