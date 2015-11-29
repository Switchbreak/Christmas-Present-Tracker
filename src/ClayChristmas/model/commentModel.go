package model

import (
	"time"
)

type Comment struct {
	ID		string `datastore:"-"`
	Comment	string
	Author	string
	Private	bool
	Date	time.Time
}