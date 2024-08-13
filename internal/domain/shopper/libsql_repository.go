package shopper

import (
	"github.com/timeitel/groceries/internal/common/types"
	"github.com/timeitel/groceries/internal/infrastructure/data"
	"github.com/timeitel/groceries/internal/infrastructure/data/db"
)

type libSqlRepository struct {
	db *db.Queries
}

type ReadWriter interface {
	GetProducts() (types.Products, error)
	GetShopper() (shopper, error)
	AddProductToCart(productId, quantity int) error
}

func NewLibSqlRepository() ReadWriter {
	conn := data.NewLibSqlDB()

	return &libSqlRepository{
		db: conn,
	}
}
