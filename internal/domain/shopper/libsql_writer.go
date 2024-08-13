package shopper

import (
	"context"
	"fmt"

	"github.com/timeitel/groceries/internal/common/types"
	"github.com/timeitel/groceries/internal/infrastructure/data/db"
)

func (r *libSqlRepository) AddProductToCart(productId, quantity int) error {
	res, err := r.db.GetShopper(context.Background())
	if err != nil {
		fmt.Println("Shopper", err)
		return err
	}

	params := db.AddCartItemParams{
		CartID:    res.User.ActiveCartID,
		ProductID: types.NewSqlInt(productId),
		Quantity:  types.NewSqlInt(quantity),
	}

	if _, err = r.db.AddCartItem(context.Background(), params); err != nil {
		fmt.Println("Unable to add cart item", err)
		return err
	}

	return nil
}
