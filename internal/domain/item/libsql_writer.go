package item

import (
	"context"

	"github.com/timeitel/groceries/internal/domain"
	"github.com/timeitel/groceries/internal/infrastructure/data/db"
)

func (r *libSqlRepo) Create(name, description string) (*Item, error) {
	params := db.CreateItemParams{
		Name:        name,
		Description: domain.NewSqlNullString(description),
	}

	i, err := r.db.CreateItem(context.Background(), params)
	if err != nil {
		return nil, domain.NewSQLError("Creating item", err)
	}

	item := New(i)

	return &item, nil
}

func (r *libSqlRepo) Delete(id int64) error {
	err := r.db.DeleteItem(context.Background(), id)
	if err != nil {
		return domain.NewSQLError("Deleting product", err)
	}

	return nil
}

func (r *libSqlRepo) Update(id int64, name, description string) (*Item, error) {
	params := db.UpdateItemParams{
		ID:          id,
		Name:        name,
		Description: domain.NewSqlNullString(description),
	}

	p, err := r.db.UpdateItem(context.Background(), params)
	if err != nil {
		return nil, domain.NewSQLError("Updating item", err)
	}

	item := New(p)

	return &item, nil
}
