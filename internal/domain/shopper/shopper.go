package shopper

import "github.com/timeitel/groceries/internal/types/models"

type Aggregate struct {
	User       models.User // root
	activeCart models.Cart
}
