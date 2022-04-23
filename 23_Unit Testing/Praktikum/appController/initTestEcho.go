package appController

import (
	"Praktikum/appMiddleware"
	"Praktikum/appModel"

	"github.com/labstack/echo/v4"
)

func initTestEcho() (e *echo.Echo, pc PersonController) {
	personModel := appModel.NewPersonMemModel()
	e = echo.New()
	appMiddleware.AddGlobalMiddlewares(e)
	jwtSecret := "rahasiaBanget"
	pc = HandleRoutes(e, jwtSecret, personModel)
	return e, pc
}
