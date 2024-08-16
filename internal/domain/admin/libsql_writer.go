package admin

import (
	"context"

	"github.com/timeitel/groceries/internal/infrastructure/data/db"
	"github.com/timeitel/groceries/internal/types"
)

func (r *libSqlRepository) CreateProduct(name, description string) (*db.Product, error) {
	params := db.CreateProductParams{
		Name:        name,
		Description: types.NewSqlNullString(description),
	}

	p, err := r.db.CreateProduct(context.Background(), params)
	if err != nil {
		return nil, types.NewSQLError("Admin creating product", err)
	}

	return &p, nil
}
