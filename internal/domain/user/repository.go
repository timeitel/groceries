package user

import (
	"github.com/timeitel/groceries/internal/models"
)

type Repository interface {
	GetItems() (models.Items, error)
}
