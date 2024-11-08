package models

import "database/sql"

type ResellerAddress struct {
	ID           int32
	Reseller     Reseller
	ResellerID   int32
	Country      *Country
	CountryId    sql.NullInt32
	Address      sql.NullString `gorm:"size:255"`
	Address2     sql.NullString `gorm:"size:255"`
	City         sql.NullString `gorm:"size:255"`
	Zip          sql.NullString `gorm:"size:255"`
	VatNumber    sql.NullString `gorm:"size:10"`
	VatType      sql.NullString `gorm:"size:10"`
	PhoneNumber  sql.NullString `gorm:"size:50"`
	InvoiceEmail sql.NullString `gorm:"size:255"`
	Siret        sql.NullString `gorm:"size:20"`
	VatExempted  bool
}
