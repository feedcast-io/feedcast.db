package models

import (
	"database/sql"
	"time"
)

type Subscription struct {
	ID               int32
	Code             string `gorm:"size:50"`
	Created          time.Time
	Status           string `gorm:"size:32"`
	DateStartPeriod  time.Time
	DateEndPeriod    time.Time
	DateCancel       sql.NullTime
	EndTrial         sql.NullTime
	Merchant         *Merchant
	MerchantId       sql.NullInt32
	Reseller         *Reseller
	ResellerId       sql.NullInt32
	Complete         bool
	IsChargeAuto     bool
	LastInvoicePayed bool
	LastInvoice      *Invoice
	LastInvoiceId    sql.NullInt32 `gorm:"column:lastest_invoice_id"`
}
