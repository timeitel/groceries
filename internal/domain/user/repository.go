package user

import (
	"database/sql"

	"github.com/timeitel/groceries/internal/domain/user/libsql"
	"github.com/timeitel/groceries/internal/models"
)

type RepoReader interface {
	GetItems() (models.Items, error)
	GetUserName() (string, error)
}

func NewLibSqlRepository(db *sql.DB) RepoReader {
	return &libsql.Repository{
		DB: db,
	}
}
