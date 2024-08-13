package shopper

import (
	"github.com/timeitel/groceries/internal/common/models"
	"github.com/timeitel/groceries/internal/infrastructure/data"
	"github.com/timeitel/groceries/internal/infrastructure/data/db"
)

type libSqlRepository struct {
	db *db.Queries
}

type ReadWriter interface {
	GetProducts() (models.Products, error)
	GetShopper() (shopper, error)
	AddProductToCart(productId, cartId, quantity int) error
}

func NewLibSqlRepository() ReadWriter {
	conn := data.NewLibSqlDB()

	return &libSqlRepository{
		db: conn,
	}
}
