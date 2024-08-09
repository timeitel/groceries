package shopper

import (
	"database/sql"
	"log"
	"os"

	"github.com/timeitel/groceries/internal/common/models"
	"github.com/timeitel/groceries/internal/domain/shopper/libsql"
)

type RepoReader interface {
	GetItems() (models.Items, error)
	GetUser() (models.User, error)
}

func NewLibSqlRepository() RepoReader {
	url := os.Getenv("DB_URL")

	db, err := sql.Open("libsql", url)
	if err != nil {
		log.Fatal("Unable to open db", url, err)
	}

	return &libsql.Repository{
		DB: db,
	}
}
