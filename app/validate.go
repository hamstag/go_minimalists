package app

import (
	"github.com/go-playground/validator/v10"
)

func (app *App) initValidate() {
	app.validator = validator.New(validator.WithRequiredStructEnabled())
}
