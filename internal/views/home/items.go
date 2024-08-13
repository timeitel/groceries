package home

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var added = []string{}

type addProductFn func(id string) error

func AddProductToCart(addItem addProductFn) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		err := addItem(id)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest)
		}

		return c.Render(http.StatusOK, "added", added)
	}
}
