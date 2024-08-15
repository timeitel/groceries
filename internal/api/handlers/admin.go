package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/timeitel/groceries/internal/services"
	"github.com/timeitel/groceries/internal/types"
)

func AdminGetHome(c echo.Context, s *services.Admin) error {
	p, err := s.GetProducts()
	if err != nil {
		return err
	}

	params := struct {
		Products types.Products
	}{Products: p}

	return c.Render(http.StatusOK, "admin", params)
}

func AdminCreateProduct(c echo.Context, s *services.Admin) error {
	name := c.FormValue("name")
	description := c.FormValue("description")

	p, _ := s.CreateProduct(name, description)
	// if err != nil {
	// 	if libsqlErr, ok := err.(*libsql.Option.Error()); ok { // Assuming libsql.Error is the custom error type
	// 		fmt.Printf("LibSQL error code: %d\n", libsqlErr.Code)
	// 		fmt.Printf("LibSQL error message: %s\n", libsqlErr.Message)
	// 	} else {
	// 		log.Fatalf("Scan failed: %v", err)
	// 	}
	//
	// 	fmt.Println("error: ", err.Error())
	//
	// 	return err
	// }

	c.Render(http.StatusOK, "product-created", p)

	return c.Render(http.StatusOK, "product-form", nil)
}
