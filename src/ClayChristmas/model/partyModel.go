package model

import (
	"time"
	
	"appengine/user"
	"appengine/datastore"
)

type Party struct {
	Title		string
	Description	string `datastore:,noindex`
	Date		time.Time
	Invited		[]Person
	CreatedBy	user.User
	CreatedDate	time.Time
}

type InvitedPerson struct {
	Person		datastore.Key
}