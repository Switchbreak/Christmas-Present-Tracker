package model

import (

)

type Item struct {
	Name		string
	Description	string
	Link		string
}

type WishlistItem struct {
	Item		Item
	BoughtBy	Person
}

type BoughtItem struct {
	Item		Item
	BoughtFor	Person
}