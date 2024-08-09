package libsql

import (
	"fmt"
	"log"

	"github.com/timeitel/groceries/internal/common/models"
)

func (r *Repository) GetUser() (models.User, error) {
	query := "SELECT id, name, is_admin FROM users LIMIT 1"

	row := r.DB.QueryRow(query)

	var id string
	var name string
	var isAdmin bool

	err := row.Scan(&id, &name, &isAdmin)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("USER: %s, %s %b \n", id, name, isAdmin)

	user := models.User{
		Id:      id,
		Name:    name,
		IsAdmin: isAdmin,
	}

	return user, nil
}
