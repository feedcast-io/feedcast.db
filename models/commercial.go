package models

import "database/sql"

type Commercial struct {
	ID             int32
	FirstName      string
	LastName       string
	Authentication Authentication
	Reseller       *Reseller
	ResellerId     sql.NullInt32
}
