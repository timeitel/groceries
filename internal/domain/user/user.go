package user

import "github.com/timeitel/groceries/internal/models"

type User struct {
	Id    int
	lists models.Lists
}
