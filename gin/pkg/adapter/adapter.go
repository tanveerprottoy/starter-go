package adapter

import (
	"errors"
	"io"
	"log"
	"reflect"
	"strconv"
	"time"

	"github.com/tanveerprottoy/starter-go/pkg/json"
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

func InterfaceToStruct[T any](inter interface{}) (T, error) {
	s, ok := inter.(T)
	if ok {
		return s, errors.New("TypeCast error")
	}
	return s, nil
}

func StringToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

func StringToFloat(s string, bitSize int) (float64, error) {
	return strconv.ParseFloat(s, bitSize)
}

func ValuesToStruct[T any](params []any, t *T) {
	value := reflect.Indirect(
		reflect.ValueOf(t).Elem(),
	)
	for i := 0; i < value.NumField(); i++ {
		f := value.Field(i)
		if f.CanSet() {
			param := params[i]
			switch f.Kind() {
			case reflect.String:
				f.SetString(
					reflect.ValueOf(param).Elem().Interface().(string),
				)
			case reflect.Int32, reflect.Int64:
				f.SetInt(reflect.ValueOf(param).Elem().Interface().(int64))
			case reflect.Float32, reflect.Float64:
				f.SetFloat(reflect.ValueOf(param).Elem().Interface().(float64))
			case reflect.Bool:
				f.SetBool(reflect.ValueOf(param).Elem().Interface().(bool))
			case reflect.Struct:
				// currently only handle time.Time type
				f.Set(reflect.ValueOf(
					reflect.ValueOf(param).Elem().Interface().(time.Time),
				))
			default:
				log.Println("type unknown")
			}
		}
	}
}
