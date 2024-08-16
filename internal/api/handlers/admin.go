package handlers

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/timeitel/groceries/internal/services"
	"github.com/timeitel/groceries/internal/types"
	"github.com/timeitel/groceries/internal/web/views"
)

type adminPageData struct {
	Products types.Products
	Error    string
}

type formErr struct {
	Error string
}

func AdminGetProduct(c echo.Context, s *services.Admin) error {
	id, err := getIdFromPath(c)
	if err != nil {
		return c.Render(http.StatusOK, "admin-product-not-found", nil)
	}

	p, err := s.GetProduct(*id)
	if err != nil {
		return c.Render(http.StatusOK, "admin-product-not-found", nil)
	}

	return render(c, views.AdminProductPage(*p))
}

func AdminGetHome(c echo.Context, s *services.Admin) error {
	p, err := s.GetProducts()
	if err != nil {
		return err
	}

	data := adminPageData{Products: p, Error: ""}

	return c.Render(http.StatusOK, "admin", data)
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

func AdminDeleteProduct(c echo.Context, s *services.Admin) error {
	id, err := getIdFromPath(c)
	if err != nil {
		return c.Render(http.StatusBadRequest, "index", nil)
	}

	err = s.DeleteProduct(*id)
	if err != nil {
		return c.Render(http.StatusInternalServerError, "index", nil)
	}

	return render(c, views.DeletedProduct())
}

func AdminUpdateProduct(c echo.Context, s *services.Admin) error {
	id, err := getIdFromPath(c)
	if err != nil {
		return c.Render(http.StatusBadRequest, "index", nil)
	}

	name := c.FormValue("name")
	description := c.FormValue("description")

	p, err := s.UpdateProduct(*id, name, description)
	if err != nil {
		return c.Render(http.StatusInternalServerError, "index", nil)
	}

	return render(c, views.UpdatedProduct(*p))
}
