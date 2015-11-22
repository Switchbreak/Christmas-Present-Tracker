package model

import (
	"time"
)

type Comment struct {
	Comment	string
	Author	Person
	Private	bool
	Date	time.Time
}