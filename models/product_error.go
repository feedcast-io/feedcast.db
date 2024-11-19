package models

import "database/sql"

type ProductError struct {
	ID             int32
	Description    string `gorm:"size:1024"`
	IsBlocking     bool
	NeedUserAction bool
	Hash           sql.NullString `gorm:"size:32,unique"`
	Attribute      sql.NullString `gorm:"size:32,unique"`
	Code           sql.NullString `gorm:"size:255"`
	ToDisplay      bool
}
