package cart

import (
	"github.com/google/uuid"
)

type RepoWriter interface {
	GetItems(cartID uuid.UUID) (*Items, error)
	AddItem(cartID, itemID uuid.UUID, quantity int) (*Item, error)
	RemoveItem(cartID, itemID uuid.UUID) error
	UpdateItemQuantity(quantity int) (*int, error)
}
