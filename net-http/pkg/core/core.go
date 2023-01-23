package core

import (
	"errors"
	"reflect"
)

func Dereference[T any](obj *T) (T, error) {
	if obj == nil {
		var retObj T
		return retObj, errors.New("input obj cannot be nil")
	}
	return *obj, nil
}

func ExtractFieldsFromStruct[T any](t *T) []any {
	v := reflect.Indirect(
		reflect.ValueOf(t).Elem(),
	)
	f := make([]any, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		f[0] = v.Field(i).Interface()
	}
	return f
}

func TypeCastIfInt(x any) int32 {
	switch v := x.(type) {
	case int, int32, int64 :
		return v.(int32)
	default:
		return -1
	}
}

func TypeCastIfFloat(x any) float32 {
	switch v := x.(type) {
	case float32, float64 :
		return v.(float32)
	default:
		return -1
	}
}
