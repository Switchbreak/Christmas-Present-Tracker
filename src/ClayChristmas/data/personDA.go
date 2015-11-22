package data

import (
	"errors"
	
	"appengine"
	"appengine/datastore"
	"appengine/user"

	"ClayChristmas/model"
)

func GetLoggedInPerson(appContext appengine.Context) (*datastore.Key, error) {
	currentUser := user.Current(appContext)
	
	query := datastore.NewQuery("Person").Filter("User.ID =", currentUser.ID).Limit(1).KeysOnly();
	keys, err := query.GetAll( appContext, nil )
	if err != nil {
		return nil, err
	}
	if len( keys ) < 1 {
		return nil, errors.New( "User not found" )
	}
	return keys[0], nil
}

func GetPerson(appContext appengine.Context, personKey *datastore.Key) (*model.Person, error) {
	var person model.Person
	if err := datastore.Get( appContext, personKey, &person ); err != nil {
		return nil, err
	} 
	
	return &person, nil
}

func GetPersonByName(appContext appengine.Context, name string) (*model.Person, error) {
	personKey := datastore.NewKey( appContext, "Person", name, 0, nil )
	return GetPerson( appContext, personKey )
}

func UpdatePerson(appContext appengine.Context, person *model.Person) {
	
}

func checkPersonOwner( appContext appengine.Context, person *model.Person) error {
	if !user.IsAdmin( appContext ) {
		currentUser := user.Current( appContext )
		if currentUser.ID != person.User.ID {
			return errors.New( "User does not have permission" )
		}
	}
	
	return nil
}