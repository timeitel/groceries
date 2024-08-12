package libsql

import "github.com/timeitel/groceries/internal/infrastructure/data/db"

type Repository struct {
	DB *db.Queries
}
