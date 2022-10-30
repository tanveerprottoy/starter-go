package coreutil

import (
	"encoding/json"
	"net/http"
)

func Marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

func Unmarshal(data []byte, spec any) error {
	return json.Unmarshal(data, spec)
}

func Decode(r *http.Request, v any) error {
	return json.NewDecoder(r.Body).Decode(&v)
}
