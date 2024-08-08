package shopping

import (
	"github.com/timeitel/groceries/internal/domain/user"
	"github.com/timeitel/groceries/internal/models"
	_ "github.com/tursodatabase/go-libsql"
)

func NewService(repo user.Repository) Service {
	return Service{
		repo,
	}
}

type Service struct {
	Repo user.Repository
}

func (s *Service) GetItems() models.Items {
	items, _ := s.Repo.GetItems()
	return items
}
