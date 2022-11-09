package json

import (
	"encoding/json"
	"io"
)

func Marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

func Unmarshal(data []byte, spec any) error {
	return json.Unmarshal(data, spec)
}

func Decode(r io.Reader, v any) error {
	return json.NewDecoder(r).Decode(&v)
}
