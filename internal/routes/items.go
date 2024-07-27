package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func UpdateItem(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
