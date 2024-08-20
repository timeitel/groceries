package user

import (
	"github.com/google/uuid"
	"github.com/timeitel/groceries/internal/infrastructure/data/db"
)

type user struct {
	ID   uuid.UUID
	Name string
}

func New(u db.User) user {
	return user{
		ID:   u.ID,
		Name: u.Name,
	}
}
