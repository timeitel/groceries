package shopper

import (
	"github.com/labstack/echo/v4"
	"github.com/timeitel/groceries/internal/services"
)

func Routes(e *echo.Echo) {
	service := services.NewShopper()

	e.GET("/", func(c echo.Context) error {
		return getHome(c, &service)
	})

	e.GET("/products/:id", func(c echo.Context) error {
		return getItem(c, &service)
	})

	e.POST("/products/:id", func(c echo.Context) error {
		return addItem(c, &service)
	})
}
