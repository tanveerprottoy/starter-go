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

func CastTypeIfInt(x any) int {
	switch v := x.(type) {
	case int, int32, int64 :
		return v.(int)
	default:
		return -1
	}
}

func CastTypeIfFloat(x any) float64 {
	switch v := x.(type) {
	case float32, float64 :
		return v.(float64)
	default:
		return -1
	}
}
