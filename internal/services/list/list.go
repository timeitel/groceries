package services

import (
	"github.com/timeitel/groceries/internal/domain/user"
	"github.com/timeitel/groceries/internal/models"
	_ "github.com/tursodatabase/go-libsql"
)

func NewList(repo user.Repository) List {
	return List{
		repo,
	}
}

type List struct {
	repo user.Repository
}

func (s *List) GetItems() models.Items {
	items, _ := s.repo.GetItems()
	return items
}
