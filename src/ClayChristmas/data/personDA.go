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

	query := datastore.NewQuery("Person").Filter("User.ID =", currentUser.ID).Limit(1).KeysOnly()
	keys, err := query.GetAll(appContext, nil)
	if err != nil {
		return nil, err
	}
	if len(keys) < 1 {
		return nil, errors.New("User not found")
	}
	return keys[0], nil
}

func GetPersonByKey(appContext appengine.Context, personKey *datastore.Key) (*model.Person, error) {
	var person model.Person
	if err := datastore.Get(appContext, personKey, &person); err != nil {
		return nil, err
	}

	return &person, nil
}

func GetPerson(appContext appengine.Context, name string) (*model.Person, error) {
	personKey := datastore.NewKey(appContext, "Person", name, 0, nil)
	return GetPersonByKey(appContext, personKey)
}

func UpdatePerson(appContext appengine.Context, person *model.Person) error {
	personKey := datastore.NewKey(appContext, "Person", person.Name, 0, nil)

	//	var checkPerson model.Person
	//	err := datastore.Get(appContext, personKey, &checkPerson)
	//	if err == nil {
	//		if err := checkPersonOwner(appContext, &checkPerson); err != nil {
	//			return err
	//		}
	//
	//		person.User = checkPerson.User
	//		person.LastLogin = checkPerson.LastLogin
	//	} else if err == datastore.ErrNoSuchEntity {
	//		person.User = *user.Current(appContext)
	//		person.LastLogin = time.Now()
	//	} else {
	//		return err
	//	}

	_, err := datastore.Put(appContext, personKey, &person)
	return err
}

func DeletePerson(appContext appengine.Context, person *model.Person) error {
	personKey := datastore.NewKey(appContext, "Person", person.Name, 0, nil)

	//	var checkPerson model.Person
	//	if err := datastore.Get(appContext, personKey, &checkPerson); err != nil {
	//		return err
	//	}

	return datastore.Delete(appContext, personKey)
}

//func checkPersonOwner(appContext appengine.Context, person *model.Person) error {
//	if !user.IsAdmin(appContext) {
//		currentUser := user.Current(appContext)
//		if currentUser.ID != person.User.ID {
//			return errors.New("User does not have permission")
//		}
//	}
//
//	return nil
//}