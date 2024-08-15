package handlers

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/timeitel/groceries/internal/services"
	"github.com/timeitel/groceries/internal/types"
)

func ShopperAddProduct(c echo.Context, service *services.User) error {
	if err := service.AddProductToCart(1, 1); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	// items, err := service.GetHomeData()
	// if err != nil {
	// 	fmt.Println("Getting cart items", err)
	// 	return echo.NewHTTPError(http.StatusInternalServerError)
	// }
	//

	return c.Render(http.StatusOK, "added", nil)
}

type data struct {
	Products types.Products
	Name     string
}

func ShopperHome(c echo.Context, service *services.User) error {
	products, err := service.GetProducts()
	if err != nil {
		log.Fatalln("getting products")
	}

	data := data{
		Products: products,
		Name:     "Cool guy",
	}

	return c.Render(http.StatusOK, "index", data)
}
