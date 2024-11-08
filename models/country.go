package models

import (
	"database/sql"
	"github.com/feedcast-io/feedcast.db/types"
)

type Country struct {
	ID                int32
	Code              string         `gorm:"size:2,unique"`
	Name              string         `gorm:"size:255"`
	Zone              sql.NullString `gorm:"size:16"`
	PhonePrefix       sql.NullString `gorm:"size:16"`
	Vats              *[]CountryVat
	Timezones         types.ArrayString
	DefaultCurrency   *Currency
	DefaultCurrencyId sql.NullInt32
}
