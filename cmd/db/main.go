package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/tursodatabase/go-libsql"
)

func main() {
	ctx := context.Background()
	url, exists := os.LookupEnv("DB_URL")
	if !exists {
		log.Fatalln("No env set under DB_URL")
	}

	db, err := sql.Open("libsql", url)
	if err != nil {
		log.Fatal("Unable to open db", url, err)
	}

	res, err := db.ExecContext(ctx, "schema.sql")
	if err != nil {
		log.Fatalln("Unable to create schema", err)
	}

	// TODO: db seeding for products and admin user

	fmt.Printf("Rows Affected: %d\n", res)
	fmt.Printf("Created database schema")
}
