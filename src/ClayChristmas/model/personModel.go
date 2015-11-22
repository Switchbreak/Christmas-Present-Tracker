package model

import (
	"time"
	
	"appengine/user"
)

type Person struct {
	User		user.User `json:"-"`
	Name		string
	LastLogin	time.Time
}