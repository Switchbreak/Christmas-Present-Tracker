package data

import (
	"strconv"
	
	"appengine"
	"appengine/datastore"

	"ClayChristmas/model"
)

func GetWishlistItem(appContext appengine.Context, partyID string, personID string, wishlistID string) (*model.WishlistItem, error) {
	id, err := strconv.Atoi( wishlistID )
	if err != nil {
		return nil, err
	}
	
	partyKey := datastore.NewKey(appContext, "Party", partyID, 0, nil)
	personKey := datastore.NewKey(appContext, "InvitedPerson", personID, 0, partyKey)
	itemKey := datastore.NewKey(appContext, "WishlistItem", "", int64( id ), personKey)

	var wishlistItem model.WishlistItem
	err = datastore.Get(appContext, itemKey, &wishlistItem)

	wishlistItem.ID = wishlistID

	return &wishlistItem, err
}

func GetWishlist(appContext appengine.Context, partyID string, personID string) ([]model.WishlistItem, error) {
	partyKey := datastore.NewKey(appContext, "Party", partyID, 0, nil)
	personKey := datastore.NewKey(appContext, "InvitedPerson", personID, 0, partyKey)

	var wishlist []model.WishlistItem
	keys, err := datastore.NewQuery("WishlistItem").Ancestor(personKey).GetAll(appContext, &wishlist)
	
	for index, key := range keys {
		wishlist[index].ID = strconv.Itoa( int( key.IntID() ) )
	}

	return wishlist, err
}

func GetWishlistCount(appContext appengine.Context, partyID string, personID string) (int, error) {
	partyKey := datastore.NewKey(appContext, "Party", partyID, 0, nil)
	personKey := datastore.NewKey(appContext, "InvitedPerson", personID, 0, partyKey)

	return datastore.NewQuery("WishlistItem").Ancestor(personKey).KeysOnly().Count( appContext ) 
}

func CreateWishlistItem(appContext appengine.Context, partyID string, personID string, wishlistItem *model.WishlistItem) (*model.WishlistItem, error) {
	partyKey := datastore.NewKey(appContext, "Party", partyID, 0, nil)
	personKey := datastore.NewKey(appContext, "InvitedPerson", personID, 0, partyKey)
	itemKey := datastore.NewIncompleteKey(appContext, "WishlistItem", personKey)

	key, err := datastore.Put(appContext, itemKey, wishlistItem)
	wishlistItem.ID = strconv.Itoa( int( key.IntID() ) )

	return wishlistItem, err
}

func UpdateWishlistItem(appContext appengine.Context, partyID string, personID string, wishlistItem *model.WishlistItem) (*model.WishlistItem, error) {
	id, err := strconv.Atoi( wishlistItem.ID )
	if err != nil {
		return nil, err
	}
	
	partyKey := datastore.NewKey(appContext, "Party", partyID, 0, nil)
	personKey := datastore.NewKey(appContext, "InvitedPerson", personID, 0, partyKey)
	itemKey := datastore.NewKey(appContext, "WishlistItem", "", int64( id ), personKey)

	key, err := datastore.Put(appContext, itemKey, wishlistItem)
	wishlistItem.ID = strconv.Itoa( int( key.IntID() ) )

	return wishlistItem, err
}

func DeleteWishlistItem(appContext appengine.Context, partyID string, personID string, wishlistItem *model.WishlistItem) error {
	id, err := strconv.Atoi( wishlistItem.ID )
	if err != nil {
		return err
	}
	
	partyKey := datastore.NewKey(appContext, "Party", partyID, 0, nil)
	personKey := datastore.NewKey(appContext, "InvitedPerson", personID, 0, partyKey)
	itemKey := datastore.NewKey(appContext, "WishlistItem", "", int64( id ), personKey)

	return datastore.Delete(appContext, itemKey)
}

func GetBoughtItem(appContext appengine.Context, partyID string, personID string, boughtItemID string) (*model.BoughtItem, error) {
	id, err := strconv.Atoi( boughtItemID )
	if err != nil {
		return nil, err
	}
	
	partyKey := datastore.NewKey(appContext, "Party", partyID, 0, nil)
	personKey := datastore.NewKey(appContext, "InvitedPerson", personID, 0, partyKey)
	itemKey := datastore.NewKey(appContext, "BoughtItem", "", int64( id ), personKey)

	var boughtItem model.BoughtItem
	err = datastore.Get(appContext, itemKey, &boughtItem)

	return &boughtItem, err
}

func GetBoughtItems(appContext appengine.Context, partyID string, personID string) ([]model.BoughtItem, error) {
	partyKey := datastore.NewKey(appContext, "Party", partyID, 0, nil)
	personKey := datastore.NewKey(appContext, "InvitedPerson", personID, 0, partyKey)

	var boughtItems []model.BoughtItem
	keys, err := datastore.NewQuery("BoughtItem").Ancestor(personKey).GetAll(appContext, &boughtItems)
	
	for index, key := range keys {
		boughtItems[index].ID = strconv.Itoa( int( key.IntID() ) )
	}

	return boughtItems, err
}

func CreateBoughtItem(appContext appengine.Context, partyID string, personID string, boughtItem *model.BoughtItem) (*model.BoughtItem, error) {
	partyKey := datastore.NewKey(appContext, "Party", partyID, 0, nil)
	personKey := datastore.NewKey(appContext, "InvitedPerson", personID, 0, partyKey)
	itemKey := datastore.NewIncompleteKey(appContext, "BoughtItem", personKey)

	key, err := datastore.Put(appContext, itemKey, boughtItem)
	boughtItem.ID = strconv.Itoa( int( key.IntID() ) )

	return boughtItem, err
}

func UpdateBoughtItem(appContext appengine.Context, partyID string, personID string, boughtItem *model.BoughtItem) (*model.BoughtItem, error) {
	id, err := strconv.Atoi( boughtItem.ID )
	if err != nil {
		return nil, err
	}
	
	partyKey := datastore.NewKey(appContext, "Party", partyID, 0, nil)
	personKey := datastore.NewKey(appContext, "InvitedPerson", personID, 0, partyKey)
	itemKey := datastore.NewKey(appContext, "BoughtItem", "", int64( id ), personKey)

	key, err := datastore.Put(appContext, itemKey, boughtItem)
	boughtItem.ID = strconv.Itoa( int( key.IntID() ) )

	return boughtItem, err
}

func DeleteBoughtItem(appContext appengine.Context, partyID string, personID string, boughtItem *model.BoughtItem) error {
	id, err := strconv.Atoi( boughtItem.ID )
	if err != nil {
		return err
	}
	
	partyKey := datastore.NewKey(appContext, "Party", partyID, 0, nil)
	personKey := datastore.NewKey(appContext, "InvitedPerson", personID, 0, partyKey)
	itemKey := datastore.NewKey(appContext, "BoughtItem", "", int64( id ), personKey)

	return datastore.Delete(appContext, itemKey)
}
