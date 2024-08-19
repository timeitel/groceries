package handlers

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/timeitel/groceries/internal/domain"
	"github.com/timeitel/groceries/internal/domain/item"
	"github.com/timeitel/groceries/internal/services"
	"github.com/timeitel/groceries/internal/web/views/components"
	"github.com/timeitel/groceries/internal/web/views/pages"
)

type adminPageData struct {
	Items item.Items
	Error string
}

type formErr struct {
	Error string
}

func AdminGetProduct(c echo.Context, s *services.Admin) error {
	id, err := getIDFromPath(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	p, err := s.GetProduct(*id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound)
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
		if errors.Is(err, domain.ErrSQLUnique) {
			formErr.Error = "This product name already exists"
			return render(c, components.AddProductForm())
		}

		return render(c, components.AddProductForm())
	}

	render(c, components.CreatedProductCard(*p))

	return render(c, components.AddProductForm())
}

func AdminDeleteProduct(c echo.Context, s *services.Admin) error {
	id, err := getIDFromPath(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	err = s.DeleteProduct(*id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.Redirect(http.StatusSeeOther, "/admin")
}

func AdminUpdateProduct(c echo.Context, s *services.Admin) error {
	id, err := getIDFromPath(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	name := c.FormValue("name")
	description := c.FormValue("description")

	p, err := s.UpdateProduct(*id, name, description)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return render(c, components.UpdatedProduct(*p))
}
