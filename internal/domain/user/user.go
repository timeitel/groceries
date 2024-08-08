package user

import "github.com/timeitel/groceries/internal/models"

type User struct {
	Id    int
	Name  string
	lists models.Lists
	// recipesl models.Recipes
}
