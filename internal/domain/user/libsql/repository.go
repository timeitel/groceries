package libsql

import (
	"database/sql"
)

type Repository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *Repository {
	return &Repository{
		db,
	}
}
