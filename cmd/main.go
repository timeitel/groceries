package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/timeitel/groceries/internal/routes"
	"github.com/timeitel/groceries/internal/services"
	"github.com/timeitel/groceries/internal/services/db"
)

func main() {
	DB := db.Init()
	db.QueryUsers(DB)
	defer DB.Close()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Renderer = services.NewTemplate("internal/views/*.html")

	e.Static("/static", "static")

	e.GET("/", routes.Index)

	e.POST("/items/:name", routes.AddItem)

	e.Logger.Fatal(e.Start(":8080"))
}
