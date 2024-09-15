package cart

import (
	"context"

	"github.com/google/uuid"
	"github.com/timeitel/groceries/internal/domain"
	"github.com/timeitel/groceries/internal/infrastructure/data/db"
)

func (r *libSqlRepo) AddItem(cartId uuid.UUID, itemId int64, quantity int) (*Item, error) {
	params := db.CreateCartItemParams{
		ItemID: itemId, CartID: cartId, Quantity: domain.NewSqlNullInt(quantity),
	}

	ci, err := r.db.CreateCartItem(context.Background(), params)
	if err != nil {
		return nil, domain.NewSQLError("Creating cart item", err)
	}

	cartItem := newItem(ci)

	return &cartItem, nil
}

func (r *libSqlRepo) RemoveItem(cartItemId uuid.UUID) error {
	err := r.db.DeleteCartItem(context.Background(), cartItemId)
	if err != nil {
		return domain.NewSQLError("deleting cart item", err)
	}

	return nil
}

func (r *libSqlRepo) UpdateItemQuantity(quantity int) (*int, error) {
	params := db.UpdateCartItemQuantityParams{
		Quantity: domain.NewSqlNullInt(quantity),
	}

	p, err := r.db.UpdateCartItemQuantity(context.Background(), params)
	if err != nil {
		return nil, domain.NewSQLError("Updating item quantity", err)
	}

	res := int(p.Quantity.Int64)

	return &res, nil
}
