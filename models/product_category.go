package models

import (
	"database/sql"
	"gorm.io/gorm"
	"strings"
)

type ProductCategory struct {
	ID       int32
	GoogleId int32
	Hash     string `gorm:"size:4"`
	Parent   *ProductCategory
	ParentId sql.NullInt32
	Texts    []ProductCategoryText
}

func (c *ProductCategory) GetTitleFromLang(lang string) string {
	title := ""

	for _, t := range c.Texts {
		if strings.ToLower(t.Lang) == strings.ToLower(lang) {
			title = t.Title
		}
	}

	return title
}

func GetHierarchyCategory(conn *gorm.DB, googleCategoryId int32) []ProductCategory {
	var categories []ProductCategory

	var ids []int32

	sql := `
WITH RECURSIVE an AS (
	SELECT pc.id, google_id, pc.parent_id
	FROM product_category pc
	WHERE pc.google_id = ?
	UNION
	SELECT pc.id, pc.google_id, pc.parent_id
	FROM product_category pc
	JOIN an ON an.parent_id = pc.id
)
SELECT google_id FROM an;
`
	conn.Raw(sql, googleCategoryId).Find(&ids)

	for _, id := range ids {
		var cat ProductCategory
		if e := conn.Find(&cat, id).Error; nil == e {
			categories = append(categories, cat)
		}
	}

	// reverse array
	for i, j := 0, len(categories)-1; i < j; i, j = i+1, j-1 {
		categories[i], categories[j] = categories[j], categories[i]
	}

	return categories
}
