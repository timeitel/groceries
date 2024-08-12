package shopper

import (
	"github.com/timeitel/groceries/internal/infrastructure/data/db"
)

type Shopper struct {
	User       db.User
	ActiveCart db.Cart
}
