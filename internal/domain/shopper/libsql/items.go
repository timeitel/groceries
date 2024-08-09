package libsql

import (
	"fmt"
	"os"

	"github.com/timeitel/groceries/internal/types/models"
)

func (r *Repository) GetItems() (models.Items, error) {
	rows, err := r.DB.Query("SELECT * FROM items")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to execute query: %v\n", err)
		os.Exit(1)
	}
	defer rows.Close()

	var items models.Items

	for rows.Next() {
		var item models.Item

		if err := rows.Scan(&item.Id, &item.Name); err != nil {
			fmt.Println("Error scanning row:", err)
			return nil, err
		}

		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error during rows iteration:", err)
	}

	return items, nil
}
