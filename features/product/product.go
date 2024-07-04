package product

import (
	"go-minimalists/app"
	"net/http"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model

	Name     string
	Surname  string
	Username string
}

type productRequestData struct {
	Name string `json:"name" validate:"required"`
	Qty  string `json:"qty" validate:"required"`
}

func (rd *productRequestData) Bind(r *http.Request) error {
	return app.AppFromContext(r.Context()).Validate(rd)
}
