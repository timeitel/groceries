package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/timeitel/groceries/internal/domain"
	"github.com/timeitel/groceries/internal/infrastructure/data/db"
	_ "github.com/tursodatabase/go-libsql"
)

func main() {
	var dbPath string

	if len(os.Args) < 2 {
		dbPath = "file:/data/groceries.db"
	} else {
		dbPath = os.Args[1]
	}

	conn, err := sql.Open("libsql", dbPath)
	if err != nil {
		log.Fatal("Unable to open db", err)
	}

	q := db.New(conn)
	ctx := context.Background()

	user, err := q.CreateUser(ctx, "Cool guy")
	if err != nil {
		log.Fatalln("Creating user", err)
	}

	cart, err := q.CreateCart(ctx, db.CreateCartParams{
		UserID: user.ID,
		Name:   domain.NewSqlNullString("cart juan"),
	})
	if err != nil {
		log.Fatalln("Creating cart", err)
	}

	item, err := createItem(conn, ctx, "Apple", "Delicious")
	if err != nil {
		log.Fatalln("Creating item", err)
	}

	fmt.Printf("User created: %s, admin: %v\n", user.Name, user.IsAdmin.Bool)
	fmt.Printf("Cart created: %v\n", cart.Name.String)
	fmt.Printf("Initial item created: %v\n", item)
}

type item = db.Item

func createItem(db *sql.DB, ctx context.Context, name, description string) (db.Item, error) {
	row := db.QueryRowContext(ctx, createItemQuery, name, description)
	var i item
	err := row.Scan(&i.ID, &i.Name, &i.Description)
	return i, err
}

const createItemQuery = `-- name: CreateItem :one
INSERT INTO items (id, name, description)
    VALUES (1000, ?, ?)
RETURNING
    id, name, description
`
