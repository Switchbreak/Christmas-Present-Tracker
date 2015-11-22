package model

import (
	"time"
	
	"appengine/datastore"
)

type Party struct {
	Title			string
	Description		string `datastore:,noindex`
	Date			time.Time
	CreatedBy		string
	CreatedDate		time.Time
}

type InvitedPerson struct {
	Person			*datastore.Key
}