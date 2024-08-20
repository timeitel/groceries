package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/timeitel/groceries/internal/domain"
	"github.com/timeitel/groceries/internal/infrastructure/data/db"
	_ "github.com/tursodatabase/go-libsql"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("DB path not passed as arg")
	}

	dbPath := os.Args[1]

	conn, err := sql.Open("libsql", dbPath)
	if err != nil {
		log.Fatal("Unable to open db", err)
	}

	q := db.New(conn)
	ctx := context.Background()

	user, err := q.CreateUser(ctx, db.CreateUserParams{
		ID:   uuid.New(),
		Name: "Cool guy",
	})
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

	fmt.Printf("User created %v", user)
	fmt.Printf("Cart created %v", cart)
}
