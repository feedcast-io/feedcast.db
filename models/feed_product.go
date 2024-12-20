package models

import (
	"cmp"
	"database/sql"
	"fmt"
	"github.com/feedcast-io/feedcast.db/types"
	"gorm.io/gorm"
	"regexp"
	"sort"
	"strings"
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

func (from *FeedProduct) ToGoogleProduct() (*types.GoogleProduct, error) {
	p := &types.GoogleProduct{}

	if from.ReferenceId.Valid && nil == from.Reference {
		return nil, fmt.Errorf("product reference relationship is missing")
	}

	p.Reference = from.Reference.Reference

	p.Condition = fmt.Sprintf("%s", from.ConditionId)
	p.Available = fmt.Sprintf("%s", from.AvailabilityId)
	p.AgeGroup = fmt.Sprintf("%s", from.AgeGroupId)
	p.Gender = fmt.Sprintf("%s", from.GenderId)

	if nil == from.Feed.Country {
		return nil, fmt.Errorf("product feed country relationship is missing")
	}
	if nil == from.Feed.Language {
		return nil, fmt.Errorf("product feed language relationship is missing")
	}
	if from.CurrencyId.Valid && nil == from.Currency {
		return nil, fmt.Errorf("product currency relationship is missing")
	}

	p.Language = from.Feed.Language.Code
	p.Country = from.Feed.Country.Code

	if from.IsBundle.Bool {
		p.Bundle = "yes"
	}

	if from.CategoryId.Valid {
		if nil == from.Category {
			return nil, fmt.Errorf("product category relationship is missing")
		}
		p.Category = fmt.Sprintf("%d", from.Category.GoogleId)
	}

	if from.ProductBrandId.Valid {
		if nil == from.ProductBrand {
			return nil, fmt.Errorf("product brand relationship is missing")
		}
		p.Brand = from.ProductBrand.Name
	}

	if nil != from.Text {
		p.Title = from.Text.Title.String
		p.Description = from.Text.Description.String
		p.Label0 = from.Text.Label0.String
		p.Label1 = from.Text.Label1.String
		p.Label2 = from.Text.Label3.String
		p.Label3 = from.Text.Label3.String
		p.Label4 = from.Text.Label4.String
		p.ProductType = from.Text.ProductType.String
	}

	if nil != from.Url {
		p.Link = from.Url.Link.String
		p.Image = from.Url.ImageLink.String
		p.AdsRedirect = from.Url.AdsRedirect.String
		p.MobileLink = from.Url.MobileLink.String
	}

	if nil != from.Various {
		p.AvailableDate = from.Various.AvailableDate.String
		p.Color = from.Various.Color.String
		p.Size = from.Various.Size.String
		p.Mpn = from.Various.Mpn.String
		p.ItemGroupId = from.Various.ItemGroupId.String
		p.Material = from.Various.Material.String

		re := regexp.MustCompile("^[0-9]{8,14}$")
		if re.Match([]byte(from.Various.Gtin.String)) {
			p.Gtin = from.Various.Gtin.String
		}
	}

	if from.Shipping != nil {
		if from.Shipping.ProductWeight.Int32 > 0 {
			p.ProductWeight = fmt.Sprintf("%.3f %s", float32(from.Shipping.ProductWeight.Int32)/100.0, from.Shipping.ProductWeightUnit)
		}
		if from.Shipping.ShippingWeight.Int32 > 0 {
			p.ShippingWeight = fmt.Sprintf("%.3f %s", float32(from.Shipping.ShippingWeight.Int32)/100.0, from.Shipping.ShippingWeightUnit)
		}

		if from.Shipping.ShippingValue.Valid {
			p.Shipping = &types.GoogleShipping{
				Country: p.Country,
				Price:   fmt.Sprintf("%.2f %s", float32(from.Shipping.ShippingValue.Int32)/100.0, from.Currency.Code),
			}
		}
	}

	// Use product weight as shipping weight if missing
	p.ShippingWeight = cmp.Or(p.ShippingWeight, p.ProductWeight)

	var prices []float64

	if from.Price.Int32 > 0 {
		prices = append(prices, float64(from.Price.Int32)/100.0)
	}

	if from.SalePrice.Int32 > 0 {
		prices = append(prices, float64(from.SalePrice.Int32)/100.0)
	}

	if len(prices) > 0 {
		sort.Float64s(prices)
		p.Price = fmt.Sprintf("%.2f %s", prices[len(prices)-1], from.Currency.Code)
		if 2 == len(prices) {
			p.SalePrice = fmt.Sprintf("%.2f %s", prices[0], from.Currency.Code)
		}
	}

	if len(p.Gtin) == 0 && (len(p.Mpn) == 0 || len(p.Brand) == 0) {
		p.IdExists = "false"
	}

	if nil != from.CustomData {
		p.CustomData = from.CustomData.Data
		p.CustomDataAi = from.CustomData.DataIa
	}

	return p, nil
}

// Return feed products ID from feedID & array of references as map with lowered-case references as key
func GetFeedProductsByReferences(conn *gorm.DB, feedId int32, references []string) (map[string]types.FeedProductIdentifier, error) {
	sql := `
SELECT
    fp.id AS feed_product_id, fpr.id AS reference_id, fpr.reference, fpt.id AS text_id, fpu.id AS url_id,
    fps.id AS shipping_id, fpv.id AS various_id
FROM
    feed_product_reference fpr
    LEFT JOIN feed_product fp ON (fp.reference_id = fpr.id AND fp.feed_id = ?)
    LEFT JOIN feed_product_text fpt ON fp.id = fpt.feed_product_id
    LEFT JOIN feed_product_url fpu ON fp.id = fpu.feed_product_id
    LEFT JOIN feed_product_shipping fps ON fp.id = fps.feed_product_id
    LEFT JOIN feed_product_various fpv ON fp.id = fpv.feed_product_id
WHERE
    fpr.reference IN ?`

	var rows []types.FeedProductIdentifier
	result := make(map[string]types.FeedProductIdentifier)
	res := conn.Unscoped().Raw(sql, feedId, references).Scan(&rows)

	for _, row := range rows {
		result[strings.ToLower(row.Reference)] = row
	}

	return result, res.Error
}
