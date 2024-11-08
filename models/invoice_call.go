package models

import (
	"github.com/feedcast-io/feedcast.db/types"
	"time"
)

type InvoiceCall struct {
	ID         int32
	Reseller   Reseller
	ResellerId int32
	Status     types.InvoiceCallStatus
	Code       string `gorm:"size:16,index:unique"`
	Hash       string `gorm:"size:32,index:unique"`
	Amount     int64
	Tax        int64
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
