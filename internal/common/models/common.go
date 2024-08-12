package models

import (
	"database/sql"

	"github.com/timeitel/groceries/internal/infrastructure/data/db"
)

type Products []db.Product
type CartItems []db.CartItem
type Carts []db.Cart

type SqlInt = sql.NullInt64
