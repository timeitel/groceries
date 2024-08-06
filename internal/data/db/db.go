package db

import (
	"database/sql"
	"log"
	"os"
)

func NewLibSql() *sql.DB {
	url := os.Getenv("DB_URL")

	db, err := sql.Open("libsql", url)
	if err != nil {
		log.Fatal("Unable to open db", url, err)
	}

	return db
}
