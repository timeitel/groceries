package services

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/timeitel/groceries/internal/domain/cart"
	"github.com/timeitel/groceries/internal/domain/catalogue"
	"github.com/timeitel/groceries/internal/domain/item"
	"github.com/timeitel/groceries/internal/domain/user"
	_ "github.com/tursodatabase/go-libsql"
)

type Shopper struct {
	catalogue catalogue.RepoReader
	cart      cart.RepoWriter
	user      user.RepoReader
}

// TODO: inject repos
func NewShopper() Shopper {
	return Shopper{
		catalogue: catalogue.NewLibSqlRepo(),
		cart:      cart.NewLibSqlRepo(),
		user:      user.NewLibSqlRepo(),
	}
}

func (s *Shopper) AddItemToCart(itemId uuid.UUID, quantity int) (*cart.Item, error) {
	cart, err := s.cart.Get()
	if err != nil {
		return nil, err
	}

	item, err := s.cart.AddItem(itemId, cart.ID, quantity)
	fmt.Println(item)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (s *Shopper) GetItems() (*item.Items, error) {
	items, err := s.catalogue.GetItems()
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (s *Shopper) GetItem(id uuid.UUID) (*item.Item, error) {
	item, err := s.catalogue.GetItem(id)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (s *Shopper) GetCartItems() (*cart.Items, error) {
	cart, err := s.cart.Get()
	if err != nil {
		return nil, err
	}

	return &cart.Items, nil
}
