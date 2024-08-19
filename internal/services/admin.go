package services

import (
	"github.com/google/uuid"
	"github.com/timeitel/groceries/internal/domain/catalogue"
	"github.com/timeitel/groceries/internal/domain/item"
	_ "github.com/tursodatabase/go-libsql"
)

type Admin struct {
	items     item.RepoWriter
	catalogue catalogue.RepoReader
}

// TODO: inject
func NewAdmin() Admin {
	return Admin{
		items:     item.NewLibSqlRepo(),
		catalogue: catalogue.NewLibSqlRepo(),
	}
}

func (s *Admin) CreateItem(name, description string) (*item.Item, error) {
	i, err := s.items.Create(name, description)
	if err != nil {
		return nil, err
	}

	return i, nil
}

func (s *Admin) GetItem(id uuid.UUID) (*item.Item, error) {
	i, err := s.catalogue.GetItem(id)
	if err != nil {
		return nil, err
	}

	return i, nil
}

func (s *Admin) GetItems() (item.Items, error) {
	items, err := s.catalogue.GetItems()
	if err != nil {
		return nil, err
	}

	return *items, nil
}

func (s *Admin) UpdateItem(id uuid.UUID, name string, description string) (*item.Item, error) {
	i, err := s.items.Update(id, name, description)
	if err != nil {
		return nil, err
	}

	return i, nil
}

func (s *Admin) DeleteItem(id uuid.UUID) error {
	err := s.items.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
