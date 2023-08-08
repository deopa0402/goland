package api

import (
	"github.com/labstack/echo/v4"
	"main/handler"
	"main/middleware"
)

func InitAPI(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})

	//use
	e.POST("/login", handler.Login)
	e.POST("/users", handler.CreateUser)

	g := e.Group("")
	g.Use(middleware.JWTMiddleware)
	g.GET("/users/:id", handler.GetUserByID)
	g.GET("/users", handler.GetUser)

	//e.PUT("/users/:id", handler.UpdateUser)
	//e.DELETE("/users/:id", handler.DeleteUser)
}
