package models

import (
	"database/sql"
	"github.com/feedcast-io/feedcast.db/types"
)

type FeedProductShipping struct {
	ID                 int32
	FeedProductId      int32
	ProductWeight      sql.NullInt32
	ProductWeightUnit  types.WeightUnit
	ShippingWeight     sql.NullInt32
	ShippingWeightUnit types.WeightUnit
	ShippingValue      sql.NullInt32
}
