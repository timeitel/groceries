package libsql

import (
	"fmt"
	"log"

	"github.com/timeitel/groceries/internal/common/models"
)

func (r *Repository) GetUser() (models.User, error) {
	query := "SELECT id, name, is_admin FROM users LIMIT 1"
	row := r.DB.QueryRow(query)

	var (
		id         string
		name       string
		isAdminInt int
	)

	err := row.Scan(&id, &name, &isAdminInt)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("USER: %s, %s, %d \n", id, name, isAdminInt)

	user := models.User{
		Id:      id,
		Name:    name,
		IsAdmin: isAdminInt == 1,
	}

	return user, nil
}
