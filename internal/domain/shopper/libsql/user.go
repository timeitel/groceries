package libsql

import "github.com/timeitel/groceries/internal/types/models"

func (r *Repository) GetUser() (models.User, error) {
	r.DB.QueryRow("")

	user := models.User{
		Id:   "asdfasd",
		Name: "cool guy",
	}

	return user, nil
}
