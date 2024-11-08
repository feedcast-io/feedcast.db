package models

import (
	"database/sql"
	"time"
)

type FeedScoreDate struct {
	ID     int32
	FeedId int32
	Date   time.Time

	FeedcastScore    sql.NullFloat64
	GoogleScore      sql.NullFloat64
	BingScore        sql.NullFloat64
	MetaScore        sql.NullFloat64
	FreeListingScore sql.NullFloat64
}
