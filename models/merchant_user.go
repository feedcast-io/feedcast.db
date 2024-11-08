package models

import "github.com/feedcast-io/feedcast.db/types"

type MerchantUser struct {
	ID             int32
	FirstName      string `gorm:"size:64"`
	LastName       string `gorm:"size:64"`
	PhoneNumber    string `gorm:"size:16"`
	Authentication Authentication
	Roles          types.ArrayString `gorm:"type:text"`
	MerchantID     int32
	Merchant       *Merchant
	Credentials    []MerchantUserCredential
}
