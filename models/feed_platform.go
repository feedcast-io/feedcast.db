package models

import "database/sql"

type FeedPlatform struct {
	ID     int32
	FeedId int32

	AdsAccount    *Account
	AdsAccountId  sql.NullInt32
	GoogleEnabled bool

	ShoppingAccount   *Account
	ShoppingAccountId sql.NullInt32

	BingAccount   *Account
	BingAccountId sql.NullInt32
	BingEnabled   bool

	BingMerchantAccount   *Account
	BingMerchantAccountId sql.NullInt32

	MetaAccount   *Account
	MetaAccountId sql.NullInt32
	MetaEnabled   bool
}
