package main

import (
	"fmt"
	"go-minimalists/app"
	_ "go-minimalists/features/product"
	_ "go-minimalists/features/user"
	"go-minimalists/middleware"
	"go-minimalists/util/security"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
)

func init() {
	app.OnInitRoutes(func(app *app.App) {
		r := app.Router()

		r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
			hash := security.HashPassword("hello hamstag")
			fmt.Printf("%s\n", hash)

			match := security.CheckPasswordHash("hello hamstag", hash)
			fmt.Println(match)

			render.JSON(w, r, render.M{
				"message": "Hello! Hamstag.",
			})
		})

		r.Group(func(r chi.Router) {
			r.Use(middleware.Authenticator(app.JWTAuth()).Handler)

			r.Get("/private/hello", func(w http.ResponseWriter, r *http.Request) {
				_, claims, _ := jwtauth.FromContext(r.Context())

				fmt.Println(claims)

				render.JSON(w, r, render.M{
					"message": "Hello! Private Hamstag.",
				})
			})
		})
	})
}
