package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/timeitel/groceries/internal/routes"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/static", "../static/")

	http.HandleFunc("/item", routes.PostItemHandler)

	e.Logger.Fatal(e.Start(":3000"))
}
