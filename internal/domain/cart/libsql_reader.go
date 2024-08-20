package cart

import (
	"context"

	"github.com/google/uuid"
	"github.com/timeitel/groceries/internal/domain"
)

func (r *libSqlRepo) GetItems(userID uuid.UUID) (*Items, error) {
	cart, err := r.db.GetCart(context.Background(), userID)
	if err != nil {
		return nil, err
	}

	dbItems, err := r.db.GetCartItems(context.Background(), cart.ID)
	if err != nil {
		return nil, domain.NewSQLError("Creating cart item", err)
	}

	var items Items
	for _, i := range dbItems {
		items = append(items, newItem(i))
	}

	return &items, nil
}

func (r *libSqlRepo) Get(userID uuid.UUID) (*cart, error) {
	dbCart, err := r.db.GetCart(context.Background(), userID)
	if err != nil {
		return nil, domain.NewSQLError("Getting cart", err)
	}

	cart := newCart(dbCart)

	return &cart, nil
}
