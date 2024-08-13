package types

import (
	"database/sql"

	"github.com/timeitel/groceries/internal/infrastructure/data/db"
)

type Products []db.Product
type CartItems []db.CartItem
type Carts []db.Cart

type SqlInt = sql.NullInt64

func NewSqlInt(value int) sql.NullInt64 {
	return sql.NullInt64{
		Int64: int64(value),
		Valid: true,
	}
}
