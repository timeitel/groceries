package home

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/timeitel/groceries/internal/models"
)

var options = []models.Item{{Name: "strawberry", Icon: "🍓"}, {Name: "banana", Icon: "🍌"}, {Icon: "🍎", Name: "apple"}, {Name: "nectarine", Icon: "🍑"}}

type Data struct {
	Items []models.Item
	Name  string
}

func Index(items models.Items) func(c echo.Context) error {
	return func(c echo.Context) error {
		d := Data{Items: items, Name: "Michael Connor"}
		return c.Render(http.StatusOK, "index", d)
	}
}
