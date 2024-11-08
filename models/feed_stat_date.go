package models

import (
	"database/sql"
	"time"
)

type FeedStatDate struct {
	ID     int32
	FeedId int32
	Date   time.Time

	DateStartLastImport sql.NullTime
	DateEndLastImport   sql.NullTime

	Imported sql.NullInt32
	Errors   sql.NullInt32
	Warnings sql.NullInt32
	Pendings sql.NullInt32
}
