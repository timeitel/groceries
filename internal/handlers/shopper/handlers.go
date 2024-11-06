package shopper

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/timeitel/groceries/internal/services"
	"github.com/timeitel/groceries/internal/web/views/components"
	"github.com/timeitel/groceries/internal/web/views/pages"
)

func addItem(c echo.Context, service *services.Shopper) error {
	id, err := getIDFromPath(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid ID")
	}

	item, err := service.AddItemToCart(*id, 1)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Unable to add item, please try again")
	}

	return render(c, components.CartItemCard(*item))
}

func getHome(c echo.Context, service *services.Shopper) error {
	items, err := service.GetItems()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	cartItems, err := service.GetCartItems()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	vm := pages.ShopperHomeViewModel{
		Items:     *items,
		UserName:  "cool guy",
		CartItems: *cartItems,
	}

	return render(c, pages.ShopperHome(vm))
}

func getItem(c echo.Context, s *services.Shopper) error {
	id, err := getIDFromPath(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid id")
	}

	i, err := s.GetItem(*id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "No item with that id")
	}

	return render(c, pages.ShopperItem(*i))
}
