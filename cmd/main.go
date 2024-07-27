package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/timeitel/groceries/internal/routes"
	"github.com/timeitel/groceries/internal/services"
)

func main() {
	db := services.InitDatabase()
	defer db.Close()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Renderer = services.NewTemplate("../internal/views/**/*.html")

	e.Static("/static", "../static/")

	e.GET("/", routes.Index)

	e.POST("/item/:id", routes.UpdateItem)

	e.Logger.Fatal(e.Start(":3000"))
}
