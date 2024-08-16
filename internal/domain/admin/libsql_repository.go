package admin

import (
	"github.com/timeitel/groceries/internal/infrastructure/data"
	"github.com/timeitel/groceries/internal/infrastructure/data/db"
)

type libSqlRepository struct {
	db *db.Queries
}

type RepoWriter interface {
	CreateProduct(name, description string) (*db.Product, error)
	DeleteProduct(id int64) error
	UpdateProduct(id int64, name, description string) (*db.Product, error)
}

func NewLibSqlRepository() RepoWriter {
	conn := data.NewLibSqlDB()

	return &libSqlRepository{
		db: conn,
	}
}
