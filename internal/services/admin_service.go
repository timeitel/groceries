package services

import (
	"github.com/timeitel/groceries/internal/domain/admin"
	"github.com/timeitel/groceries/internal/domain/shopper"
	"github.com/timeitel/groceries/internal/infrastructure/data/db"
	"github.com/timeitel/groceries/internal/types"
	_ "github.com/tursodatabase/go-libsql"
)

type Admin struct {
	adminRepo   admin.RepoWriter
	shopperRepo shopper.RepoReadWriter
}

func NewAdmin(adminRepo admin.RepoWriter, shopperRepo shopper.RepoReadWriter) Admin {
	return Admin{
		adminRepo:   adminRepo,
		shopperRepo: shopperRepo,
	}
}

func (s *Admin) CreateProduct(name, description string) (*db.Product, error) {
	p, err := s.adminRepo.CreateProduct(name, description)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (s *Admin) GetProduct(id int64) (*db.Product, error) {
	p, err := s.shopperRepo.GetProduct(id)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (s *Admin) GetProducts() (types.Products, error) {
	p, err := s.shopperRepo.GetProducts()
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (s *Admin) UpdateProduct(id int64, name string, description string) error {
	err := s.adminRepo.UpdateProduct(id, name, description)
	if err != nil {
		return err
	}

	return nil
}

func (s *Admin) DeleteProduct(id int64) error {
	err := s.adminRepo.DeleteProduct(id)
	if err != nil {
		return err
	}

	return nil
}
