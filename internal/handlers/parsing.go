package handlers

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func getIDFromPath(c echo.Context) (*uuid.UUID, error) {
	idParam := c.Param("id")
	uuid, err := uuid.Parse(idParam)
	if err != nil {
		return nil, err
	}

	return &uuid, nil
}
