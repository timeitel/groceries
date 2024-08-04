package services

import (
	"bufio"
	"database/sql"
	"log"
	"os"

	_ "github.com/tursodatabase/go-libsql"
)

func InitDatabase() {
	path := os.Getenv("DB_URL_FILE")

	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Unable to get db url file: %s", path)
	}

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	url := scanner.Text()

	db, err := sql.Open("libsql", url)
	if err != nil {
		log.Fatalf("Failed to open db: %s", err)
	}
	defer db.Close()
}
