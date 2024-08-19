package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/timeitel/groceries/internal/handlers"
	"github.com/timeitel/groceries/internal/services"
)

func main() {
	shopperService := services.NewShopper()
	adminService := services.NewAdmin()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/static", "internal/web/static")

	e.GET("/", func(c echo.Context) error {
		return handlers.ShopperGetHome(c, &shopperService)
	})
	e.GET("/products/:id", func(c echo.Context) error {
		return handlers.ShopperGetItem(c, &shopperService)
	})

	e.POST("/products/:id", func(c echo.Context) error {
		return handlers.ShopperAddItem(c, &shopperService)
	})

	protected := e.Group("/admin")
	protected.Use(handlers.Middleware)
	protected.GET("", func(c echo.Context) error {
		return handlers.AdminGetHome(c, &adminService)
	})
	protected.GET("/products/:id", func(c echo.Context) error {
		return handlers.AdminGetProduct(c, &adminService)
	})
	protected.POST("/products", func(c echo.Context) error {
		return handlers.AdminCreateProduct(c, &adminService)
	})
	protected.PUT("/products/:id", func(c echo.Context) error {
		return handlers.AdminUpdateProduct(c, &adminService)
	})
	protected.DELETE("/products/:id", func(c echo.Context) error {
		return handlers.AdminDeleteProduct(c, &adminService)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
