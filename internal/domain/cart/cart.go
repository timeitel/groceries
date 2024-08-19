package cart

import (
	"github.com/google/uuid"
	"github.com/timeitel/groceries/internal/infrastructure/data/db"
)

type cart struct {
	ID    uuid.UUID
	Name  string
	Items Items
}

type Items []Item
type Item struct {
	Name     string
	Quantity int
}

func newCartItem(ci db.CartItem) Item {
	return Item{
		Name:     "asd",
		Quantity: int(ci.Quantity.Int64),
	}
}

func newCart(c db.Cart) cart {
	return cart{
		Name: c.Name.String,
	}
}
