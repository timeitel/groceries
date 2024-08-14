package home

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/timeitel/groceries/internal/common/types"
	"github.com/timeitel/groceries/internal/services"
)

type data struct {
	Products types.Products
	IsAdmin  bool
	Name     string
}

func HomeHandler(c echo.Context, service *services.User) error {
	products, err := service.GetProducts()
	if err != nil {
		log.Fatalln("getting products")
	}

	data := data{
		Products: products,
		IsAdmin:  false,
		Name:     "Cool guy",
	}

	return c.Render(http.StatusOK, "index", data)
}
