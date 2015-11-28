package model

import (
	"time"
)

type Comment struct {
	ID		string `datastore:"-"`
	Comment	string
	Author	Person
	Private	bool
	Date	time.Time
}