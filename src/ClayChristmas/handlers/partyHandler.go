package handlers

import (
	"net/http"

	"appengine"
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

	var party model.Party
	getRequestObject(r, &party)
	party.Title = vars["id"]
	if err := data.UpdateParty(appContext, &party); err != nil {
		panic(err)
	}
}

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
