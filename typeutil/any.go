package typeutil

import (
	"fmt"
	"reflect"
	"time"
)

type AnyType struct {
	v interface{}
}

// 任意类型
func Any(v interface{}) AnyType {
	return AnyType{v: v}
}

// 获取interface值
func (a AnyType) Value() interface{} {
	return a.v
}

// 获取指定索引的值（适用于 array, slice, map）
func (a AnyType) Get(key interface{}) (AnyType, bool) {
	v := reflect.ValueOf(a.v)
	k := reflect.ValueOf(key)
	vt := v.Type()
	kt := k.Type()
	switch vt.Kind().String() {
	case "map":
		if vt.Key().String() != kt.String() {
			return Any(nil), false
		}
		ret := v.MapIndex(k)
		if !ret.IsValid() {
			return Any(nil), false
		}
		return Any(ret.Interface()), true
	case "slice":
		if kt.String() != "int" {
			return Any(nil), false
		}
		i := k.Interface().(int)
		if v.Len() <= i {
			return Any(nil), false
		}
		ret := v.Index(i)
		if !ret.IsValid() {
			return Any(nil), false
		}
		return Any(ret.Interface()), true
	case "array":
		if kt.String() != "int" {
			return Any(nil), false
		}
		i := k.Interface().(int)
		if v.Len() <= i {
			return Any(nil), false
		}
		ret := v.Index(i)
		if !ret.IsValid() {
			return Any(nil), false
		}
		return Any(ret.Interface()), true
	default:
		return Any(nil), false
	}
}

// 获取指定索引的值（适用于 array, slice, map），如果不存在则报错
func (a AnyType) MustGet(key interface{}) AnyType {
	ret, ok := a.Get(key)
	if !ok {
		panic(fmt.Sprintf(`index or key "%s" does not in value %+v`, key, a))
	}
	return ret
}

func (a AnyType) ToBool(defaultValue bool) bool {
	if v, ok := a.v.(bool); ok {
		return v
	} else {
		return defaultValue
	}
}

func (a AnyType) MustToBool() bool {
	if v, ok := a.v.(bool); ok {
		return v
	} else {
		panic(fmt.Sprintf("failed to cast %+v to bool", a.v))
	}
}

func (a AnyType) ToInt(defaultValue int) int {
	if v, ok := a.v.(int); ok {
		return v
	} else {
		return defaultValue
	}
}

func (a AnyType) MustToInt() int {
	if v, ok := a.v.(int); ok {
		return v
	} else {
		panic(fmt.Sprintf("failed to cast %+v to int", a.v))
	}
}

func (a AnyType) MustToInt8() int8 {
	if v, ok := a.v.(int8); ok {
		return v
	} else {
		panic(fmt.Sprintf("failed to cast %+v to int8", a.v))
	}
}

func (a AnyType) ToInt8(defaultValue int8) int8 {
	if v, ok := a.v.(int8); ok {
		return v
	} else {
		return defaultValue
	}
}

func (a AnyType) ToInt16(defaultValue int16) int16 {
	if v, ok := a.v.(int16); ok {
		return v
	} else {
		return defaultValue
	}
}

func (a AnyType) MustToInt16() int16 {
	if v, ok := a.v.(int16); ok {
		return v
	} else {
		panic(fmt.Sprintf("failed to cast %+v to int16", a.v))
	}
}

func (a AnyType) ToInt32(defaultValue int32) int32 {
	if v, ok := a.v.(int32); ok {
		return v
	} else {
		return defaultValue
	}
}

func (a AnyType) MustToInt32() int32 {
	if v, ok := a.v.(int32); ok {
		return v
	} else {
		panic(fmt.Sprintf("failed to cast %+v to int32", a.v))
	}
}

func (a AnyType) ToInt64(defaultValue int64) int64 {
	if v, ok := a.v.(int64); ok {
		return v
	} else {
		return defaultValue
	}
}

func (a AnyType) MustToInt64() int64 {
	if v, ok := a.v.(int64); ok {
		return v
	} else {
		panic(fmt.Sprintf("failed to cast %+v to int64", a.v))
	}
}

func (a AnyType) ToUint(defaultValue uint) uint {
	if v, ok := a.v.(uint); ok {
		return v
	} else {
		return defaultValue
	}
}

func (a AnyType) MustToUint() uint {
	if v, ok := a.v.(uint); ok {
		return v
	} else {
		panic(fmt.Sprintf("failed to cast %+v to uint", a.v))
	}
}

func (a AnyType) ToUint8(defaultValue uint8) uint8 {
	if v, ok := a.v.(uint8); ok {
		return v
	} else {
		return defaultValue
	}
}

func (a AnyType) MustToUint8() uint8 {
	if v, ok := a.v.(uint8); ok {
		return v
	} else {
		panic(fmt.Sprintf("failed to cast %+v to uint8", a.v))
	}
}

func (a AnyType) ToUint16(defaultValue uint16) uint16 {
	if v, ok := a.v.(uint16); ok {
		return v
	} else {
		return defaultValue
	}
}

func (a AnyType) MustToUint16() uint16 {
	if v, ok := a.v.(uint16); ok {
		return v
	} else {
		panic(fmt.Sprintf("failed to cast %+v to uint16", a.v))
	}
}

func (a AnyType) ToUint32(defaultValue uint32) uint32 {
	if v, ok := a.v.(uint32); ok {
		return v
	} else {
		return defaultValue
	}
}

func (a AnyType) MustToUint32() uint32 {
	if v, ok := a.v.(uint32); ok {
		return v
	} else {
		panic(fmt.Sprintf("failed to cast %+v to uint32", a.v))
	}
}

func (a AnyType) ToUint64(defaultValue uint64) uint64 {
	if v, ok := a.v.(uint64); ok {
		return v
	} else {
		return defaultValue
	}
}

func (a AnyType) MustToUint64() uint64 {
	if v, ok := a.v.(uint64); ok {
		return v
	} else {
		panic(fmt.Sprintf("failed to cast %+v to uint64", a.v))
	}
}

func (a AnyType) ToFloat32(defaultValue float32) float32 {
	if v, ok := a.v.(float32); ok {
		return v
	} else {
		return defaultValue
	}
}

func (a AnyType) MustToFloat32() float32 {
	if v, ok := a.v.(float32); ok {
		return v
	} else {
		panic(fmt.Sprintf("failed to cast %+v to float32", a.v))
	}
}

func (a AnyType) ToFloat64(defaultValue float64) float64 {
	if v, ok := a.v.(float64); ok {
		return v
	} else {
		return defaultValue
	}
}

func (a AnyType) MustToFloat64() float64 {
	if v, ok := a.v.(float64); ok {
		return v
	} else {
		panic(fmt.Sprintf("failed to cast %+v to float64", a.v))
	}
}

func (a AnyType) ToString(defaultValue string) string {
	if v, ok := a.v.(string); ok {
		return v
	} else {
		return defaultValue
	}
}

func (a AnyType) MustToString() string {
	if v, ok := a.v.(string); ok {
		return v
	} else {
		panic(fmt.Sprintf("failed to cast %+v to string", a.v))
	}
}

func (a AnyType) ToMap(defaultValue map[string]interface{}) map[string]interface{} {
	if v, ok := a.v.(map[string]interface{}); ok {
		return v
	} else {
		return defaultValue
	}
}

func (a AnyType) MustToMap() map[string]interface{} {
	if v, ok := a.v.(map[string]interface{}); ok {
		return v
	} else {
		panic(fmt.Sprintf("failed to cast %+v to map[string]interface{}", a.v))
	}
}

func (a AnyType) ToTime(defaultValue time.Time) time.Time {
	if v, ok := a.v.(time.Time); ok {
		return v
	} else {
		return defaultValue
	}
}

func (a AnyType) MustToTime() time.Time {
	if v, ok := a.v.(time.Time); ok {
		return v
	} else {
		panic(fmt.Sprintf("failed to cast %+v to time.Time", a.v))
	}
}

func (a AnyType) ToComplex64(defaultValue complex64) complex64 {
	if v, ok := a.v.(complex64); ok {
		return v
	} else {
		return defaultValue
	}
}

func (a AnyType) MustToComplex64() complex64 {
	if v, ok := a.v.(complex64); ok {
		return v
	} else {
		panic(fmt.Sprintf("failed to cast %+v to complex64", a.v))
	}
}

func (a AnyType) ToComplex128(defaultValue complex128) complex128 {
	if v, ok := a.v.(complex128); ok {
		return v
	} else {
		return defaultValue
	}
}

func (a AnyType) MustToComplex128() complex128 {
	if v, ok := a.v.(complex128); ok {
		return v
	} else {
		panic(fmt.Sprintf("failed to cast %+v to complex128", a.v))
	}
}

func (a AnyType) ToByte(defaultValue byte) byte {
	if v, ok := a.v.(byte); ok {
		return v
	} else {
		return defaultValue
	}
}

func (a AnyType) MustToByte() byte {
	if v, ok := a.v.(byte); ok {
		return v
	} else {
		panic(fmt.Sprintf("failed to cast %+v to byte", a.v))
	}
}

func (a AnyType) ToRune(defaultValue rune) rune {
	if v, ok := a.v.(rune); ok {
		return v
	} else {
		return defaultValue
	}
}

func (a AnyType) MustToRune() rune {
	if v, ok := a.v.(rune); ok {
		return v
	} else {
		panic(fmt.Sprintf("failed to cast %+v to rune", a.v))
	}
}

func (a AnyType) ToUintptr(defaultValue uintptr) uintptr {
	if v, ok := a.v.(uintptr); ok {
		return v
	} else {
		return defaultValue
	}
}

func (a AnyType) MustToUintptr() uintptr {
	if v, ok := a.v.(uintptr); ok {
		return v
	} else {
		panic(fmt.Sprintf("failed to cast %+v to uintptr", a.v))
	}
}
