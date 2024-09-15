package catalogue

import (
	"github.com/timeitel/groceries/internal/domain/item"
)

type RepoReader interface {
	GetItems() (*item.Items, error)
	GetItem(id int64) (*item.Item, error)
	// search
}
