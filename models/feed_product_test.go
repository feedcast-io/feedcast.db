package models

import (
	feedcast_database "github.com/feedcast-io/feedcast.db"
	"strings"
	"testing"
)

func TestFeedProduct_ToGoogleProduct(t *testing.T) {
	conn := feedcast_database.GetConnection()
	defer conn.Close()

	var products []FeedProduct
	if e := conn.Gorm.
		Preload("Reference").
		Preload("Feed.Country").
		Preload("Feed.Language").
		Preload("Category").
		Preload("ProductBrand").
		Preload("Currency").
		Preload("Text").
		Preload("Url").
		Preload("Various").
		Preload("CustomData").
		Joins("INNER JOIN feed_product_custom_data ON feed_product.id = feed_product_custom_data.feed_product_id AND LENGTH(feed_product_custom_data.data) > 10").
		Limit(1000).
		Where("reference_id > 0").
		Last(&products).Error; e != nil {
		t.Error(e)
	}

	if 0 == len(products) {
		t.Error("no products found")
	}

	for _, product := range products {
		gp, e := product.ToGoogleProduct()
		gp.AppendCustomData()

		if nil != e {
			t.Error(e)
		} else {
			if len(gp.Brand) == 0 {
				t.Error()
			}
		}

		if 0 == len(gp.Title) {
			t.Error("no title found")
		}
		if 0 == len(gp.Description) {
			t.Error("no description found")
		}
		if 0 == len(gp.Price) {
			t.Error("no price found")
		}
		if !strings.Contains(gp.Link, "https://") {
			t.Error("invalid link")
		}
		if !strings.Contains(gp.Image, "https://") {
			t.Error("invalid image link")
		}

		customTitle := ""
		if t, ok := product.CustomData.DataIa["title"].(string); ok {
			customTitle = t
		}
		if t, ok := product.CustomData.Data["title"].(string); ok {
			customTitle = t
		}

		if len(customTitle) > 0 {
			if customTitle == product.Text.Title.String {
				t.Error("custom title same as origin title")
			}
			if gp.Title != customTitle {
				t.Errorf("override title error. Expected: %s, got: %s", customTitle, gp.Title)
			}
		}
	}
}
