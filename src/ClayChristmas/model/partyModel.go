package model

import (
	"time"
)

type Party struct {
	Title			string
	Description		string `datastore:,noindex`
	Date			time.Time
	CreatedBy		string
	CreatedDate		time.Time
}