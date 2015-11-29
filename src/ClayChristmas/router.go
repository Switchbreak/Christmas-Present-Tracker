package ClayChristmas

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"ClayChristmas/handlers"
)

func init() {
	router := mux.NewRouter()
	router.StrictSlash(true)

	router.HandleFunc("/api/", index).Methods("GET")

	router.HandleFunc("/api/party/",		handlers.GetParties).Methods("GET")
	router.HandleFunc("/api/party/{id}",	handlers.GetParty).Methods("GET")
	router.HandleFunc("/api/party/{id}",	handlers.UpdateParty).Methods("PUT")
	router.HandleFunc("/api/party/{id}",	handlers.DeleteParty).Methods("DELETE")

	router.HandleFunc("/api/party/{partyID}/invited",	handlers.GetInvitedPeople).Methods("GET")
	router.HandleFunc("/api/person/",					handlers.GetPeople).Methods("GET")
	router.HandleFunc("/api/person/{id}",				handlers.GetPerson).Methods("GET")
	router.HandleFunc("/api/person/{id}",				handlers.UpdatePerson).Methods("PUT")
	router.HandleFunc("/api/person/{id}",				handlers.DeletePerson).Methods("DELETE")

	http.Handle("/", router)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world")
}
