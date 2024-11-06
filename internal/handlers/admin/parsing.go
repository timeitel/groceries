package admin

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

func getIDFromPath(c echo.Context) (*int64, error) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return nil, err
	}

	return &id, nil
}
