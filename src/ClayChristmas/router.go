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

	router.HandleFunc("/api/party/",						handlers.GetParties).Methods("GET")
	router.HandleFunc("/api/party/{id}",					handlers.GetParty).Methods("GET")
	router.HandleFunc("/api/party/{id}",					handlers.CreateParty).Methods("POST")
	router.HandleFunc("/api/party/{id}",					handlers.UpdateParty).Methods("PUT")
	router.HandleFunc("/api/party/{id}",					handlers.DeleteParty).Methods("DELETE")
	router.HandleFunc("/api/party/{partyID}/invited",		handlers.GetInvitedPeople).Methods("GET")
	router.HandleFunc("/api/party/{id}/invite/{personID}",	handlers.GetInvitedPeople).Methods("POST")
	
	router.HandleFunc("/api/person/",						handlers.GetPeople).Methods("GET")
	router.HandleFunc("/api/person/{id}",					handlers.CreatePerson).Methods("POST")
	router.HandleFunc("/api/person/{id}",					handlers.GetPerson).Methods("GET")
	router.HandleFunc("/api/person/{id}",					handlers.UpdatePerson).Methods("PUT")
	router.HandleFunc("/api/person/{id}",					handlers.DeletePerson).Methods("DELETE")
	router.HandleFunc("/api/person/{id}/link",				handlers.LinkPerson).Methods("POST")
	
	router.HandleFunc("/api/markLogin",						handlers.MarkLogin).Methods("POST")

	http.Handle("/", router)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world")
}
