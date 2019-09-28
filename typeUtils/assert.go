package typeUtils

import (
	"reflect"
)

func IsArray(v interface{}) bool {
	rt := reflect.TypeOf(v)
	switch rt.Kind() {
	case reflect.Array:
		return true
	default:
		return false
	}
}

func IsSlice(v interface{}) bool {
	rt := reflect.TypeOf(v)
	switch rt.Kind() {
	case reflect.Slice:
		return true
	default:
		return false
	}
}

func IsArrayOrSlice(v interface{}) bool {
	rt := reflect.TypeOf(v)
	switch rt.Kind() {
	case reflect.Slice:
		return true
	case reflect.Array:
		return true
	default:
		return false
	}
}
