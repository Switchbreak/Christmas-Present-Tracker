package data

import (
	"errors"

	"appengine"
	"appengine/datastore"
	"appengine/user"

	"ClayChristmas/model"
)

var UserNotLoggedIn = errors.New("User not logged in")
var UserNotFound = errors.New("User not found")

func GetLoggedInPerson(appContext appengine.Context) (string, error) {
	currentUser := user.Current(appContext)
	if currentUser == nil {
		return "", UserNotLoggedIn
	}

	query := datastore.NewQuery("Person").Filter("User.ID =", currentUser.ID).Limit(1).KeysOnly()
	keys, err := query.GetAll(appContext, nil)
	if err != nil {
		return "", err
	}
	if len(keys) < 1 {
		return "", UserNotFound
	}
	return keys[0].StringID(), nil
}

func GetPeople(appContext appengine.Context) ([]model.Person, error) {
	query := datastore.NewQuery("Person")

	var people []model.Person
	_, err := query.GetAll(appContext, &people)

	return people, err
}

func GetPerson(appContext appengine.Context, name string) (*model.Person, error) {
	personKey := datastore.NewKey(appContext, "Person", name, 0, nil)

	var person model.Person
	if err := datastore.Get(appContext, personKey, &person); err != nil {
		return nil, err
	}

	return &person, nil
}

func UpdatePerson(appContext appengine.Context, person *model.Person) error {
	personKey := datastore.NewKey(appContext, "Person", person.Name, 0, nil)

	_, err := datastore.Put(appContext, personKey, person)
	return err
}

func DeletePerson(appContext appengine.Context, person *model.Person) error {
	personKey := datastore.NewKey(appContext, "Person", person.Name, 0, nil)

	return datastore.Delete(appContext, personKey)
}