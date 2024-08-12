package libsql

import (
	"context"
	"fmt"
	"log"

	"github.com/timeitel/groceries/internal/common/models"
	"github.com/timeitel/groceries/internal/infrastructure/data/db"
)

func (r *Repository) GetUser() (db.User, error) {
	// TODO:
	u, err := r.GetUser()
	if err != nil {
		log.Fatal(err)
	}

	return u, nil
}

func (r *Repository) GetProducts() (models.Products, error) {
	p, err := r.DB.GetProducts(context.Background())

	if err != nil {
		fmt.Println("Unable to get products", err)
		return nil, err
	}

	return p, nil
}

func (r *Repository) AddProductToCart(productId, cartId, quantity models.SqlInt) error {
	params := db.AddCartItemParams{
		CartID: cartId, ProductID: productId, Quantity: quantity,
	}

	_, err := r.DB.AddCartItem(context.Background(), params)

	if err != nil {
		fmt.Println("Unable to add product to cart", err)
		return err
	}

	return nil
}
