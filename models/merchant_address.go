package models

import "database/sql"

type MerchantAddress struct {
	ID           int32
	MerchantId   int32
	Address      sql.NullString `gorm:"size:255"`
	Address2     sql.NullString `gorm:"size:255"`
	City         sql.NullString `gorm:"size:255"`
	Zip          sql.NullString `gorm:"size:255"`
	Country      *Country
	CountryId    sql.NullInt32
	PhoneNumber  sql.NullString `gorm:"size:50"`
	InvoiceEmail sql.NullString `gorm:"size:255"`
	VatNumber    sql.NullString `gorm:"size:50"`
	VatType      sql.NullString `gorm:"size:10"`
	Siret        sql.NullString `gorm:"size:20"`
	VatExempted  bool
}
