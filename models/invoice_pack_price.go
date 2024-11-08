package models

type InvoicePackPrice struct {
	ID            int32
	Code          string `gorm:"size:32,unique"`
	Amount        int32
	Name          string `gorm:"size:255"`
	InvoicePack   InvoicePack
	InvoicePackId int32
	Active        bool
	IsDefault     bool
	IsMetered     bool
}
