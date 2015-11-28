package data

import (
	"appengine"
	"appengine/datastore"

	"ClayChristmas/model"
)

func GetWishlistItem(appContext appengine.Context, partyID string, personID string, wishlistID string) (*model.WishlistItem, error) {
	partyKey := datastore.NewKey(appContext, "Party", partyID, 0, nil)
	personKey := datastore.NewKey(appContext, "InvitedPerson", personID, 0, partyKey)
	itemKey := datastore.NewKey(appContext, "WishlistItem", wishlistID, 0, personKey)

	var wishlistItem model.WishlistItem
	err := datastore.Get(appContext, itemKey, &wishlistItem)

	wishlistItem.ID = wishlistID

	return &wishlistItem, err
}

func GetWishlist(appContext appengine.Context, partyID string, personID string) ([]model.WishlistItem, error) {
	partyKey := datastore.NewKey(appContext, "Party", partyID, 0, nil)
	personKey := datastore.NewKey(appContext, "InvitedPerson", personID, 0, partyKey)

	var wishlist []model.WishlistItem
	keys, err := datastore.NewQuery("WishlistItem").Ancestor(personKey).GetAll(appContext, &wishlist)
	
	for index, key := range keys {
		wishlist[index].ID = key.StringID()
	}

	return wishlist, err
}

func CreateWishlistItem(appContext appengine.Context, partyID string, personID string, wishlistItem *model.WishlistItem) (*model.WishlistItem, error) {
	partyKey := datastore.NewKey(appContext, "Party", partyID, 0, nil)
	personKey := datastore.NewKey(appContext, "InvitedPerson", personID, 0, partyKey)
	itemKey := datastore.NewIncompleteKey(appContext, "WishlistItem", personKey)

	key, err := datastore.Put(appContext, itemKey, wishlistItem)
	wishlistItem.ID = key.StringID()

	return wishlistItem, err
}

func UpdateWishlistItem(appContext appengine.Context, partyID string, personID string, wishlistItem *model.WishlistItem) (*model.WishlistItem, error) {
	partyKey := datastore.NewKey(appContext, "Party", partyID, 0, nil)
	personKey := datastore.NewKey(appContext, "InvitedPerson", personID, 0, partyKey)
	itemKey := datastore.NewKey(appContext, "WishlistItem", wishlistItem.ID, 0, personKey)

	key, err := datastore.Put(appContext, itemKey, wishlistItem)
	wishlistItem.ID = key.StringID()

	return wishlistItem, err
}

func DeleteWishlistItem(appContext appengine.Context, partyID string, personID string, wishlistItem *model.WishlistItem) error {
	partyKey := datastore.NewKey(appContext, "Party", partyID, 0, nil)
	personKey := datastore.NewKey(appContext, "InvitedPerson", personID, 0, partyKey)
	itemKey := datastore.NewKey(appContext, "WishlistItem", wishlistItem.ID, 0, personKey)

	return datastore.Delete(appContext, itemKey)
}

func GetBoughtItem(appContext appengine.Context, partyID string, personID string, boughtItemID string) (*model.BoughtItem, error) {
	partyKey := datastore.NewKey(appContext, "Party", partyID, 0, nil)
	personKey := datastore.NewKey(appContext, "InvitedPerson", personID, 0, partyKey)
	itemKey := datastore.NewKey(appContext, "BoughtItem", boughtItemID, 0, personKey)

	var boughtItem model.BoughtItem
	err := datastore.Get(appContext, itemKey, &boughtItem)

	return &boughtItem, err
}

func GetBoughtItems(appContext appengine.Context, partyID string, personID string) ([]model.BoughtItem, error) {
	partyKey := datastore.NewKey(appContext, "Party", partyID, 0, nil)
	personKey := datastore.NewKey(appContext, "InvitedPerson", personID, 0, partyKey)

	var boughtItems []model.BoughtItem
	keys, err := datastore.NewQuery("BoughtItem").Ancestor(personKey).GetAll(appContext, &boughtItems)
	
	for index, key := range keys {
		boughtItems[index].ID = key.StringID()
	}

	return boughtItems, err
}

func CreateBoughtItem(appContext appengine.Context, partyID string, personID string, boughtItem *model.BoughtItem) (*model.BoughtItem, error) {
	partyKey := datastore.NewKey(appContext, "Party", partyID, 0, nil)
	personKey := datastore.NewKey(appContext, "InvitedPerson", personID, 0, partyKey)
	itemKey := datastore.NewIncompleteKey(appContext, "BoughtItem", personKey)

	key, err := datastore.Put(appContext, itemKey, boughtItem)
	boughtItem.ID = key.StringID()

	return boughtItem, err
}

func UpdateBoughtItem(appContext appengine.Context, partyID string, personID string, boughtItem *model.BoughtItem) (*model.BoughtItem, error) {
	partyKey := datastore.NewKey(appContext, "Party", partyID, 0, nil)
	personKey := datastore.NewKey(appContext, "InvitedPerson", personID, 0, partyKey)
	itemKey := datastore.NewKey(appContext, "BoughtItem", boughtItem.ID, 0, personKey)

	key, err := datastore.Put(appContext, itemKey, boughtItem)
	boughtItem.ID = key.StringID()

	return boughtItem, err
}

func DeleteBoughtItem(appContext appengine.Context, partyID string, personID string, boughtItem *model.BoughtItem) error {
	partyKey := datastore.NewKey(appContext, "Party", partyID, 0, nil)
	personKey := datastore.NewKey(appContext, "InvitedPerson", personID, 0, partyKey)
	itemKey := datastore.NewKey(appContext, "BoughtItem", boughtItem.ID, 0, personKey)

	return datastore.Delete(appContext, itemKey)
}
