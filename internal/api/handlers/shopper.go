package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/timeitel/groceries/internal/domain/shopper"
	"github.com/timeitel/groceries/internal/infrastructure/data/db"
	"github.com/timeitel/groceries/internal/services"
	"github.com/timeitel/groceries/internal/types"
	"github.com/timeitel/groceries/internal/web/views/components"
	"github.com/timeitel/groceries/internal/web/views/pages"
)

func ShopperAddProduct(c echo.Context, service *services.User) error {
	p, err := service.AddProductToCart(1, 1)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	return render(c, components.ProductCard(*p))
}

func ShopperGetHome(c echo.Context, service *services.User) error {
	products, err := service.GetProducts()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	vm := pages.ShopperHomeViewModel{
		Products: products,
		Shopper: shopper.Shopper{
			UserName: "cool guy",
			CartItems: []db.CartItem{
				{ProductID: types.NewSqlNullInt(1), CartID: types.NewSqlNullInt(1), Quantity: types.NewSqlNullInt(1)},
			},
		},
	}

	return render(c, pages.ShopperHome(vm))
}

func ShopperGetProduct(c echo.Context, s *services.User) error {
	id, err := getIdFromPath(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	p, err := s.GetProduct(*id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	return render(c, pages.ShopperProduct(*p))
}
