package model

import (
	"time"
	
	"appengine/user"
)

type Person struct {
	User		user.User
	Name		string
	LastLogin	time.Time
}