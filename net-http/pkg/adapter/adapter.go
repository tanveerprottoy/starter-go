package adapter

import (
	"io"
	"strconv"

	"github.com/tanveerprottoy/rest-api-starter-go/net-http/pkg/json"
)

func IOReaderToBytes(r io.Reader) ([]byte, error) {
	b, err := io.ReadAll(r)
	return b, err
}

func BytesToType[T any](b []byte) (*T, error) {
	var out T
	err := json.Unmarshal(b, &out)
	return &out, err
}

func BodyToType[T any](b io.ReadCloser) (*T, error) {
	var out T
	err := json.Decode(b, &out)
	if err != nil {
		return nil, err
	}
	return AnyToType[T](out)
}

func AnyToType[T any](v any) (*T, error) {
	var out T
	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &out)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func StringToInt(s string) (int, error) {
	return strconv.Atoi(s)
}
