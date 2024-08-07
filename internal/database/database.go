package database

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

// func (r *repository) addItem(userId string, itemId string) error {
// 	_, err := r.db.Exec("INSERT INTO shopping_lists (userId, itemId) VALUES (?, ?)", userId, itemId)
//
// 	return err
// }
