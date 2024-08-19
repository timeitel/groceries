package cart

import (
	"github.com/timeitel/groceries/internal/infrastructure/data"
	"github.com/timeitel/groceries/internal/infrastructure/data/db"
)

type libSqlRepo struct {
	db *db.Queries
}

func NewLibSqlRepo() RepoWriter {
	conn := data.NewLibSqlDB()

	return &libSqlRepo{
		db: conn,
	}
}
