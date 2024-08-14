package home

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/timeitel/groceries/internal/services"
)

func HomeHandler(c echo.Context, service *services.User) error {
	products, err := service.GetProducts()
	if err != nil {
		return c.Render(http.StatusInternalServerError, "index", products)
	}

	return c.Render(http.StatusOK, "index", products)
}
