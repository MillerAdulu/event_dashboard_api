package utils

import "encoding/json"

// Unwrap -
func Unwrap(wrapped []byte) map[string]interface{} {
	var u map[string]interface{}

	_ = json.Unmarshal(wrapped, &u)

	return u
}
