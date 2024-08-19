package cart

import (
	"github.com/google/uuid"
)

type RepoWriter interface {
	Get() (*cart, error)
	AddItem(itemId, cartId uuid.UUID, quantity int) (*Item, error)
	RemoveItem(cartId, itemId uuid.UUID) error
	UpdateItemQuantity(quantity int) (*int, error)
}
