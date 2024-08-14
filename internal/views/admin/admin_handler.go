package admin

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Handler(c echo.Context) error {
	// id := c.Param("id")

	// items, err := service.GetHomeData()
	// if err != nil {
	// 	fmt.Println("Getting cart items", err)
	// 	return echo.NewHTTPError(http.StatusInternalServerError)
	// }
	//

	return c.Render(http.StatusOK, "admin-view", nil)
}
