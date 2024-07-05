package user

import (
	"go-minimalists/app"
	"net/http"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Name     string
	Surname  string
	Username string
}

type userRequestData struct {
	Name     string `json:"name" validate:"required"`
	Surname  string `json:"surname" validate:"required"`
	Username string `json:"username" validate:"required"`
}

func (rd *userRequestData) Bind(r *http.Request) error {
	return app.AppFromContext(r.Context()).Validate(rd)
}

func (rd *userRequestData) toModel() User {
	user := User{
		Name:     rd.Name + "_suffix",
		Surname:  rd.Surname,
		Username: rd.Username,
	}

	return user
}
