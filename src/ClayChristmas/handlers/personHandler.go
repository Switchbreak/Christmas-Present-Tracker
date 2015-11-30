package handlers

import (
	"net/http"
	"time"

	"appengine"
	"appengine/user"
	"github.com/gorilla/mux"

	"ClayChristmas/data"
	"ClayChristmas/model"
)

func GetInvitedPeople(w http.ResponseWriter, r *http.Request) {
	appContext := appengine.NewContext(r)
	vars := mux.Vars(r)

	if checkInvited(appContext, vars["partyID"]) {
		people, err := data.GetInvitedPeople(appContext, vars["partyID"])
		if err != nil {
			panic(err)
		}

		respond(w, people)
	} else {
		http.Error(w, "Access Denied", http.StatusForbidden)
	}
}

func GetLoggedInPerson(w http.ResponseWriter, r *http.Request) {
	appContext := appengine.NewContext(r)

	personID, err := data.GetLoggedInPerson(appContext)
	if err == data.UserNotFound {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	} else if err != nil {
		panic(err)
	}

	person, err := data.GetPerson(appContext, personID)
	if err != nil {
		panic(err)
	}

	respond(w, person)
}

func GetPeople(w http.ResponseWriter, r *http.Request) {
	appContext := appengine.NewContext(r)

	people, err := data.GetPeople(appContext)
	if err != nil {
		panic(err)
	}

	respond(w, people)
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	appContext := appengine.NewContext(r)
	vars := mux.Vars(r)

	if _, err := data.GetPerson(appContext, vars["id"]); err == nil {
		http.Error(w, "Name already taken", http.StatusForbidden)
	} else if err != data.ErrNoSuchEntity {
		panic(err)
	}

	var person model.Person
	getRequestObject(r, &person)

	person.LastLogin = time.Now()

	if err := data.UpdatePerson(appContext, &person); err != nil {
		panic(err)
	}
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	appContext := appengine.NewContext(r)
	vars := mux.Vars(r)

	person, err := data.GetPerson(appContext, vars["id"])
	if err != nil {
		panic(err)
	}

	respond(w, person)
}

func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	appContext := appengine.NewContext(r)
	vars := mux.Vars(r)

	checkPerson, err := data.GetPerson(appContext, vars["id"])
	if err != nil {
		panic(err)
	}

	if checkPersonOwner(appContext, checkPerson) {
		var person model.Person
		getRequestObject(r, &person)

		person.Name = vars["id"]
		person.LastLogin = time.Now()

		if err := data.UpdatePerson(appContext, &person); err != nil {
			panic(err)
		}
	} else {
		http.Error(w, "Access Denied", http.StatusForbidden)
	}
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	appContext := appengine.NewContext(r)
	vars := mux.Vars(r)

	checkPerson, err := data.GetPerson(appContext, vars["id"])
	if err != nil {
		panic(err)
	}

	if checkPersonOwner(appContext, checkPerson) {
		if err := data.DeletePerson(appContext, checkPerson); err != nil {
			panic(err)
		}
	} else {
		http.Error(w, "Access Denied", http.StatusForbidden)
	}
}

func LinkPerson(w http.ResponseWriter, r *http.Request) {
	appContext := appengine.NewContext(r)
	vars := mux.Vars(r)

	if _, err := data.GetLoggedInPerson(appContext); err == data.UserNotFound {
		person, err := data.GetPerson(appContext, vars["id"])
		if err != nil {
			panic(err)
		}

		if person.User.ID != "" {
			http.Error(w, "Person already linked", http.StatusForbidden)
		}

		person.User = *user.Current(appContext)
		if err = data.UpdatePerson(appContext, person); err != nil {
			panic(err)
		}
	} else if err != nil {
		panic(err)
	} else {
		http.Error(w, "User already linked", http.StatusForbidden)
	}
}

func MarkLogin(w http.ResponseWriter, r *http.Request) {
	appContext := appengine.NewContext(r)

	personID, err := data.GetLoggedInPerson(appContext)
	if err != nil {
		panic(err)
	}

	person, err := data.GetPerson(appContext, personID)
	if err != nil {
		panic(err)
	}

	person.LastLogin = time.Now()
	if err := data.UpdatePerson(appContext, person); err != nil {
		panic(err)
	}
}

func checkPersonOwner(appContext appengine.Context, checkPerson *model.Person) bool {
	if !user.IsAdmin(appContext) {
		personID, err := data.GetLoggedInPerson(appContext)
		if err != nil {
			panic(err)
		}

		if checkPerson.Name != personID {
			return false
		}
	}

	return true
}
