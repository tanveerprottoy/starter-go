package adapter

import (
	"io"
	"strconv"

	"txp/restapistarter/pkg/json"
)

func IOReaderToBytes(r io.Reader) ([]byte, error) {
	b, err := io.ReadAll(r)
	return b, err
}

func BytesToValue[T any](b []byte) (*T, error) {
	var out T
	err := json.Unmarshal(b, &out)
	return &out, err
}

func AnyToValue[T any](v any) (*T, error) {
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

func AnyToValue[T any](v any) (*T, error) {
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
