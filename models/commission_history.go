package models

import (
	"github.com/feedcast-io/feedcast.db/types"
	"time"
)

type CommissionHistory struct {
	ID            int32
	Invoice       Invoice
	InvoiceId     int32
	Reseller      Reseller
	ResellerId    int32
	Feed          Feed
	FeedId        int32
	Date          time.Time
	Type          types.CommissionHistoryType
	Amount        int64
	Available     bool
	InvoiceCall   *InvoiceCall
	InvoiceCallId int32
}
