package users

import (
	"Praktikum/infrastructures/database/sqlite"
	m "Praktikum/infrastructures/http/middleware"
	ur "Praktikum/repositories/users"

	"github.com/labstack/echo/v4"
)

func Routes(echo *echo.Echo) {
	db := sqlite.ConnectDB()
	userRepo := ur.NewUserSQLiteRepository(db)
	userHandler := NewUserHandler(userRepo)

	echo.GET("/users", userHandler.GetUsersHandler, m.JWTMiddleware())
	echo.POST("/users", userHandler.CreateUserHandler)
}
