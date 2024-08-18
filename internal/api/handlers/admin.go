package handlers

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/timeitel/groceries/internal/services"
	"github.com/timeitel/groceries/internal/types"
	"github.com/timeitel/groceries/internal/web/views/components"
	"github.com/timeitel/groceries/internal/web/views/pages"
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

	return render(c, pages.AdminProduct(*p))
}

func AdminGetHome(c echo.Context, s *services.Admin) error {
	p, err := s.GetProducts()
	if err != nil {
		return err
	}

	return render(c, pages.AdminHome(p))
}

func AdminCreateProduct(c echo.Context, s *services.Admin) error {
	name := c.FormValue("name")
	description := c.FormValue("description")
	formErr := formErr{Error: ""}

	p, err := s.CreateProduct(name, description)
	if err != nil {
		if errors.Is(err, types.ErrSQLUnique) {
			formErr.Error = "This product name already exists"
			return render(c, components.AddProductForm())
		}

		return render(c, components.AddProductForm())
	}

	render(c, components.CreatedProductCard(*p))

	return render(c, components.AddProductForm())
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

	return c.Redirect(http.StatusSeeOther, "/admin")
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

	return render(c, components.UpdatedProduct(*p))
}
