package coreutil

import "encoding/json"

func Unmarshal(data []byte, spec interface{}) error {
	return json.Unmarshal(data, spec)
}
