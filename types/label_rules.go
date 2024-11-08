package types

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type LabelRules struct {
	Rules []struct {
		LabelValue string `json:"label_value"`
		LabelKey   string `json:"label_key"`
		Overwrite  bool   `json:"overwrite"`
		IsAny      bool   `json:"condition_any"`
		Conditions []struct {
			Field    string `json:"field"`
			Operator string `json:"operator"`
			Value    string `json:"value"`
		} `json:"conditions"`
	} `json:"rules"`
}

func (r *LabelRules) Value() (driver.Value, error) {
	return json.Marshal(r)
}

func (r *LabelRules) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	var rules LabelRules
	err := json.Unmarshal(bytes, &rules)

	if nil == err {
		*r = rules
	}

	// Silent error if unable to decode, map will just be empty
	return nil
}
