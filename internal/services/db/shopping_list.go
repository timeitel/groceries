package db

import "database/sql"

type Items []Item

type Item struct {
	ID          int
	Name        string
	Description string
}

type ShoppingListRepository interface {
	AddItem(item Item, quantity int) error
	RemoveItem(id string) error
}

type LibSqlShoppingListRepository struct {
	db *sql.DB
}

func (r *LibSqlShoppingListRepository) AddItem(userId string, itemId string) error {
	_, err := r.db.Exec("INSERT INTO shopping_lists (userId, itemId) VALUES (?, ?)", userId, itemId)

	return err
}

func (r *LibSqlShoppingListRepository) RemoveItem(item Item, quantity int) error {
	return nil
}

// type ShoppingListService struct {
// 	repo ShoppingListRepository
// }
//
// func (s *ShoppingListService) AddItem(item Item, quantity int) error {
// 	return s.repo.AddItem(item, quantity)
// }
