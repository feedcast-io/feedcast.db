package models

import "database/sql"

type FeedProductShipping struct {
	ID                 int32
	FeedProductId      int32
	ProductWeight      sql.NullInt32
	ProductWeightUnit  sql.NullInt16
	ShippingWeight     sql.NullInt32
	ShippingWeightUnit sql.NullInt16
	ShippingValue      sql.NullInt32
}
