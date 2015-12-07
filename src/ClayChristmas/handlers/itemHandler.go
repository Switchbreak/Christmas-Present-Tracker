package handlers

import (
	"net/http"

	"appengine"
	"appengine/user"
	"github.com/gorilla/mux"

	"ClayChristmas/data"
	"ClayChristmas/model"
)

func GetWishlistItems(w http.ResponseWriter, r *http.Request) {
	appContext := appengine.NewContext(r)
	vars := mux.Vars(r)

	if checkInvited(appContext, vars["partyID"]) {
		wishlist, err := data.GetWishlist(appContext, vars["partyID"], vars["personID"])
		if err != nil {
			panic(err)
		}

		respond(w, wishlist)
	} else {
		http.Error(w, "Access Denied", http.StatusForbidden)
	}
}

func CreateWishlistItem(w http.ResponseWriter, r *http.Request) {
	appContext := appengine.NewContext(r)
	vars := mux.Vars(r)

	if checkInvited(appContext, vars["partyID"]) {
		personID, err := data.GetLoggedInPerson(appContext)
		if err != nil {
			panic(err)
		}
		
		if vars["personID"] != personID {
			person, err := data.GetPerson(appContext, vars["personID"])
			if err != nil {
				panic(err)
			}
			
			if person.Registered {
				http.Error(w, "Access Denied", http.StatusForbidden)
				return
			}
		}

		var wishlistItem model.WishlistItem
		getRequestObject(r, &wishlistItem)
		
		newItem, err := data.CreateWishlistItem(appContext, vars["partyID"], vars["personID"], &wishlistItem)
		if err != nil {
			panic(err)
		}

		respond(w, newItem)
	} else {
		http.Error(w, "Access Denied", http.StatusForbidden)
	}
}

func GetWishlistItem(w http.ResponseWriter, r *http.Request) {
	appContext := appengine.NewContext(r)
	vars := mux.Vars(r)

	if checkInvited(appContext, vars["partyID"]) {
		wishlistItem, err := data.GetWishlistItem(appContext, vars["partyID"], vars["personID"], vars["id"])
		if err != nil {
			panic(err)
		}

		respond(w, wishlistItem)
	} else {
		http.Error(w, "Access Denied", http.StatusForbidden)
	}
}

func UpdateWishlistItem(w http.ResponseWriter, r *http.Request) {
	appContext := appengine.NewContext(r)
	vars := mux.Vars(r)

	personID, err := data.GetLoggedInPerson(appContext)
	if err != nil {
		panic(err)
	}

	if user.IsAdmin(appContext) || vars["personID"] == personID {
		var wishlistItem model.WishlistItem
		getRequestObject(r, &wishlistItem)

		wishlistItem.ID = vars["id"]

		if _, err := data.UpdateWishlistItem(appContext, vars["partyID"], vars["personID"], &wishlistItem); err != nil {
			panic(err)
		}
	} else {
		http.Error(w, "Access Denied", http.StatusForbidden)
	}
}

func DeleteWishlistItem(w http.ResponseWriter, r *http.Request) {
	appContext := appengine.NewContext(r)
	vars := mux.Vars(r)

	personID, err := data.GetLoggedInPerson(appContext)
	if err != nil {
		panic(err)
	}

	if user.IsAdmin(appContext) || vars["personID"] == personID {
		var wishlistItem model.WishlistItem
		wishlistItem.ID = vars["id"]

		if err := data.DeleteWishlistItem(appContext, vars["partyID"], vars["personID"], &wishlistItem); err != nil {
			panic(err)
		}
	} else {
		http.Error(w, "Access Denied", http.StatusForbidden)
	}
}

func GetBoughtItems(w http.ResponseWriter, r *http.Request) {
	appContext := appengine.NewContext(r)
	vars := mux.Vars(r)

	if checkInvited(appContext, vars["partyID"]) {
		personID, err := data.GetLoggedInPerson(appContext)
		if err != nil {
			panic(err)
		}

		if personID != vars["personID"] {
			items, err := data.GetBoughtItems(appContext, vars["partyID"], vars["personID"])
			if err != nil {
				panic(err)
			}
	
			respond(w, items)
		} else {
			http.Error(w, "Access Denied", http.StatusForbidden)
		}
	} else {
		http.Error(w, "Access Denied", http.StatusForbidden)
	}
}

func CreateBoughtItem(w http.ResponseWriter, r *http.Request) {
	appContext := appengine.NewContext(r)
	vars := mux.Vars(r)

	if checkInvited(appContext, vars["partyID"]) {
		personID, err := data.GetLoggedInPerson(appContext)
		if err != nil {
			panic(err)
		}

		var boughtItem model.BoughtItem
		getRequestObject(r, &boughtItem)
		
		boughtItem.BoughtBy = personID
		
		newItem, err := data.CreateBoughtItem(appContext, vars["partyID"], vars["personID"], &boughtItem)
		if err != nil {
			panic(err)
		}

		respond(w, newItem)
	} else {
		http.Error(w, "Access Denied", http.StatusForbidden)
	}
}

func GetBoughtItem(w http.ResponseWriter, r *http.Request) {
	appContext := appengine.NewContext(r)
	vars := mux.Vars(r)

	if checkInvited(appContext, vars["partyID"]) {
		personID, err := data.GetLoggedInPerson(appContext)
		if err != nil {
			panic(err)
		}

		if personID != vars["personID"] {
			boughtItem, err := data.GetBoughtItem(appContext, vars["partyID"], vars["personID"], vars["id"])
			if err != nil {
				panic(err)
			}
	
			respond(w, boughtItem)
		} else {
			http.Error(w, "Access Denied", http.StatusForbidden)
		}
	} else {
		http.Error(w, "Access Denied", http.StatusForbidden)
	}
}

func UpdateBoughtItem(w http.ResponseWriter, r *http.Request) {
	appContext := appengine.NewContext(r)
	vars := mux.Vars(r)

	personID, err := data.GetLoggedInPerson(appContext)
	if err != nil {
		panic(err)
	}
	
	checkBoughtItem, err := data.GetBoughtItem( appContext, vars["partyID"], vars["personID"], vars["id"] )
	if err != nil {
		panic( err )
	}

	if user.IsAdmin(appContext) || checkBoughtItem.BoughtBy == personID {
		var boughtItem model.BoughtItem
		getRequestObject(r, &boughtItem)

		boughtItem.ID = checkBoughtItem.ID
		boughtItem.BoughtBy = checkBoughtItem.BoughtBy

		if _, err := data.UpdateBoughtItem(appContext, vars["partyID"], vars["personID"], &boughtItem); err != nil {
			panic(err)
		}
	} else {
		http.Error(w, "Access Denied", http.StatusForbidden)
	}
}

func DeleteBoughtItem(w http.ResponseWriter, r *http.Request) {
	appContext := appengine.NewContext(r)
	vars := mux.Vars(r)

	personID, err := data.GetLoggedInPerson(appContext)
	if err != nil {
		panic(err)
	}
	
	checkBoughtItem, err := data.GetBoughtItem( appContext, vars["partyID"], vars["personID"], vars["id"] )
	if err != nil {
		panic( err )
	}

	if user.IsAdmin(appContext) || checkBoughtItem.BoughtBy == personID {
		var boughtItem model.BoughtItem
		boughtItem.ID = vars["id"]

		if err := data.DeleteBoughtItem(appContext, vars["partyID"], vars["personID"], &boughtItem); err != nil {
			panic(err)
		}
	} else {
		http.Error(w, "Access Denied", http.StatusForbidden)
	}
}