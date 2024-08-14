package home

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/timeitel/groceries/internal/services"
)

func AddProductToCartHandler(c echo.Context, service *services.User) error {
	// id := c.Param("id")

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
