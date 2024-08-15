package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/timeitel/groceries/internal/api"
	"github.com/timeitel/groceries/internal/api/handlers"
	"github.com/timeitel/groceries/internal/domain/admin"
	"github.com/timeitel/groceries/internal/domain/shopper"
	"github.com/timeitel/groceries/internal/services"
)

func main() {
	shopperRepo := shopper.NewLibSqlRepository()
	shopperService := services.NewUser(shopperRepo)
	adminService := services.NewAdmin(admin.NewLibSqlRepository(), shopperRepo)

	e := echo.New()
	e.Renderer = api.NewTemplates()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/static", "static")

	e.GET("/", func(c echo.Context) error {
		return handlers.ShopperHome(c, &shopperService)
	})

	e.POST("/products/:id", func(c echo.Context) error {
		return handlers.ShopperAddProduct(c, &shopperService)
	})

	protected := e.Group("/admin")
	protected.Use(api.Middleware)
	protected.GET("", func(c echo.Context) error {
		return handlers.AdminGetHome(c, &adminService)
	})
	protected.POST("/products", func(c echo.Context) error {
		return handlers.AdminCreateProduct(c, &adminService)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
