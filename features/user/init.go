package user

import (
	"go-minimalists/app"

	"github.com/go-chi/chi/v5"
)

func init() {
	app.OnInitRoutes(func(app *app.App) {
		h := NewUserHandler(app)

		r := app.APIRouter()

		r.Route("/users", func(r chi.Router) {
			r.Get("/", h.Index)
			r.Post("/", h.Store)
			r.Get("/{id}", h.Show)
			r.Put("/{id}", h.Update)
			r.Delete("/{id}", h.Destroy)
		})
	})
}
