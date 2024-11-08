package models

import (
	"database/sql"
	"github.com/feedcast-io/feedcast.db/types"
	"slices"
	"time"
)

type Reseller struct {
	ID                int32
	Name              string `gorm:"size:50"`
	Authentication    Authentication
	Merchants         []Merchant
	CreatedAt         time.Time
	UpdatedAt         time.Time
	StripeCustomerId  sql.NullString `gorm:"size:32"`
	HasInvoicePayment bool
	Type              types.ResellerType
	Commission        *ResellerCommission
	Address           *ResellerAddress
	DefaultLanguage   string
}

func (r *Reseller) IsNonPayer() bool {
	return slices.Contains([]types.ResellerType{
		types.ResellerTypeNonPayer,
		types.ResellerTypeNonPayerWithCommission,
	}, r.Type)
}
