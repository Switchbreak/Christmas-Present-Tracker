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

func GetPartyComments(w http.ResponseWriter, r *http.Request) {
	appContext := appengine.NewContext(r)
	vars := mux.Vars(r)

	if checkInvited(appContext, vars["partyID"]) {
		comments, err := data.GetComments(appContext, vars["partyID"], "")
		if err != nil {
			panic(err)
		}

		respond(w, comments)
	} else {
		http.Error(w, "Access Denied", http.StatusForbidden)
	}
}

func GetPersonComments(w http.ResponseWriter, r *http.Request) {
	appContext := appengine.NewContext(r)
	vars := mux.Vars(r)

	if checkInvited(appContext, vars["partyID"]) {
		personID, err := data.GetLoggedInPerson(appContext)
		if err != nil {
			panic(err)
		}

		if personID != vars["personID"] {
			comments, err := data.GetComments(appContext, vars["partyID"], vars["personID"])
			if err != nil {
				panic(err)
			}

			respond(w, comments)
		} else {
			http.Error(w, "Access Denied", http.StatusForbidden)
		}
	} else {
		http.Error(w, "Access Denied", http.StatusForbidden)
	}
}

func CreateComment(w http.ResponseWriter, r *http.Request) {
	appContext := appengine.NewContext(r)
	vars := mux.Vars(r)

	if checkInvited(appContext, vars["partyID"]) {
		personID, err := data.GetLoggedInPerson(appContext)
		if err != nil {
			panic(err)
		}

		if personID != vars["personID"] {
			var comment model.Comment
			getRequestObject(r, &comment)

			comment.Date = time.Now()
			comment.Author = personID

			newComment, err := data.CreateComment(appContext, vars["partyID"], vars["personID"], &comment)
			if err != nil {
				panic(err)
			}

			respond(w, newComment)
		} else {
			http.Error(w, "Access Denied", http.StatusForbidden)
		}
	} else {
		http.Error(w, "Access Denied", http.StatusForbidden)
	}
}

func GetPartyComment(w http.ResponseWriter, r *http.Request) {
	appContext := appengine.NewContext(r)
	vars := mux.Vars(r)

	if checkInvited(appContext, vars["partyID"]) {
		comment, err := data.GetComment(appContext, vars["partyID"], "", vars["id"])
		if err != nil {
			panic(err)
		}

		respond(w, comment)
	} else {
		http.Error(w, "Access Denied", http.StatusForbidden)
	}
}

func GetPersonComment(w http.ResponseWriter, r *http.Request) {
	appContext := appengine.NewContext(r)
	vars := mux.Vars(r)

	if checkInvited(appContext, vars["partyID"]) {
		personID, err := data.GetLoggedInPerson(appContext)
		if err != nil {
			panic(err)
		}

		if personID != vars["personID"] {
			comment, err := data.GetComment(appContext, vars["partyID"], vars["personID"], vars["id"])
			if err != nil {
				panic(err)
			}

			respond(w, comment)
		} else {
			http.Error(w, "Access Denied", http.StatusForbidden)
		}
	} else {
		http.Error(w, "Access Denied", http.StatusForbidden)
	}
}

func UpdateComment(w http.ResponseWriter, r *http.Request) {
	appContext := appengine.NewContext(r)
	vars := mux.Vars(r)

	checkComment, err := data.GetComment(appContext, vars["partyID"], vars["personID"], vars["id"])
	if err != nil {
		panic(err)
	}

	if checkCommentOwner(appContext, checkComment) {
		var comment model.Comment
		getRequestObject(r, &comment)

		comment.ID = vars["id"]
		comment.Author = checkComment.Author
		comment.Date = checkComment.Date

		if _, err := data.UpdateComment(appContext, vars["partyID"], vars["personID"], &comment); err != nil {
			panic(err)
		}
	} else {
		http.Error(w, "Access Denied", http.StatusForbidden)
	}
}

func DeleteComment(w http.ResponseWriter, r *http.Request) {
	appContext := appengine.NewContext(r)
	vars := mux.Vars(r)

	checkComment, err := data.GetComment(appContext, vars["partyID"], vars["personID"], vars["id"])
	if err != nil {
		panic(err)
	}

	if checkCommentOwner(appContext, checkComment) {
		if err := data.DeleteComment(appContext, vars["partyID"], vars["personID"], checkComment); err != nil {
			panic(err)
		}
	} else {
		http.Error(w, "Access Denied", http.StatusForbidden)
	}
}

func checkCommentOwner(appContext appengine.Context, checkComment *model.Comment) bool {
	if !user.IsAdmin(appContext) {
		personID, err := data.GetLoggedInPerson(appContext)
		if err != nil {
			panic(err)
		}

		if checkComment.Author != personID {
			return false
		}
	}

	return true
}
