package models

import (
	"database/sql"
	"github.com/feedcast-io/feedcast.db/types"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"
)

type Log struct {
	ID        int32
	Date      time.Time
	LogType   types.LogTypes
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

func AddFeedLog(conn *gorm.DB, logType types.LogTypes, feed *Feed, data map[string]interface{}) (*Log, error) {
	var merchantUser MerchantUser

	if err := conn.
		Where("merchant_id = ?", feed.MerchantId).
		Order("id ASC").
		First(&merchantUser).
		Error; err != nil {
		return nil, err
	}

	log := Log{
		FeedId:         sql.NullInt32{feed.ID, true},
		MerchantUserId: sql.NullInt32{merchantUser.ID, true},
		Date:           time.Now(),
		LogType:        logType,
		Data:           data,
	}

	return &log, conn.Create(&log).Error
}
