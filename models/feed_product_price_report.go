package models

import "database/sql"

type FeedProductPriceReport struct {
	ID                 int32
	FeedProductId      int32
	BenchmarkPrice     sql.NullInt64
	BenchmarkLastCheck sql.NullTime
}
