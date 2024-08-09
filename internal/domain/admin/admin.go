package admin

import "github.com/timeitel/groceries/internal/common/models"

type Admin struct {
	User     models.User // root
	shoppers int
	// shopper metrics, activity
}
