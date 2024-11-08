package models

import "database/sql"

type FeedSubscriptionHistory struct {
	ID int32

	Feed           *Feed
	FeedId         sql.NullInt32
	OldFeed        *Feed
	OldFeedId      sql.NullInt32
	Reseller       *Reseller
	ResellerId     sql.NullInt32
	Subscription   Subscription
	SubscriptionId int32

	DateStartSubscription sql.NullTime
	DateEndSubscription   sql.NullTime
	DateEndTrial          sql.NullTime
	IsCancelled           bool
}
