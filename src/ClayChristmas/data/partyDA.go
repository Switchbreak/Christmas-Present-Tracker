package data

import (
	"appengine"
	"appengine/datastore"

	"ClayChristmas/model"
)

func GetParty(appContext appengine.Context, title string) (*model.Party, error) {
	partyKey := datastore.NewKey(appContext, "Party", title, 0, nil)

	var party model.Party
	err := datastore.Get(appContext, partyKey, &party)
	return &party, err
}

func IsInvited(appContext appengine.Context, title string, personID string) (bool, error) {
	partyKey := datastore.NewKey(appContext, "Party", title, 0, nil)

	invitedPersonKey := datastore.NewKey(appContext, "InvitedPerson", personID, 0, partyKey)
	err := datastore.Get(appContext, invitedPersonKey, nil)
	if err == datastore.ErrNoSuchEntity {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return true, nil
}

func GetPartiesInvited(appContext appengine.Context, personID string) ([]model.Party, error) {
	personKey := datastore.NewKey(appContext, "Person", personID, 0, nil)

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

func GetInvitedPeople(appContext appengine.Context, partyID string) ([]model.Person, error) {
	partyKey := datastore.NewKey(appContext, "Party", partyID, 0, nil)

	var invitedPeople []datastore.Key
	if _, err := datastore.NewQuery("InvitedPerson").Ancestor(partyKey).GetAll(appContext, &invitedPeople); err != nil {
		return nil, err
	}

	keys := make([]*datastore.Key, len(invitedPeople))
	for index, person := range invitedPeople {
		keys[index] = &person
	}

	var people []model.Person
	if err := datastore.GetMulti(appContext, keys, &people); err != nil {
		return nil, err
	}

	return people, nil
}

func InvitePerson(appContext appengine.Context, partyID string, personID string) error {
	partyKey := datastore.NewKey(appContext, "Party", partyID, 0, nil)
	personKey := datastore.NewKey(appContext, "Person", personID, 0, nil)

	_, err := datastore.Put(appContext, partyKey, personKey)
	return err
}

func UpdateParty(appContext appengine.Context, party *model.Party) error {
	partyKey := datastore.NewKey(appContext, "Party", party.Title, 0, nil)

	//	var checkParty model.Party
	//	err := datastore.Get( appContext, partyKey, &checkParty )
	//	if err == nil {
	//		if err = checkPartyOwner( appContext, &checkParty ); err != nil {
	//			return err
	//		}
	//
	//		party.CreatedBy = checkParty.CreatedBy
	//		party.CreatedDate = checkParty.CreatedDate
	//	} else if err == datastore.ErrNoSuchEntity {
	//		currentUser, err := GetLoggedInPerson( appContext );
	//		if err != nil {
	//			return err
	//		}
	//
	//		party.CreatedBy = currentUser.StringID()
	//		party.CreatedDate = time.Now()
	//	} else {
	//		return err
	//	}

	_, err := datastore.Put(appContext, partyKey, &party)
	return err
}

func DeleteParty(appContext appengine.Context, party *model.Party) error {
	partyKey := datastore.NewKey(appContext, "Party", party.Title, 0, nil)

	//	var checkParty model.Party
	//	if err := datastore.Get( appContext, partyKey, &checkParty ); err != nil {
	//		return err
	//	}
	//	if err := checkPartyOwner( appContext, &checkParty ); err != nil {
	//		return err
	//	}

	return datastore.Delete(appContext, partyKey)
}

//func checkPartyOwner( appContext appengine.Context, party *model.Party ) error {
//	if !user.IsAdmin( appContext ) {
//		currentUserKey, err := GetLoggedInPerson( appContext );
//		if err != nil {
//			return err
//		}
//
//		if party.CreatedBy != currentUserKey.StringID() {
//			return errors.New( "User does not have permission" )
//		}
//	}
//
//	return nil
//}
