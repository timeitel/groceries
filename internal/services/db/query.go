package db

import (
	"database/sql"
	"fmt"
	"os"
)

type Items []Item
type Item struct {
	ID          int
	Name        string
	Description string
}

func (i Item) String() string {
	return fmt.Sprintf("%d, %s, %s", i.ID, i.Name, i.Description)
}

func QueryItems(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM items")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to execute query: %v\n", err)
		os.Exit(1)
	}
	defer rows.Close()

	var items Items

	for rows.Next() {
		var item Item

		if err := rows.Scan(&item.ID, &item.Name, &item.Description); err != nil {
			fmt.Println("Error scanning row:", err)
			return
		}

		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error during rows iteration:", err)
	}

	fmt.Printf("items: %s", items)
}
