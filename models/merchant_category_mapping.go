package models

import "database/sql"

type MerchantCategoryMapping struct {
	ID int32

	Merchant   Merchant
	MerchantId int32
	Feed       Feed
	FeedId     int32
	Category   *ProductCategory
	CategoryId sql.NullInt32

	Hash          string
	OriginalValue string
	IsAuto        bool
}
