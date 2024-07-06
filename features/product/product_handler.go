package product

import (
	"go-minimalists/app"
	"go-minimalists/util/httperror"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type ProductHandler struct {
	app *app.App
}

func NewProductHandler(app *app.App) *ProductHandler {
	return &ProductHandler{app: app}
}

func (h ProductHandler) Index(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, render.M{
		"message": "Index",
	})
}

func (h ProductHandler) Store(w http.ResponseWriter, r *http.Request) {
	rd := &productRequestData{}

	if err := render.Bind(r, rd); err != nil {
		render.Render(w, r, httperror.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, render.M{
		"message": "rd",
	})
}

func (h ProductHandler) Show(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id != "1" {
		render.Render(w, r, httperror.ErrNotFound)
		return
	}

	render.JSON(w, r, render.M{
		"message": "Show " + id,
	})
}

func (h ProductHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id != "1" {
		render.Render(w, r, httperror.ErrNotFound)
		return
	}

	rd := &productRequestData{}

	if err := render.Bind(r, rd); err != nil {
		render.Render(w, r, httperror.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, render.M{
		"message": "Update " + id,
	})
}

func (h ProductHandler) Destroy(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	render.JSON(w, r, render.M{
		"message": "Destroy " + id,
	})
}
