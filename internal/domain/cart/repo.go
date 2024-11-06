package cart

import (
	"github.com/google/uuid"
	"github.com/timeitel/groceries/internal/infrastructure/data"
	"github.com/timeitel/groceries/internal/infrastructure/data/db"
)

type RepoReadWriter interface {
	Get(userID uuid.UUID) (*cart, error)
	GetItems(userID uuid.UUID) (*Items, error)
	AddItem(cartID uuid.UUID, itemID int64, quantity int) (*Item, error)
	RemoveItem(cartItemID uuid.UUID) error
	UpdateItemQuantity(quantity int) (*int, error)
}

type libSqlRepo struct {
	db *db.Queries
}

func NewLibSqlRepo() RepoReadWriter {
	conn := data.NewLibSqlDB()

	return &libSqlRepo{
		db: conn,
	}
}
