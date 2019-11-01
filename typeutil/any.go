package typeutil

import (
	"fmt"
	"reflect"
)

type AnyType struct {
	v interface{}
}

// 任意类型
func Any(v interface{}) AnyType {
	return AnyType{v: v}
}

// 获取interface值
func (a AnyType) Interface() interface{} {
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
