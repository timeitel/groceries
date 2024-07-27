package services

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Item struct {
	Id   int `json:"id"`
	Name int `json:"name"`
	Icon int `json:"icon"`
}

func InitDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", "../../main.db")

	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
    CREATE TABLE IF NOT EXISTS items (
      id INTEGER PRIMARY KEY AUTOINCREMENT,
      name TEXT,
      icon TEXT
    );
    `)

	if err != nil {
		log.Fatal(err)
	}

	return db
}
