package models

import (
	"database/sql"
	"gorm.io/gorm"
)

type SubscriptionItem struct {
	ID                 int32
	Feed               Feed
	FeedId             sql.NullInt32
	Subscription       Subscription
	SubscriptionId     int32
	InvoicePackPrice   InvoicePackPrice
	InvoicePackPriceId int32
	AutoRenewal        bool
	FeedOld            *Feed
	FeedOldId          sql.NullInt32
	DeletedAt          gorm.DeletedAt
}

func (si *SubscriptionItem) AfterDelete(tx *gorm.DB) error {
	if si.FeedId.Valid {
		return tx.Unscoped().
			Model(SubscriptionItem{}).
			Where("id = ?", si.ID).
			Updates(map[string]interface{}{
				"FeedId":      sql.NullInt32{},
				"FeedOldId":   si.FeedId,
				"AutoRenewal": false,
			}).Error
	}

	return nil
}
