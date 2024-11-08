package models

import (
	"github.com/feedcast-io/feedcast.db/types"
	"gorm.io/gorm"
)

type User struct {
	ID             int32
	Roles          types.ArrayString `gorm:"type:text"`
	DeletedAt      *gorm.DeletedAt
	Authentication Authentication
}
