package user

import (
	"github.com/timeitel/groceries/internal/models"
	_ "github.com/tursodatabase/go-libsql"
)

func NewService(repo repository) service {
	return service{
		repo,
	}
}

type servicer interface {
	GetItems() (models.Items, error)
}

type service struct {
	repo repository
}

func (s *service) GetItems() models.Items {
	items, _ := s.repo.GetItems()
	return items
}
