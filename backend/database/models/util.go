package models

import (
	"database/sql/driver"
	"fmt"
)

type JSONField []byte

func (j *JSONField) Scan(value interface{}) error {
	if value == nil {
		*j = []byte("{}")
		return nil
	}

	switch v := value.(type) {
	case []byte:
		*j = v
	case string:
		*j = []byte(v)
	default:
		return fmt.Errorf("unsupported type: %T", value)
	}
	return nil
}

func (j JSONField) Value() (driver.Value, error) {
	if len(j) == 0 {
		return []byte("{}"), nil
	}
	return []byte(j), nil
}
