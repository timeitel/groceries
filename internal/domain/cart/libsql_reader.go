package cart

import (
	"context"

	"github.com/google/uuid"
	"github.com/timeitel/groceries/internal/domain"
)

func (r *libSqlRepo) GetItems(cartID uuid.UUID) (*Items, error) {
	dbItems, err := r.db.GetCartItems(context.Background(), cartID)
	if err != nil {
		return nil, domain.NewSQLError("Creating cart item", err)
	}

	var items Items
	for _, i := range dbItems {
		items = append(items, newItem(i))
	}

	return &items, nil
}
