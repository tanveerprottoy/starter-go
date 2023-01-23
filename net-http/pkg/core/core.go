package core

import (
	"errors"
	"fmt"
	"reflect"
)

func InterfaceToStruct[T any](inter interface{}) T {
	s, ok := inter.(T)
	if ok {
		fmt.Println(s)
	}
	return s
}

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
