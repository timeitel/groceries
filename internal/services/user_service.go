package services

import (
	"fmt"

	"github.com/timeitel/groceries/internal/domain/shopper"
	"github.com/timeitel/groceries/internal/infrastructure/data/db"
	"github.com/timeitel/groceries/internal/types"
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

func (s *User) AddProductToCart(productID int, quantity int) (*db.Product, error) {
	user, err := s.repo.GetShopper()
	fmt.Println(user)
	if err != nil {
		return nil, err
	}

	if err = s.repo.AddProductToCart(productID, 1); err != nil {
		return nil, err
	}

	return nil, err
}

func (s *User) GetProducts() (types.Products, error) {
	products, err := s.repo.GetProducts()
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (s *User) GetProduct(id int64) (*db.Product, error) {
	p, err := s.repo.GetProduct(id)
	if err != nil {
		return nil, err
	}

	return p, nil
}
