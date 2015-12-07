package handlers

import (
	"net/http"
	"net/url"
	
	"appengine"
	"appengine/user"
	"github.com/gorilla/mux"
)

func LogoutScreen( w http.ResponseWriter, r *http.Request ) {
	appContext := appengine.NewContext(r)
	vars := mux.Vars(r)
	
	returnURL, err := url.QueryUnescape( vars["returnURL"] );
	if err != nil {
		http.Error( w, err.Error(), http.StatusInternalServerError )
		return
	}
	
	url, err := user.LogoutURL( appContext, returnURL )
	if err != nil {
		http.Error( w, err.Error(), http.StatusInternalServerError )
		return
	}
	
	w.Header().Set( "Location", url )
	w.WriteHeader( http.StatusFound )
}