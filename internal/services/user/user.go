package user

import (
	"fmt"

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
	products, _ := s.Repo.GetProducts()
	user, _ := s.Repo.GetUser()

	d := home.Data{
		Products: products,
		Name:     user.Name,
	}

	return d
}

func (s *Service) AddItem(id string) error {
	user, _ := s.Repo.GetUser()
	fmt.Println(user)

	return nil
}
