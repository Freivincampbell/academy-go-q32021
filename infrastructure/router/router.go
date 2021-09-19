package router

import (
	"academy-go-q32021/interface/controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewRouter(e *echo.Echo, c controller.App) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/read-users", func(context echo.Context) error {
		return c.User.ReadUsers(context)
	})
	e.GET("/read-users-by-key", func(context echo.Context) error {
		return c.User.ReadUsersByKey(context)
	})
	e.GET("/users", func(context echo.Context) error {
		return c.User.GetUsers(context)
	})

	return e
}
