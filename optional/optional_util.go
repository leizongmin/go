package optional

import (
	"reflect"
)

// 如果值为nil则返回None
func OfNullable(v interface{}) Optional {
	if v == nil {
		return none
	}
	return Some(v)
}

// 如果值为nil、0、空字符串、零长度的array、slice、map则返回None
func OfZeroable(v interface{}) Optional {
	if v == nil {
		return none
	}
	switch reflect.TypeOf(v).Kind().String() {
	case "int":
		if v.(int) == 0 {
			return none
		}
	case "int8":
		if v.(int8) == 0 {
			return none
		}
	case "int16":
		if v.(int16) == 0 {
			return none
		}
	case "int32":
		if v.(int32) == 0 {
			return none
		}
	case "int64":
		if v.(int64) == 0 {
			return none
		}
	case "uint":
		if v.(uint) == 0 {
			return none
		}
	case "uint8":
		if v.(uint8) == 0 {
			return none
		}
	case "uint16":
		if v.(uint16) == 0 {
			return none
		}
	case "uint32":
		if v.(uint32) == 0 {
			return none
		}
	case "uint64":
		if v.(uint64) == 0 {
			return none
		}
	case "float32":
		if v.(float32) == 0 {
			return none
		}
	case "float64":
		if v.(float64) == 0 {
			return none
		}
	case "rune":
		if v.(rune) == 0 {
			return none
		}
	case "string":
		if len(v.(string)) < 1 {
			return none
		}
	case "slice":
		if reflect.ValueOf(v).Len() < 1 {
			return none
		}
	case "array":
		if reflect.ValueOf(v).Len() < 1 {
			return none
		}
	case "map":
		if reflect.ValueOf(v).Len() < 1 {
			return none
		}
	default:
	}
	return Some(v)
}
