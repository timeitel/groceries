package models

type Carts []Cart

type Cart struct {
	Id     string
	Items  CartItems
	Name   string
	UserId string
}

type CartItems []CartItem

type CartItem struct {
	ItemId   string
	Quantity int
}
