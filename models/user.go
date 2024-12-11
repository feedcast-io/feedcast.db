package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type User struct {
	ID             int32
	Roles          datatypes.JSONSlice[string] `gorm:"type:text"`
	DeletedAt      *gorm.DeletedAt
	Authentication Authentication
}
