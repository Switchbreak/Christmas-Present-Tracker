package handlers

import (
	"encoding/json"
	"net/http"
)

func getRequestObject(r *http.Request, v interface{}) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(v)
	if err != nil {
		panic(err)
	}
}

func respond(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		panic(err)
	}
}
