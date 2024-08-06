package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/tursodatabase/go-libsql"
)

type service struct {
	conn *sql.DB
}

func NewService() service {
	return service{
		conn: newLibSql(),
	}
}

func newLibSql() *sql.DB {
	url := os.Getenv("DB_URL")

	db, err := sql.Open("libsql", url)
	if err != nil {
		log.Fatal("Unable to open db", url, err)
	}

	return db
}
