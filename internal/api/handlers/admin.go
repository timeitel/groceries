package handlers

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/timeitel/groceries/internal/services"
	"github.com/timeitel/groceries/internal/types"
)

type formErr struct {
	Error string
}

func AdminGetHome(c echo.Context, s *services.Admin) error {
	p, err := s.GetProducts()
	if err != nil {
		return err
	}

	params := struct {
		Products types.Products
		Error    string
	}{Products: p, Error: ""}

	return c.Render(http.StatusOK, "admin", params)
}

func AdminCreateProduct(c echo.Context, s *services.Admin) error {
	name := c.FormValue("name")
	description := c.FormValue("description")
	formErr := formErr{Error: ""}

	p, err := s.CreateProduct(name, description)
	if err != nil {
		if errors.Is(err, types.ErrSQLUnique) {
			formErr.Error = "This product name already exists"
			return c.Render(http.StatusUnprocessableEntity, "product-form", formErr)
		}

		return c.Render(http.StatusInternalServerError, "product-form", formErr)
	}

	c.Render(http.StatusOK, "product-created", p)

	return c.Render(http.StatusOK, "product-form", formErr)
}
