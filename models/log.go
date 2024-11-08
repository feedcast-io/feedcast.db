package models

import (
	"database/sql"
	"gorm.io/datatypes"
	"time"
)

type Log struct {
	ID        int32
	Date      time.Time
	LogType   int16
	Data      datatypes.JSONMap
	AddressIp sql.NullString
	Seen      bool

	User   *User
	UserId sql.NullInt32

	Feed   *Feed
	FeedId sql.NullInt32

	MerchantUser   *MerchantUser
	MerchantUserId sql.NullInt32

	Reseller   *Reseller
	ResellerId sql.NullInt32

	Commercial   *Commercial
	CommercialId sql.NullInt32

	Account   *Account
	AccountId sql.NullInt32

	Guest   *Authentication
	GuestId sql.NullInt32
}
