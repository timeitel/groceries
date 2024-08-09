package home

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var added = []string{}

type addItemFn func(id string) error

func AddItem(addItem addItemFn) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		err := addItem(id)
		if err != nil {
			return c.Render(http.StatusBadRequest, "", nil)
		}

		return c.Render(http.StatusOK, "added", added)
	}
}
