package app

import (
	"context"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-playground/validator/v10"
)

type appCtxKey struct{}

func AppFromContext(ctx context.Context) *App {
	return ctx.Value(appCtxKey{}).(*App)
}

func ContextWithApp(ctx context.Context, app *App) context.Context {
	ctx = context.WithValue(ctx, appCtxKey{}, app)
	return ctx
}

type App struct {
	ctx context.Context
	cfg *AppConfig

	router    *chi.Mux
	apiRouter *chi.Mux
	jwtauth   *jwtauth.JWTAuth
	validator *validator.Validate

	// lazy init
	dbOnce sync.Once
	db     *DB
}

func NewApp(ctx context.Context) *App {
	app := &App{}

	app.ctx = ContextWithApp(ctx, app)
	app.initAppConfig()
	app.jwtauth = jwtauth.New("HS256", []byte(app.cfg.JWTSecret), nil)
	app.initValidate()
	app.initRouter()

	return app
}

func (app *App) Context() context.Context {
	return app.ctx
}

func (app *App) Config() *AppConfig {
	return app.cfg
}

func (app *App) Router() *chi.Mux {
	return app.router
}

func (app *App) APIRouter() *chi.Mux {
	return app.apiRouter
}

func (app *App) RouteList() error {
	return app.routeList()
}

func (app *App) WithDBConnect() *App {
	app.DB()

	return app
}

func (app *App) DB() *DB {
	app.dbOnce.Do(func() {
		app.db = app.initDB()
	})

	return app.db
}

func (app *App) JWTAuth() *jwtauth.JWTAuth {
	return app.jwtauth
}

func (app *App) Validate(i interface{}) error {
	return app.validator.Struct(i)
}
