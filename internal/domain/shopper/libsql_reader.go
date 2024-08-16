package shopper

import (
	"context"
	"log"

	"github.com/timeitel/groceries/internal/infrastructure/data/db"
	"github.com/timeitel/groceries/internal/types"
)

func (r *libSqlRepository) GetShopper() (shopper, error) {
	// TODO:
	u, err := r.db.GetShopper(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	return u, nil
}

func (r *libSqlRepository) GetProducts() (types.Products, error) {
	p, err := r.db.GetProducts(context.Background())

	if err != nil {
		return nil, err
	}

	return p, nil
}

func (r *libSqlRepository) GetProduct(id int64) (*db.Product, error) {
	p, err := r.db.GetProduct(context.Background(), id)
	if err != nil {
		return nil, types.ErrSQLNotFound
	}

	return &p, nil
}
