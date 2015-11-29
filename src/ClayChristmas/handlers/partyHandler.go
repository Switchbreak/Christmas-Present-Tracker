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

func GetParties(w http.ResponseWriter, r *http.Request) {
	appContext := appengine.NewContext(r)

	personID, err := data.GetLoggedInPerson(appContext)
	if err != nil {
		panic(err)
	}

	parties, err := data.GetPartiesInvited(appContext, personID)
	if err != nil {
		panic(err)
	}

	respond(w, parties)
}

func CreateParty(w http.ResponseWriter, r *http.Request) {
	appContext := appengine.NewContext(r)
	vars := mux.Vars(r)

	if _, err := data.GetParty(appContext, vars["id"]); err == nil {
		http.Error(w, "Party title already taken", http.StatusForbidden)
	} else if err != data.ErrNoSuchEntity {
		panic(err)
	}

	personID, err := data.GetLoggedInPerson(appContext)
	if err != nil {
		panic(err)
	}

	var party model.Party
	getRequestObject(r, &party)

	party.Title = vars["id"]
	party.CreatedBy = personID
	party.CreatedDate = time.Now()

	if err := data.UpdateParty(appContext, &party); err != nil {
		panic(err)
	}
}

// TODO: Invite people

func GetParty(w http.ResponseWriter, r *http.Request) {
	appContext := appengine.NewContext(r)
	vars := mux.Vars(r)

	if checkInvited(appContext, vars["id"]) {
		party, err := data.GetParty(appContext, vars["id"])
		if err != nil {
			panic(err)
		}

		respond(w, party)
	} else {
		http.Error(w, "Access Denied", http.StatusForbidden)
	}
}

func UpdateParty(w http.ResponseWriter, r *http.Request) {
	appContext := appengine.NewContext(r)
	vars := mux.Vars(r)

	checkParty, err := data.GetParty(appContext, vars["id"])
	if err != nil {
		panic(err)
	}

	if checkPartyOwner(appContext, checkParty) {
		var party model.Party
		getRequestObject(r, &party)

		party.Title = vars["id"]
		party.CreatedBy = checkParty.CreatedBy
		party.CreatedDate = checkParty.CreatedDate

		if err := data.UpdateParty(appContext, &party); err != nil {
			panic(err)
		}
	} else {
		http.Error(w, "Access Denied", http.StatusForbidden)
	}
}

func DeleteParty(w http.ResponseWriter, r *http.Request) {
	appContext := appengine.NewContext(r)
	vars := mux.Vars(r)

	checkParty, err := data.GetParty(appContext, vars["id"])
	if err != nil {
		panic(err)
	}

	if checkPartyOwner(appContext, checkParty) {
		if err := data.DeleteParty(appContext, checkParty); err != nil {
			panic(err)
		}
	} else {
		http.Error(w, "Access Denied", http.StatusForbidden)
	}
}

func checkInvited(appContext appengine.Context, partyID string) bool {
	personID, err := data.GetLoggedInPerson(appContext)
	if err != nil {
		panic(err)
	}

	isInvited, err := data.IsInvited(appContext, partyID, personID)
	if err != nil {
		panic(err)
	}

	return isInvited
}

func checkPartyOwner(appContext appengine.Context, checkParty *model.Party) bool {
	if !user.IsAdmin(appContext) {
		personID, err := data.GetLoggedInPerson(appContext)
		if err != nil {
			panic(err)
		}

		if checkParty.CreatedBy != personID {
			return false
		}
	}

	return true
}
