package router

import (
	"academy-go-q32021/interface/controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewRouter(e *echo.Echo, c controller.App) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Read users from CSV
	e.GET("/read-users", func(context echo.Context) error {
		return c.User.ReadUsers(context)
	})
	// Read users from CSV BY KEY
	e.GET("/read-users-by-key", func(context echo.Context) error {
		return c.User.ReadUsersByKey(context)
	})
	// Read users from API
	e.GET("/users", func(context echo.Context) error {
		return c.User.GetUsers(context)
	})
	// Read users from APY BY UNIQ ID
	e.GET("/users/:id", func(context echo.Context) error {
		return c.User.GetUserById(context)
	})
	// Read users from CSV CONCURRENTLY
	e.GET("/users/concurrently", func(context echo.Context) error {
		return c.User.GetUsersConcurrently(context)
	})

	return e
}
