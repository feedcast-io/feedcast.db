package models

import (
	feedcast_database "github.com/feedcast-io/feedcast.db"
	"testing"
)

func TestProductErrorModel(t *testing.T) {
	conn := feedcast_database.GetConnection()
	defer conn.Close()

	var productErrors []ProductError
	if e := conn.Gorm.
		Where(ProductError{ToDisplay: true}).
		Limit(100).Last(&productErrors).Error; e != nil {
		t.Fatal(e)
	}

	if 0 == len(productErrors) {
		t.Error("no product error found")
	}
}
