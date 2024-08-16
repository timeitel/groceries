package shopper

import (
	"github.com/timeitel/groceries/internal/infrastructure/data"
	"github.com/timeitel/groceries/internal/infrastructure/data/db"
	"github.com/timeitel/groceries/internal/types"
)

type libSqlRepository struct {
	db *db.Queries
}

type RepoReadWriter interface {
	GetProduct(id int64) (*db.Product, error)
	GetProducts() (types.Products, error)
	GetShopper() (shopper, error)
	AddProductToCart(productId, quantity int) error
}

func NewLibSqlRepository() RepoReadWriter {
	conn := data.NewLibSqlDB()

	return &libSqlRepository{
		db: conn,
	}
}
