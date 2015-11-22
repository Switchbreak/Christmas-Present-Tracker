package data

import (
	"errors"
	"time"
	
	"appengine"
	"appengine/datastore"
	"appengine/user"

	"ClayChristmas/model"
)

func GetParties(appContext appengine.Context) ([]model.Party, error) {
	personKey, err := GetLoggedInPerson(appContext)
	if err != nil {
		return nil, err
	}

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
	partyKey := datastore.NewKey( appContext, "Party", title, 0, nil )
	
	var party model.Party
	err := datastore.Get( appContext, partyKey, &party )
	return &party, err
}

func UpdateParty(appContext appengine.Context, party *model.Party) error {
	partyKey := datastore.NewKey( appContext, "Party", party.Title, 0, nil )
	
	var checkParty model.Party
	err := datastore.Get( appContext, partyKey, &checkParty )
	if err == nil {
		if err = checkPartyOwner( appContext, &checkParty ); err != nil {
			return err
		}
		
		party.CreatedBy = checkParty.CreatedBy
		party.CreatedDate = checkParty.CreatedDate
	} else if err == datastore.ErrNoSuchEntity {
		currentUser, err := GetLoggedInPerson( appContext );
		if err != nil {
			return err
		}
		
		party.CreatedBy = currentUser.StringID() 
		party.CreatedDate = time.Now()
	} else {
		return err
	}
	
	_, err = datastore.Put( appContext, partyKey, &party )
	return err
}

func DeleteParty(appContext appengine.Context, party *model.Party) error {
	partyKey := datastore.NewKey( appContext, "Party", party.Title, 0, nil )
	
	var checkParty model.Party
	if err := datastore.Get( appContext, partyKey, &checkParty ); err != nil {
		return err
	}
	if err := checkPartyOwner( appContext, &checkParty ); err != nil {
		return err
	}
	
	return datastore.Delete( appContext, partyKey )
}

func checkPartyOwner( appContext appengine.Context, party *model.Party ) error {
	if !user.IsAdmin( appContext ) {
		currentUserKey, err := GetLoggedInPerson( appContext );
		if err != nil {
			return err
		}
		
		if party.CreatedBy != currentUserKey.StringID() {
			return errors.New( "User does not have permission" )
		}
	}
	
	return nil
} 