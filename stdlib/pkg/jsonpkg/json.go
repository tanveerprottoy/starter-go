package jsonpkg

import (
	"encoding/json"
	"io"
)

// Marshal returns the JSON encoding of v
func Marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

// Unmarshal parses the JSON-encoded data and stores the result in the value pointed to by v
// The caller must pass the address for the v any param, ex: &v
func Unmarshal(d []byte, v any) error {
	return json.Unmarshal(d, v)
}

// Encode Encode writes the JSON encoding of v to the stream which is provided by the encoder created from the passed io.writer
func Encode(w io.Writer, v any) error {
	return json.NewEncoder(w).Encode(v)
}

// Decode reads the JSON value from decoder created from the passed io.reader
// The caller must pass the address for the v any param, ex: &v
func Decode(r io.Reader, v any) error {
	return json.NewDecoder(r).Decode(v)
}
