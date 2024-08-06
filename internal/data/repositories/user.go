package repositories

import (
	"database/sql"

	"github.com/timeitel/groceries/internal/data/models"
)

type getter interface {
	AddItem(item models.Item, quantity int) error
}

type User struct {
	db *sql.DB
}

func NewUser(db *sql.DB) User {
	return User{
		db,
	}
}

func (r *User) AddItem(userId string, itemId string) error {
	_, err := r.db.Exec("INSERT INTO shopping_lists (userId, itemId) VALUES (?, ?)", userId, itemId)

	return err
}
