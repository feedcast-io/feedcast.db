package models

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type Merchant struct {
	ID                int32
	Users             []MerchantUser
	Name              string `gorm:"size:100"`
	DeletedAt         *gorm.DeletedAt
	Reseller          *Reseller
	ResellerId        sql.NullInt32
	Commercial        *Commercial
	CommercialId      sql.NullInt32
	CreatedAt         time.Time
	UpdatedAt         time.Time
	Feeds             []Feed
	Address           *MerchantAddress
	StripeCustomerId  sql.NullString `gorm:"size:32"`
	HasInvoicePayment bool
	DefaultLanguage   string
	HearAboutUs       sql.NullInt16
	IsSpamSuspicion   sql.NullBool
}

func (m *Merchant) IsDirectCustomer() bool {
	return !m.ResellerId.Valid || m.Reseller.IsNonPayer()
}
