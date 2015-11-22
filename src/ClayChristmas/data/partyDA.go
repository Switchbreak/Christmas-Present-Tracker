package data

import (
	"appengine"
	"appengine/datastore"

	"ClayChristmas/model"
)

func GetParties(appContext appengine.Context) ([]model.Party, error) {
	personKey := GetLoggedInPersonKey(appContext)

	query := datastore.NewQuery("InvitedPerson").Filter("Person =", personKey).KeysOnly()
	keys, err := query.GetAll(appContext, nil)
	if err != nil {
		return nil, err
	}

	for index, key := range keys {
		keys[index] = key.Parent()
	}

	var parties []model.Party
	err = datastore.GetMulti(appContext, keys, &parties)
	return parties, err
}

func GetParty(appContext appengine.Context, title string) (*model.Party, error) {
	partyKey := datastore.NewKey( appContext, "Party", title, 0, nil );
	
	var party model.Party
	err := datastore.Get( appContext, partyKey, &party )
	return &party, err
}