package admin

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

func getItem(c echo.Context, s *services.Admin) error {
	id, err := getIDFromPath(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	p, err := s.GetItem(*id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	return render(c, pages.AdminItem(*p))
}

func getHome(c echo.Context, s *services.Admin) error {
	p, err := s.GetItems()
	if err != nil {
		return err
	}

	return render(c, pages.AdminHome(p))
}

func createItem(c echo.Context, s *services.Admin) error {
	name := c.FormValue("name")
	description := c.FormValue("description")
	formErr := formErr{Error: ""}

	p, err := s.CreateItem(name, description)
	if err != nil {
		if errors.Is(err, domain.ErrSQLUnique) {
			formErr.Error = "This product name already exists"
			return render(c, components.AddItemForm())
		}

		return render(c, components.AddItemForm())
	}

	render(c, components.CreatedItemCard(*p))

	return render(c, components.AddItemForm())
}

func deleteItem(c echo.Context, s *services.Admin) error {
	id, err := getIDFromPath(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	err = s.DeleteItem(*id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.Redirect(http.StatusSeeOther, "/admin")
}

func updateItem(c echo.Context, s *services.Admin) error {
	id, err := getIDFromPath(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	name := c.FormValue("name")
	description := c.FormValue("description")

	p, err := s.UpdateItem(*id, name, description)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return render(c, components.UpdatedItem(*p))
}
