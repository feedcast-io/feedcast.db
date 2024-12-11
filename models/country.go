package models

import (
	"database/sql"
	"gorm.io/datatypes"
)

type Country struct {
	ID                int32
	Code              string         `gorm:"size:2,unique"`
	Name              string         `gorm:"size:255"`
	Zone              sql.NullString `gorm:"size:16"`
	PhonePrefix       sql.NullString `gorm:"size:16"`
	Vats              *[]CountryVat
	Timezones         datatypes.JSONSlice[string]
	DefaultCurrency   *Currency
	DefaultCurrencyId sql.NullInt32
}
