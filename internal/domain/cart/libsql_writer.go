package cart

import (
	"context"

	"github.com/google/uuid"
	"github.com/timeitel/groceries/internal/domain"
	"github.com/timeitel/groceries/internal/infrastructure/data/db"
)

func (r *libSqlRepo) AddItem(itemId, cartId uuid.UUID, quantity int) (*Item, error) {
	params := db.CreateCartItemParams{
		ItemID: itemId, CartID: cartId, Quantity: domain.NewSqlNullInt(quantity),
	}

	ci, err := r.db.CreateCartItem(context.Background(), params)
	if err != nil {
		return nil, domain.NewSQLError("Creating cart item", err)
	}

	cartItem := newCartItem(ci)

	return &cartItem, nil
}

func (r *libSqlRepo) RemoveItem(cartId, itemId uuid.UUID) error {
	params := db.DeleteCartItemParams{
		CartID: cartId, ItemID: itemId,
	}

	err := r.db.DeleteCartItem(context.Background(), params)
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
