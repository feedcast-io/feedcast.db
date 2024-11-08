package types

import "database/sql"

type FeedScore struct {
	Feedcast    sql.NullFloat64
	Google      sql.NullFloat64
	Meta        sql.NullFloat64
	Bing        sql.NullFloat64
	FreeListing sql.NullFloat64
}
