package server

import (
	"Praktikum/interfaces/users"

	"github.com/labstack/echo/v4"
)

func Server() *echo.Echo {
	app := echo.New()

	users.Routes(app)

	return app
}
