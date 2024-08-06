package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/timeitel/groceries/internal/data/db"
	"github.com/timeitel/groceries/internal/data/repositories"
	"github.com/timeitel/groceries/internal/data/services"
	"github.com/timeitel/groceries/internal/routes"
	"github.com/timeitel/groceries/internal/views"
)

func main() {
	DB := db.NewLibSql()
	repo := repositories.NewUser(DB)
	service := services.NewUser(repo)

	e := echo.New()

	// e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Renderer = views.NewTemplate("internal/views/*.html")

	e.Static("/static", "static")

	e.GET("/", routes.Index)

	e.POST("/items/:name", routes.AddItem)

	e.Logger.Fatal(e.Start(":8080"))
}
