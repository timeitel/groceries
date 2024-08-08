package user

import (
	"database/sql"

	"github.com/timeitel/groceries/internal/domain/user/libsql"
	"github.com/timeitel/groceries/internal/models"
)

type Repository interface {
	GetItems() (models.Items, error)
}

func NewLibSqlRepository(db *sql.DB) *libsql.Repository {
	return &libsql.Repository{
		DB: db,
	}
}
