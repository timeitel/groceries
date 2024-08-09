package shopper

import "github.com/timeitel/groceries/internal/common/models"

type Shopper struct {
	User       models.User // root
	activeCart models.Cart
}
