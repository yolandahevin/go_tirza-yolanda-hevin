package auth

import (
	"Praktikum/infrastructures/database/sqlite"
	ar "Praktikum/repositories/auth"

	"github.com/labstack/echo/v4"
)

func Routes(echo *echo.Echo) {
	db := sqlite.ConnectDB()
	authRepo := ar.NewAuthSQLiteRepository(db)
	authHandler := NewAuthHandler(authRepo)

	echo.POST("/auth/login", authHandler.LoginHandler)
}
