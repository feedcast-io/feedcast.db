package models

import (
	feedcast_database "github.com/feedcast-io/feedcast.db"
	"testing"
)

func TestProductCategory_GetTitleFromLang(t *testing.T) {
	conn := feedcast_database.GetConnection()
	defer conn.Close()

	var pct []ProductCategoryText

	if e := conn.Gorm.
		Limit(100).
		Preload("ProductCategory").
		Preload("ProductCategory.Texts").
		Where(ProductCategoryText{Lang: "FR"}).
		Find(&pct).Error; nil != e {
		t.Error(e)
	}

	for _, p := range pct {
		if title := p.ProductCategory.GetTitleFromLang("FR"); title != p.Title {
			t.Errorf("wrong title. Expected '%s', got '%s'", p.Title, title)
		}
	}
}
