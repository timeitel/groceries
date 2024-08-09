package models

type Users []User

type User struct {
	Id      string
	Name    string
	IsAdmin bool
}
