package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/timeitel/groceries/internal/handlers"
	"github.com/timeitel/groceries/internal/handlers/admin"
	"github.com/timeitel/groceries/internal/handlers/shopper"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(handlers.Middleware)
	e.Static("/static", "internal/web/static")

	shopper.Routes(e)
	admin.Routes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
