package body

import (
	"database/sql/driver"
	"encoding/json"
	"github.com/mitchellh/mapstructure"
)

type Body map[string]interface{}

// Mapstructure
func (b Body) Mapstructure(input interface{}) error {
	return mapstructure.Decode(input, &b)
}

// Delete
func (b Body) Delete(key string) Body {
	delete(b, key)
	return b
}

func (b Body) Value() (driver.Value, error) {
	return json.Marshal(b)
}

func (b *Body) Scan(d interface{}) (err error) {
	return json.Unmarshal(d.([]byte), b)
}