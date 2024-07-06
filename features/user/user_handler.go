package user

import (
	"fmt"
	"go-minimalists/app"
	"go-minimalists/util/httperror"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type UserHandler struct {
	app *app.App
}

func NewUserHandler(app *app.App) *UserHandler {
	return &UserHandler{app: app}
}

func (h UserHandler) Index(w http.ResponseWriter, r *http.Request) {
	// get app from context
	appFromContext := app.AppFromContext(r.Context())
	fmt.Printf("Address: http://%s\n", appFromContext.Config().Address)

	// eg: redis
	h.app.DB().Redis.Set(r.Context(), "hello", "Hamstag", time.Second*60)

	users := []User{}

	h.app.DB().MySql.Limit(10).Find(&users)

	render.JSON(w, r, users)
}

func (h UserHandler) Store(w http.ResponseWriter, r *http.Request) {
	rd := &userRequestData{}

	if err := render.Bind(r, rd); err != nil {
		render.Render(w, r, httperror.ErrInvalidRequest(err))
		return
	}

	user := rd.toModel()

	if err := h.app.DB().MySql.Create(&user).Error; err != nil {
		render.Render(w, r, httperror.ErrRender(err))
		return
	}

	render.JSON(w, r, render.M{
		"id": user.ID,
	})
}

func (h UserHandler) Show(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id != "1" {
		render.Render(w, r, httperror.ErrNotFound)
		return
	}

	render.JSON(w, r, render.M{
		"message": "Show " + id,
	})
}

func (h UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id != "1" {
		render.Render(w, r, httperror.ErrNotFound)
		return
	}

	rd := &userRequestData{}

	if err := render.Bind(r, rd); err != nil {
		render.Render(w, r, httperror.ErrRender(err))
		return
	}

	render.JSON(w, r, render.M{
		"status": "Update " + id,
	})
}

func (h UserHandler) Destroy(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	render.JSON(w, r, render.M{
		"status": "Destroy " + id,
	})
}
