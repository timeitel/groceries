package services

import (
	"github.com/timeitel/groceries/internal/domain/shopper"
	"github.com/timeitel/groceries/internal/services/user"
)

func NewUser(repo shopper.ReadWriter) user.Service {
	return user.Service{
		Repo: repo,
	}
}
