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

func (s *Admin) GetProducts() (types.Products, error) {
	p, err := s.shopperRepo.GetProducts()
	if err != nil {
		return nil, err
	}

	return p, nil
}
