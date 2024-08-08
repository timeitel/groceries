package libsql

import (
	"database/sql"
)

type Repository struct {
	DB *sql.DB
}

func (r *Repository) GetUserName() (string, error) {
	return "cool guy", nil
}
