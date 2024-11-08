package models

import "database/sql"

type InvoiceNote struct {
	ID           int32
	Invoice      *Invoice
	InvoiceId    sql.NullInt32
	Comment      string `gorm:"size:1024"`
	PaymentError string `gorm:"size:1024"`
}
