package home

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/timeitel/groceries/internal/common/models"
)

type Data struct {
	Items models.Items
	Name  string
}

type getDataFn func(jwt string) Data

func Index(getData getDataFn) echo.HandlerFunc {
	return func(c echo.Context) error {
		// jwt := c.Param("jwt")
		data := getData("jwt")

		return c.Render(http.StatusOK, "index", data)
	}
}
