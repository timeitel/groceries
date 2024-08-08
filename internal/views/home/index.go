package home

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/timeitel/groceries/internal/models"
)

type Data struct {
	Items []models.Item
	Name  string
}

func Index(data Data) func(c echo.Context) error {
	return func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", data)
	}
}
