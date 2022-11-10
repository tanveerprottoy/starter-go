package adapter

import (
	"io"
	"txp/restapistarter/pkg/json"
)

func IOReaderToBytes(r io.Reader) ([]byte, error) {
	b, err := io.ReadAll(r)
	return b, err
}

func BytesToValue[T any](d []byte) (*T, error) {
	var out T
	err := json.Unmarshal(d, &out)
	return &out, err
}

func AnyToValue[T any](d any) (*T, error) {
	var out T
	b, err := json.Marshal(d)
	if err != nil {
		return &out, err
	}
	err = json.Unmarshal(b, &out)
	if err != nil {
		return nil, err
	}
	return &out, err
}
