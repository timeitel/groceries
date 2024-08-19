package user

import (
	"github.com/google/uuid"
	"github.com/timeitel/groceries/internal/infrastructure/data/db"
)

type user struct {
	Name       string
	ActiveCart uuid.UUID
}

func New(u db.User) user {
	return user{
		Name:       u.Name,
		ActiveCart: u.ActiveCartID,
	}
}
