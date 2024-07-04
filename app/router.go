package app

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

var initRoutes []func(app *App)

func (app *App) initRouter() {
	app.router = chi.NewRouter()

	app.router.Use(
		render.SetContentType(render.ContentTypeJSON),
		cors.AllowAll().Handler,
		middleware.RequestID,
		middleware.RealIP,
		middleware.Logger,
		middleware.Compress(5),
		middleware.Heartbeat("/ping"),
		middleware.Recoverer,
	)

	app.apiRouter = chi.NewMux()
	app.apiRouter.Use(app.router.Middlewares()...)

	app.router.Mount(app.cfg.APIPrefix, app.apiRouter)

	// init routes
	for i := range initRoutes {
		initRoutes[i](app)
	}
}

func OnInitRoutes(fn func(app *App)) {
	initRoutes = append(initRoutes, fn)
}

func (app *App) routeList() error {
	walkFunc := func(method string, route string, handler http.Handler, middleware ...func(http.Handler) http.Handler) error {
		route = strings.Replace(route, "/*/", "/", -1)
		fmt.Printf("%s %s %s\n", method, strings.Repeat(" ", 7-len(method)), route)
		return nil
	}

	if err := chi.Walk(app.router, walkFunc); err != nil {
		fmt.Printf("Logging err: %s\n", err.Error())
		return err
	}

	return nil
}
