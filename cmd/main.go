package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/timeitel/groceries/internal/database"
	"github.com/timeitel/groceries/internal/domain/user/libsql"
	services "github.com/timeitel/groceries/internal/services/list"
	"github.com/timeitel/groceries/internal/views"
	"github.com/timeitel/groceries/internal/views/home"
)

func main() {
	db := database.NewLibSql()
	repo := libsql.NewUserRepository(db)
	service := services.NewList(repo)

	e := echo.New()

	// e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Renderer = views.NewTemplate("internal/views/*.html")

	e.Static("/static", "static")

	e.GET("/", home.Index(service.GetItems()))

	e.POST("/items/:name", home.AddItem)

	e.Logger.Fatal(e.Start(":8080"))
}
