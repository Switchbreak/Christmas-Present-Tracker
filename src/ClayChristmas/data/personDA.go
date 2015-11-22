package data

import (
	"appengine"
	"appengine/datastore"
	"appengine/user"

	"ClayChristmas/model"
)

func GetLoggedInPersonKey(appContext appengine.Context) (*datastore.Key) {
	currentUser := user.Current(appContext)
	return datastore.NewKey(appContext, "Person", currentUser.ID, 0, nil)
}

func GetLoggedInPerson(appContext appengine.Context) (*model.Person, error) {
	key := GetLoggedInPersonKey( appContext )
	
	var person model.Person
	if err := datastore.Get( appContext, key, &person ); err != nil {
		return nil, err
	} 
	
	return &person, nil
}
