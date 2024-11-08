package models

import (
	"gorm.io/datatypes"
)

type FeedProductCustomData struct {
	ID            int32
	FeedProductId int32
	Data          datatypes.JSONMap
	DataIa        datatypes.JSONMap
}
