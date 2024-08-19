package user

import "github.com/timeitel/groceries/internal/infrastructure/data/db"

type user struct {
	name    string
	isAdmin bool
}

func New(u db.User) user {
	return user{
		name:    u.Name,
		isAdmin: u.IsAdmin.Valid,
	}
}
