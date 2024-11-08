package models

import "github.com/feedcast-io/feedcast.db/types"

type InvoicePack struct {
	ID                 int32
	Code               types.InvoiceProduct `gorm:"size:32,unique"`
	Name               string               `gorm:"size:255"`
	ProductId          string
	AdditionalPlatform int32
	MaxProducts        int32
	Prices             []InvoicePackPrice
	Visible            bool
}
