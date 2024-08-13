package shopper

import (
	"context"
	"fmt"
	"log"

	"github.com/timeitel/groceries/internal/common/types"
)

func (r *libSqlRepository) GetShopper() (shopper, error) {
	// TODO:
	u, err := r.db.GetShopper(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	return u, nil
}

func (r *libSqlRepository) GetProducts() (types.Products, error) {
	p, err := r.db.GetProducts(context.Background())

	if err != nil {
		fmt.Println("Unable to get products", err)
		return nil, err
	}

	return p, nil
}
