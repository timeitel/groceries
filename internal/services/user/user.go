package user

import (
	"fmt"
	"log"

	"github.com/timeitel/groceries/internal/domain/shopper"
	"github.com/timeitel/groceries/internal/views/home"
	_ "github.com/tursodatabase/go-libsql"
)

func NewService(repo shopper.ReadWriter) Service {
	return Service{
		repo,
	}
}

type Service struct {
	Repo shopper.ReadWriter
}

func (s *Service) GetHomeData(jwt string) home.Data {
	products, err := s.Repo.GetProducts()
	if err != nil {
		log.Fatal(err)
	}

	shopper, err := s.Repo.GetShopper()
	if err != nil {
		log.Fatal(err)
	}

	d := home.Data{
		Products: products,
		Name:     shopper.User.Name,
	}

	return d
}

func (s *Service) AddProductToCart(productID int, quantity int) error {
	user, err := s.Repo.GetShopper()
	if err != nil {
		fmt.Println(err)
		return err
	}

	if err = s.Repo.AddProductToCart(productID, 1); err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println(user)

	return nil
}
