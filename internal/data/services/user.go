package services

import (
	"github.com/timeitel/groceries/internal/data/models"
	"github.com/timeitel/groceries/internal/data/repositories"
	_ "github.com/tursodatabase/go-libsql"
)

type getter interface {
	getItems() (models.Items, error)
}

type service struct {
	repo repositories.User
}

func (s *service) GetItems() error {
	return nil
}

func NewUser(repo repositories.User) service {
	return service{
		repo,
	}
}
