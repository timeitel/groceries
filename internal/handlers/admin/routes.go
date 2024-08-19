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
	protected.GET("/products/:id", func(c echo.Context) error {
		return getProduct(c, &service)
	})

	protected.POST("/products", func(c echo.Context) error {
		return createProduct(c, &service)
	})

	protected.PUT("/products/:id", func(c echo.Context) error {
		return updateProduct(c, &service)
	})

	protected.DELETE("/products/:id", func(c echo.Context) error {
		return deleteProduct(c, &service)
	})

}
