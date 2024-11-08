package feedcast_database

import (
	"strings"
	"testing"
)

func TestGetConnection(t *testing.T) {
	conn := GetConnection()
	defer conn.Close()

	var version string

	if e := conn.Gorm.Raw("SELECT @@version").Scan(&version).Error; nil != e {
		t.Errorf("Get server version from raw query failed: %s", e.Error())
	}

	expected := "mariadb"

	if lowered := strings.ToLower(version); !strings.Contains(lowered, expected) {
		t.Errorf("Unexpected server version: %s (expected content: %s)", lowered, expected)
	}
}
