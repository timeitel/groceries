package home

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var added = []string{}

func AddItem(c echo.Context) error {
	name := c.Param("name")
	added = append(added, name)
	return c.Render(http.StatusOK, "added", added)
}
