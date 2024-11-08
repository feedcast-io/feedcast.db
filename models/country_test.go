package models

import (
	feedcast_database "github.com/feedcast-io/feedcast.db"
	"slices"
	"testing"
)

func TestCountryModel(t *testing.T) {
	conn := feedcast_database.GetConnection()
	defer conn.Close()

	testData := map[string][]string{
		"FR": {"Europe/Paris"},
		"IT": {"Europe/Rome"},
		"ES": {"Africa/Ceuta", "Atlantic/Canary", "Europe/Madrid"},
	}

	for code, expectedTimezones := range testData {
		var country Country

		if e := conn.Gorm.Where("code = ?", code).Find(&country).Error; e != nil {
			t.Error(e)
		}

		if len(expectedTimezones) != len(country.Timezones) {
			t.Errorf("Country %s should have 1 timezone", country.Name)
		}

		for _, timezone := range expectedTimezones {
			if !slices.Contains(country.Timezones, timezone) {
				t.Errorf("Country %s should contain %s", country.Name, timezone)
			}
		}
	}
}
