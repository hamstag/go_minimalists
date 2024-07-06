package app

import (
	"fmt"
	"net/http"
	"os"
	"reflect"
	"runtime"
	"strings"
	"text/tabwriter"

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
	app.router.Mount(app.cfg.APIPrefix, app.apiRouter)

	// init routes
	for _, fn := range initRoutes {
		fn(app)
	}
}

func OnInitRoutes(fn func(app *App)) {
	initRoutes = append(initRoutes, fn)
}

func (app *App) routeList() error {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 4, ' ', 0)
	fmt.Fprintf(w, "%s\t%s\t%s\n", "Method", "Route", "Handler")
	fmt.Fprintf(w, "%s\t%s\t%s\n", "------", "-----", "-------")

	walkFunc := func(method string, route string, handler http.Handler, middleware ...func(http.Handler) http.Handler) error {
		route = strings.Replace(route, "/*/", "/", -1)
		pathHandler := runtime.FuncForPC(reflect.ValueOf(handler).Pointer()).Name()
		replacer := strings.NewReplacer("go-minimalists", "")
		pathHandler = replacer.Replace(pathHandler)

		fmt.Fprintf(w, "%s\t%s\t%s\n", method, route, pathHandler)

		return nil
	}

	if err := chi.Walk(app.router, walkFunc); err != nil {
		fmt.Printf("Logging err: %s\n", err.Error())
		return err
	}

	fmt.Fprintln(w)
	w.Flush()

	return nil
}
