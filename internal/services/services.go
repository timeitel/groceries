package services

import (
	"github.com/timeitel/groceries/internal/domain/shopper"
	"github.com/timeitel/groceries/internal/services/user"
)

func NewUser(repo shopper.RepoReader) user.Service {
	return user.Service{
		Repo: repo,
	}
}
