package shopping

import (
	"github.com/timeitel/groceries/internal/domain/user"
	"github.com/timeitel/groceries/internal/views/home"
	_ "github.com/tursodatabase/go-libsql"
)

func NewService(repo user.RepoReader) Service {
	return Service{
		repo,
	}
}

type Service struct {
	Repo user.RepoReader
}

func (s *Service) GetHomeData() home.Data {
	items, _ := s.Repo.GetItems()
	name, _ := s.Repo.GetUserName()

	d := home.Data{
		Items: items,
		Name:  name,
	}

	return d
}
