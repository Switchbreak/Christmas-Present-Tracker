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
	
	router.HandleFunc("/api/party/{id}/invite/{personID}",	handlers.InvitePerson).Methods("POST")
	router.HandleFunc("/api/party/{partyID}/invited",		handlers.GetInvitedPeople).Methods("GET")
	router.HandleFunc("/api/party/{partyID}/comment",		handlers.GetPartyComments).Methods("GET")
	router.HandleFunc("/api/party/{partyID}/comment",		handlers.CreateComment).Methods("POST")
	router.HandleFunc("/api/party/{partyID}/comment/{id}",	handlers.GetPartyComment).Methods("GET")
	router.HandleFunc("/api/party/{partyID}/comment/{id}",	handlers.UpdateComment).Methods("PUT")
	router.HandleFunc("/api/party/{partyID}/comment/{id}",	handlers.DeleteComment).Methods("DELETE")
	router.HandleFunc("/api/party/{partyID}/{personID}/comment",		handlers.GetPersonComments).Methods("GET")
	router.HandleFunc("/api/party/{partyID}/{personID}/comment",		handlers.CreateComment).Methods("POST")
	router.HandleFunc("/api/party/{partyID}/{personID}/comment/{id}",	handlers.GetPersonComment).Methods("GET")
	router.HandleFunc("/api/party/{partyID}/{personID}/comment/{id}",	handlers.UpdateComment).Methods("PUT")
	router.HandleFunc("/api/party/{partyID}/{personID}/comment/{id}",	handlers.DeleteComment).Methods("DELETE")
	
	router.HandleFunc("/api/party/{partyID}/{personID}/wishlist",		handlers.GetWishlistItems).Methods("GET")
	router.HandleFunc("/api/party/{partyID}/{personID}/wishlist",		handlers.CreateWishlistItem).Methods("POST")
	router.HandleFunc("/api/party/{partyID}/{personID}/wishlist/{id}",	handlers.GetWishlistItem).Methods("GET")
	router.HandleFunc("/api/party/{partyID}/{personID}/wishlist/{id}",	handlers.UpdateWishlistItem).Methods("PUT")
	router.HandleFunc("/api/party/{partyID}/{personID}/wishlist/{id}",	handlers.DeleteWishlistItem).Methods("DELETE")
	
	router.HandleFunc("/api/party/{partyID}/{personID}/bought",			handlers.GetBoughtItems).Methods("GET")
	router.HandleFunc("/api/party/{partyID}/{personID}/bought",			handlers.CreateBoughtItem).Methods("POST")
	router.HandleFunc("/api/party/{partyID}/{personID}/bought/{id}",	handlers.GetBoughtItem).Methods("GET")
	router.HandleFunc("/api/party/{partyID}/{personID}/bought/{id}",	handlers.UpdateBoughtItem).Methods("PUT")
	router.HandleFunc("/api/party/{partyID}/{personID}/bought/{id}",	handlers.DeleteBoughtItem).Methods("DELETE")
	
	router.HandleFunc("/api/party/{id}",					handlers.GetParty).Methods("GET")
	router.HandleFunc("/api/party/{id}",					handlers.CreateParty).Methods("POST")
	router.HandleFunc("/api/party/{id}",					handlers.UpdateParty).Methods("PUT")
	router.HandleFunc("/api/party/{id}",					handlers.DeleteParty).Methods("DELETE")
	
	router.HandleFunc("/api/person/",						handlers.GetPeople).Methods("GET")
	router.HandleFunc("/api/person/current",				handlers.GetLoggedInPerson).Methods("GET")
	router.HandleFunc("/api/person/{id}/link",				handlers.LinkPerson).Methods("POST")
	router.HandleFunc("/api/person/{id}",					handlers.CreatePerson).Methods("POST")
	router.HandleFunc("/api/person/{id}",					handlers.GetPerson).Methods("GET")
	router.HandleFunc("/api/person/{id}",					handlers.UpdatePerson).Methods("PUT")
	router.HandleFunc("/api/person/{id}",					handlers.DeletePerson).Methods("DELETE")
	
	router.HandleFunc("/api/markLogin",						handlers.MarkLogin).Methods("POST")
	router.HandleFunc("/api/sign-out",						handlers.LogoutScreen).Queries("returnURL", "{returnURL}")

	http.Handle("/", router)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world")
}
