package admin

import (
	"github.com/labstack/echo/v4"
	"github.com/timeitel/groceries/internal/services"
)

func Routes(e *echo.Echo) {
	service := services.NewAdmin()
	protected := e.Group("/admin")

	protected.GET("", func(c echo.Context) error {
		return getHome(c, &service)
	})
	protected.GET("/items/:id", func(c echo.Context) error {
		return getItem(c, &service)
	})

	protected.POST("/items", func(c echo.Context) error {
		return createItem(c, &service)
	})

	protected.PUT("/items/:id", func(c echo.Context) error {
		return updateItem(c, &service)
	})

	protected.DELETE("/items/:id", func(c echo.Context) error {
		return deleteItem(c, &service)
	})

}
