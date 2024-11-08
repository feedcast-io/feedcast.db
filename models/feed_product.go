package models

import (
	"database/sql"
	"github.com/feedcast-io/feedcast.db/types"
	"gorm.io/gorm"
)

type FeedProduct struct {
	ID                int32
	Feed              Feed `gorm:"index:feed_ref,priority:2"`
	FeedId            int32
	Reference         *FeedProductReference `gorm:"index:feed_ref,priority:1"`
	ReferenceId       sql.NullInt32
	CreatedAt         sql.NullTime
	UpdatedAt         sql.NullTime
	DeletedAt         *gorm.DeletedAt
	GoogleAdsStatus   int16 `gorm:"column:status"`
	BingStatus        int16
	FacebookStatus    int16
	FreeListingStatus int16
	AvailabilityId    types.Availability
	ConditionId       types.ConditionType
	AgeGroupId        types.AgeGroup
	GenderId          types.Gender
	Quantity          sql.NullInt32
	Currency          *Currency
	CurrencyId        sql.NullInt32
	Price             sql.NullInt32
	SalePrice         sql.NullInt32
	IsAdult           sql.NullBool
	IsBundle          sql.NullBool
	HasIdentifier     sql.NullBool
	ProductBrand      *ProductBrand
	ProductBrandId    sql.NullInt32 `gorm:"column:brand_id"`

	Category   *ProductCategory
	CategoryId sql.NullInt32

	Text        *FeedProductText
	Shipping    *FeedProductShipping
	Various     *FeedProductVarious
	CustomData  *FeedProductCustomData
	Url         *FeedProductUrl
	Feedback    []FeedProductFeedback
	PriceReport *FeedProductPriceReport
}
