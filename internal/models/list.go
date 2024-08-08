package models

type Lists []List

type List struct {
	Id    string
	Items Items
	Name  string
}
