package libsql

import (
	"fmt"
	"log"
	"os"

	"github.com/timeitel/groceries/internal/common/models"
	"github.com/timeitel/groceries/internal/infrastructure/data"
	"github.com/timeitel/groceries/internal/infrastructure/data/db"
)

func (r *Repository) GetUser() (db.User, error) {
	// TODO:
	u, err := r.GetUser()
	if err != nil {
		log.Fatal(err)
	}

	return u, nil
}

func (r *Repository) GetProducts() (models.Products, error) {
	rows, err := r.DB.Query("SELECT * FROM products")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to execute query: %v\n", err)
		os.Exit(1)
	}
	defer rows.Close()

	var products models.Products

	for rows.Next() {
		var p data.Product

		if err := rows.Scan(&p.Id, &p.Name, &p.Description); err != nil {
			fmt.Printf("Error scanning row: %v", err)
			return nil, err
		}

		products = append(products, p)
	}

	if err := rows.Err(); err != nil {
		fmt.Printf("Error in rows: %v", err)
	}

	return products, nil
}

func (r *Repository) AddProduct(id string) (models.Product, error) {
	// rows, err := r.DB.Query("SELECT * FROM products")
	//
	var item models.Product

	return item, nil
}
