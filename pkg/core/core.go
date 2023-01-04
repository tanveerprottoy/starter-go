package core

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"time"
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
