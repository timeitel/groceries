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

	e.Static("/static", "internal/web/static")

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
