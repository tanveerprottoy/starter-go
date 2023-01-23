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

func CheckTypeInt(x any) bool {
	switch x.(type) {
	case int, int32, int64 :
		return true
	default:
		return false
	}
}

func CheckTypeFloat(x any) bool {
	switch x.(type) {
	case float32, float64 :
		return true
	default:
		return false
	}
}

func TypeCastIfInt(x any) (int32, error) {
	switch v := x.(type) {
	case int, int32, int64 :
		return v.(int32), nil
	default:
		return -1, errors.New("TypeCastIfInt error")
	}
}

func TypeCastIfFloat(x any) (float32, error) {
	switch v := x.(type) {
	case float32, float64 :
		return v.(float32), nil
	default:
		return -1, errors.New("TypeCastIfFloat error")
	}
}
