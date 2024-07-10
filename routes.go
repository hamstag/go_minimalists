package main

import (
	"encoding/json"
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

		r.Get("/sabaidee", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("sabaidee hamstag"))
		})

		r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
			// Hash
			hash := security.HashPassword("hello hamstag")
			fmt.Printf("hash: %s\n", hash)

			match := security.CheckPasswordHash("hello hamstag", hash)
			fmt.Printf("match: %t\n", match)

			// Encryption
			text := map[string]interface{}{
				"hello":    "Hamstag",
				"sabaidee": "ສະບາຍດີ",
				"number":   1.23,
			}

			textString, _ := json.Marshal(text)

			encrypted := security.Encrypt(string(textString), app.Config().EncryptionSecret)
			fmt.Printf("encrypted: %s\n", encrypted)

			decrypted := security.Decrypt(encrypted, app.Config().EncryptionSecret)
			fmt.Printf("decrypted: %s\n", decrypted)

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
