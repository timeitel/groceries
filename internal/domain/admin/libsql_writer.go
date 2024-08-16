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

func (r *libSqlRepository) DeleteProduct(id int64) error {
	err := r.db.DeleteProduct(context.Background(), id)
	if err != nil {
		return types.NewSQLError("admin deleting product", err)
	}

	return nil
}

func (r *libSqlRepository) UpdateProduct(id int64, name, description string) error {
	params := db.UpdateProductParams{
		Name:        name,
		Description: types.NewSqlNullString(description),
	}

	err := r.db.UpdateProduct(context.Background(), params)
	if err != nil {
		return types.NewSQLError("Admin creating product", err)
	}

	return nil
}
