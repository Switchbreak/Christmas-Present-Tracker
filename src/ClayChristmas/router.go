package ClayChristmas

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func init() {
	router := mux.NewRouter()
	router.StrictSlash(true)

	router.HandleFunc("/api/", index).Methods("GET")

	http.Handle("/", router)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world")
}
