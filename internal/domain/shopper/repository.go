package shopper

import (
	"github.com/timeitel/groceries/internal/common/models"
	"github.com/timeitel/groceries/internal/domain/shopper/libsql"
	"github.com/timeitel/groceries/internal/infrastructure/data"
	"github.com/timeitel/groceries/internal/infrastructure/data/db"
)

type RepoReadWriter interface {
	GetProducts() (models.Products, error)
	GetUser() (db.User, error)
	AddProduct(id string) (db.Product, error)
}

func NewLibSqlRepository() RepoReadWriter {
	DB := data.NewLibSqlQueries()

	return &libsql.Repository{
		DB,
	}
}
