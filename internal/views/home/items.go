package home

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var added = []string{}

type addProductFn func(productID int, quantity int) error

func AddProductToCart(addItem addProductFn) echo.HandlerFunc {
	return func(c echo.Context) error {
		// id := c.Param("id")

		if err := addItem(1, 1); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest)
		}

		return c.Render(http.StatusOK, "added", added)
	}
}
