package utils

import (
	"fmt"
	"reflect"
)

func pluck(list interface{}, fieldName string, defaultVal interface{}) interface{} {
	reflectVal := reflect.ValueOf(list)
	switch reflectVal.Kind() {
	case reflect.Array, reflect.Slice:
		if reflectVal.Len() == 0 {
			return defaultVal
		}

		valElem := reflectVal.Type().Elem()
		for valElem.Kind() == reflect.Ptr {
			valElem = valElem.Elem()
		}

		if valElem.Kind() != reflect.Struct {
			panic("list element is not struct")
		}

		field, ok := valElem.FieldByName(fieldName)
		if !ok {
			panic(fmt.Sprintf("field %s not found", fieldName))
		}

		result := reflect.MakeSlice(reflect.SliceOf(field.Type), reflectVal.Len(), reflectVal.Len())

		for i := 0; i < reflectVal.Len(); i++ {
			ev := reflectVal.Index(i)
			for ev.Kind() == reflect.Ptr {
				ev = ev.Elem()
			}
			if ev.Kind() != reflect.Struct {
				panic("element is not a struct")
			}
			if !ev.IsValid() {
				continue
			}
			result.Index(i).Set(ev.FieldByIndex(field.Index))
		}

		return result.Interface()
	default:
		panic("list must be an array or slice")
	}
}

func PluckUint64(list interface{}, fieldName string) []uint64 {
	return pluck(list, fieldName, []uint64{}).([]uint64)
}
