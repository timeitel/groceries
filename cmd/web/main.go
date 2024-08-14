package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/timeitel/groceries/internal/domain/shopper"
	"github.com/timeitel/groceries/internal/services"
	"github.com/timeitel/groceries/internal/views"
	"github.com/timeitel/groceries/internal/views/home"
)

func main() {
	shopperRepo := shopper.NewLibSqlRepository()
	service := services.NewUser(shopperRepo)

	e := echo.New()
	e.Renderer = views.NewTemplate("internal/views/**/*.html")
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/static", "static")

	e.GET("/", func(c echo.Context) error {
		return home.HomeHandler(c, &service)
	})

	e.POST("/products/:id", func(c echo.Context) error {
		return home.AddProductToCartHandler(c, &service)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
