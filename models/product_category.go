package models

import (
	"database/sql"
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
