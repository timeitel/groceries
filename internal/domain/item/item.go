package item

import (
	"github.com/google/uuid"
	"github.com/timeitel/groceries/internal/infrastructure/data/db"
)

type Items []Item

type Item struct {
	ID          uuid.UUID
	Name        string
	Description string
}

func New(i db.Item) Item {
	return Item{
		ID:          i.ID,
		Name:        i.Name,
		Description: i.Description.String,
	}
}
