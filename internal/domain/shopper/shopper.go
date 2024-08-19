package shopper

import (
	"github.com/timeitel/groceries/internal/infrastructure/data/db"
)

type Shopper struct {
	UserName  string
	CartItems []db.CartItem
}
