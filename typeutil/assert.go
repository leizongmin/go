package typeutil

import (
	"reflect"
)

func IsArray(value interface{}) bool {
	rt := reflect.TypeOf(value)
	switch rt.Kind() {
	case reflect.Array:
		return true
	default:
		return false
	}
}

func IsSlice(value interface{}) bool {
	rt := reflect.TypeOf(value)
	switch rt.Kind() {
	case reflect.Slice:
		return true
	default:
		return false
	}
}

func IsArrayOrSlice(value interface{}) bool {
	rt := reflect.TypeOf(value)
	switch rt.Kind() {
	case reflect.Slice:
		return true
	case reflect.Array:
		return true
	default:
		return false
	}
}
