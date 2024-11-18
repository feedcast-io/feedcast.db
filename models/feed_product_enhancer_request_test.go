package models

import (
	"encoding/json"
	feedcast_database "github.com/feedcast-io/feedcast.db"
	"slices"
	"testing"
)

func TestFeedProductEnhancerRequestModel(t *testing.T) {
	conn := feedcast_database.GetConnection()
	defer conn.Close()

	var results []FeedProductEnhancerRequest
	if e := conn.Gorm.Unscoped().
		Preload("Feed").
		Limit(100).
		Find(&results).Error; e != nil {
		t.Error(e)
	}

	if len(results) == 0 {
		t.Error("no results found")
	}

	for _, result := range results {
		if 0 == len(result.Products) {
			t.Error("no products found")
		}

		if result.Feed.ID != result.FeedId {
			t.Error("no feed found")
		}
	}
}

func TestFeedProductEnhancerRequestArrayOfScalar(t *testing.T) {
	conn := feedcast_database.GetConnection()
	defer conn.Close()

	tx := conn.Gorm.Begin()

	arrayOfInt := []int{
		128,
		256,
	}
	buf, _ := json.Marshal(arrayOfInt)

	res := tx.Exec("INSERT INTO feed_product_enhancer_request (feed_id, products, pending) VALUES (10956, ?, 0)", string(buf))

	if e := res.Error; nil != e {
		t.Error(e)
	}

	var req FeedProductEnhancerRequest

	if e := tx.Limit(1).Last(&req).Error; e != nil {
		t.Error(e)
	}

	for _, i := range arrayOfInt {
		if !slices.Contains(req.Products, i) {
			t.Error("Product id not found")
		}
	}

	if len(arrayOfInt) != len(req.Products) {
		t.Error("Number of products mismatch")
	}

	tx.Rollback()
}
