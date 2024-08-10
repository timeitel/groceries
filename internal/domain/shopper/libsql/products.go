package libsql

import (
	"fmt"
	"os"

	"github.com/timeitel/groceries/internal/common/models"
)

func (r *Repository) GetProducts() (models.Products, error) {
	rows, err := r.DB.Query("SELECT * FROM products")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to execute query: %v\n", err)
		os.Exit(1)
	}
	defer rows.Close()

	var products models.Products

	for rows.Next() {
		var item models.Product

		if err := rows.Scan(&item.Id, &item.Name, &item.Description); err != nil {
			fmt.Println("Error scanning row:", err)
			return nil, err
		}

		products = append(products, item)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error during rows iteration:", err)
	}

	return products, nil
}

func (r *Repository) AddProduct(id string) (models.Product, error) {
	// rows, err := r.DB.Query("SELECT * FROM products")
	//
	var item models.Product

	return item, nil
}
