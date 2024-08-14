package services

import (
	"fmt"

	"github.com/timeitel/groceries/internal/common/types"
	"github.com/timeitel/groceries/internal/domain/shopper"
	_ "github.com/tursodatabase/go-libsql"
)

type User struct {
	repo shopper.RepoReadWriter
}

func NewUser(r shopper.RepoReadWriter) User {
	return User{
		repo: r,
	}
}

func (s *User) AddProductToCart(productID int, quantity int) error {
	user, err := s.repo.GetShopper()
	if err != nil {
		fmt.Println(err)
		return err
	}

	if err = s.repo.AddProductToCart(productID, 1); err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println(user)

	return nil
}

func (s *User) GetProducts() (types.Products, error) {
	products, err := s.repo.GetProducts()
	if err != nil {
		fmt.Println("getting products", err)
		return nil, err
	}

	return products, nil
}
