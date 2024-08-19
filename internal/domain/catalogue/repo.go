package catalogue

import (
	"github.com/google/uuid"
	"github.com/timeitel/groceries/internal/domain/item"
)

type RepoReader interface {
	GetItems() (*item.Items, error)
	GetItem(id uuid.UUID) (*item.Item, error)
	// search
}
