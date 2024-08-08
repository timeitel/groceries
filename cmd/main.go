package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/timeitel/groceries/internal/domain/user"
	"github.com/timeitel/groceries/internal/services"
	"github.com/timeitel/groceries/internal/views"
	"github.com/timeitel/groceries/internal/views/home"
)

func main() {
	repo := user.NewLibSqlRepository()
	service := services.NewShopping(repo)

	e := echo.New()

	// e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Renderer = views.NewTemplate("internal/views/*.html")

	e.Static("/static", "static")

	e.GET("/", home.Index(service.GetHomeData()))

	e.POST("/items/:name", home.AddItem(service.AddItem))

	e.Logger.Fatal(e.Start(":8080"))
}
