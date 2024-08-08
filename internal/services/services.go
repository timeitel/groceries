package services

import (
	"github.com/timeitel/groceries/internal/domain/user"
	"github.com/timeitel/groceries/internal/services/shopping"
)

func NewShopping(repo user.Repository) shopping.Service {
	return shopping.Service{
		Repo: repo,
	}
}
