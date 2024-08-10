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

	e.GET("/", home.Index(service.GetCart))

	e.POST("/items/:id", home.AddItem(service.AddItem))

	e.Logger.Fatal(e.Start(":8080"))
}
