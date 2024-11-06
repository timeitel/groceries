package catalogue

import (
	"context"

	"github.com/timeitel/groceries/internal/domain/item"
)

func (r *libSqlRepo) GetItem(id int64) (*item.Item, error) {
	i, err := r.db.GetItem(context.Background(), id)
	if err != nil {
		return nil, err
	}

	item := item.New(i)

	return &item, nil
}

func (r *libSqlRepo) GetItems() (*item.Items, error) {
	dbItems, err := r.db.GetItems(context.Background())
	if err != nil {
		return nil, err
	}

	var items item.Items
	for _, i := range dbItems {
		item := item.New(i)
		items = append(items, item)
	}

	return &items, nil
}
