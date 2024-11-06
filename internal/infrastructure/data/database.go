package data

import (
	"database/sql"
	"log"
	"os"

	"github.com/timeitel/groceries/internal/infrastructure/data/db"
	_ "github.com/tursodatabase/go-libsql"
)

func NewLibSqlDB() *db.Queries {
	url := os.Getenv("DB_URL")

	conn, err := sql.Open("libsql", url)
	if err != nil {
		log.Fatal("Unable to open db", url, err)
	}

	return db.New(conn)
}
