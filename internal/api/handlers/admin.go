package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/timeitel/groceries/internal/services"
	"github.com/timeitel/groceries/internal/types"
)

func AdminGetHome(c echo.Context, s *services.Admin) error {
	p, err := s.GetProducts()
	if err != nil {
		return err
	}

	params := struct {
		Products types.Products
	}{Products: p}

	return c.Render(http.StatusOK, "admin", params)
}

type createProductParams struct {
	Name        string
	Description string
}

func AdminCreateProduct(c echo.Context, s *services.Admin) error {
	name := c.FormValue("name")
	description := c.FormValue("description")

	params := createProductParams{
		Name:        name,
		Description: description,
	}

	return c.Render(http.StatusOK, "created-product", params)
}
