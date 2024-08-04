package db

import (
	"bufio"
	"database/sql"
	"os"

	_ "github.com/tursodatabase/go-libsql"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func Init() *sql.DB {
	path := os.Getenv("DB_URL_FILE")

	file, err := os.Open(path)
	check(err)

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	url := scanner.Text()

	db, err := sql.Open("libsql", url)
	check(err)

	const create string = `
  CREATE TABLE IF NOT EXISTS items (
  id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  name VARCHAR NOT NULL UNIQUE,
  description TEXT
  );`

	_, err = db.Exec(create)
	check(err)

	const insert string = `
  INSERT INTO items (name, description)
  VALUES ('Apple', 'A delicious fruit.');
  `
	_, err = db.Exec(insert)
	check(err)

	return db
}
