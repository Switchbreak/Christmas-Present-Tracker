package model

type Item struct {
	Name		string
	Description	string
	Link		string
}

type WishlistItem struct {
	ID			string `datastore:"-"`
	Item		Item
	BoughtBy	string
}

type BoughtItem struct {
	ID			string `datastore:"-"`
	Item		Item
	BoughtBy	string
}