package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var options = []Item{{Name: "strawberry", Icon: "🍓"}, {Name: "banana", Icon: "🍌"}, {Icon: "🍎", Name: "apple"}, {Name: "nectarine", Icon: "🍑"}}

type Data struct {
	Items []Item
	Name  string
}

func Index(c echo.Context) error {
	d := Data{Items: options, Name: "Michael Connor"}
	return c.Render(http.StatusOK, "index", d)
}
