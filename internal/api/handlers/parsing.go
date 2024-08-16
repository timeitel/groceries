package handlers

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

func getIdFromPath(c echo.Context) (*int64, error) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return nil, err
	}

	return &id, nil
}
