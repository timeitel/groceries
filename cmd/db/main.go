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

	tx, err := conn.Begin()
	if err != nil {
		log.Fatal("Begin tx", err)
	}
	defer tx.Rollback()

	queries := db.New(conn)
	qtx := queries.WithTx(tx)
	ctx := context.Background()
	userID := uuid.New()
	cartParams := db.CreateCartParams{
		UserID: userID,
		Name:   domain.NewSqlNullString("cart juan"),
	}

	cart, err := qtx.CreateCart(ctx, cartParams)
	if err != nil {
		log.Fatal("Create cart", err)
	}

	userParams := db.CreateUserParams{
		ActiveCartID: cart.ID,
		Name:         "cool guy",
	}

	user, err := qtx.CreateUser(ctx, userParams)
	if err != nil {
		log.Fatal("Create user", err)
	}

	tx.Commit()

	fmt.Printf("User created %v", user)
}
