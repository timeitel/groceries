package user

import (
	"github.com/timeitel/groceries/internal/domain/shopper"
	"github.com/timeitel/groceries/internal/views/home"
	_ "github.com/tursodatabase/go-libsql"
)

func NewService(repo shopper.RepoReadWriter) Service {
	return Service{
		repo,
	}
}

type Service struct {
	Repo shopper.RepoReadWriter
}

func (s *Service) GetCart(jwt string) home.Data {
	items, _ := s.Repo.GetItems()
	user, _ := s.Repo.GetUser()

	d := home.Data{
		Items: items,
		Name:  user.Name,
	}

	return d
}

func (s *Service) AddItem(id string) error {
	user, _ := s.Repo.GetUser()

	return nil
}
