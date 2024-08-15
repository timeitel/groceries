package api

import (
	"github.com/labstack/echo/v4"
)

func Middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// token := c.Request().Header.Get("Authorization")

		// TODO: implement
		// if token == "" {
		// 	return c.JSON(http.StatusUnauthorized, map[string]string{
		// 		"error": "Unauthorized",
		// 	})
		// }
		//
		// if token != "valid-token" { // Replace with your actual validation logic
		// 	return c.JSON(http.StatusUnauthorized, map[string]string{
		// 		"error": "Invalid token",
		// 	})
		// }

		return next(c)
	}
}
