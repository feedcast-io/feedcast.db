package models

import (
	"database/sql"
	"time"
)

type Invoice struct {
	ID             int32
	Code           string `gorm:"size:32"`
	Amount         int64
	Tax            sql.NullInt64
	Currency       *Currency
	CurrencyId     sql.NullInt32
	InvoiceId      string `gorm:"size:64"`
	Created        time.Time
	DueDate        sql.NullTime
	AmountPayed    int64
	Status         string `gorm:"size:16"`
	Url            string `gorm:"size:500"`
	InvoiceNote    *InvoiceNote
	Subscription   *Subscription
	SubscriptionId sql.NullInt32
}
