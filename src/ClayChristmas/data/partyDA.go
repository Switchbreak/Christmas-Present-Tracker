package data

import (
	"appengine"
	"appengine/datastore"

	"ClayChristmas/model"
)

var ErrNoSuchEntity = datastore.ErrNoSuchEntity

func GetParty(appContext appengine.Context, title string) (*model.Party, error) {
	partyKey := datastore.NewKey(appContext, "Party", title, 0, nil)

	var party model.Party
	err := datastore.Get(appContext, partyKey, &party)
	return &party, err
}

func IsInvited(appContext appengine.Context, title string, personID string) (bool, error) {
	partyKey := datastore.NewKey(appContext, "Party", title, 0, nil)

	invitedPersonKey := datastore.NewKey(appContext, "InvitedPerson", personID, 0, partyKey)

	var person model.InvitedPerson
	err := datastore.Get(appContext, invitedPersonKey, &person)
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

	parties := make([]model.Party, len(keys))
	err = datastore.GetMulti(appContext, keys, parties)

	return parties, err
}

func GetInvitedPeople(appContext appengine.Context, partyID string) ([]model.Person, error) {
	partyKey := datastore.NewKey(appContext, "Party", partyID, 0, nil)

	var invitedPeople []model.InvitedPerson
	if _, err := datastore.NewQuery("InvitedPerson").Ancestor(partyKey).GetAll(appContext, &invitedPeople); err != nil {
		return nil, err
	}

	keys := make([]*datastore.Key, len(invitedPeople))
	for index, person := range invitedPeople {
		keys[index] = person.Person
	}

	people := make([]model.Person, len(invitedPeople))
	if err := datastore.GetMulti(appContext, keys, people); err != nil {
		return nil, err
	}

	return people, nil
}

func InvitePerson(appContext appengine.Context, partyID string, personID string) error {
	partyKey := datastore.NewKey(appContext, "Party", partyID, 0, nil)
	personKey := datastore.NewKey(appContext, "InvitedPerson", personID, 0, partyKey)

	var person model.InvitedPerson
	person.Person = datastore.NewKey(appContext, "Person", personID, 0, nil)

	_, err := datastore.Put(appContext, personKey, &person)
	return err
}

func UpdateParty(appContext appengine.Context, party *model.Party) error {
	partyKey := datastore.NewKey(appContext, "Party", party.Title, 0, nil)

	_, err := datastore.Put(appContext, partyKey, party)
	return err
}

func DeleteParty(appContext appengine.Context, party *model.Party) error {
	partyKey := datastore.NewKey(appContext, "Party", party.Title, 0, nil)

	return datastore.Delete(appContext, partyKey)
}
