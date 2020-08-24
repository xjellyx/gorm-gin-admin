package body

import "github.com/mitchellh/mapstructure"

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
