package typeUtils

import (
	"fmt"
	"time"
)

func toInterfaceArray(value interface{}) ([]interface{}, bool) {
	ret, ok := value.([]interface{})
	return ret, ok
}

func ToBoolArray(value interface{}, defaultValue []bool) []bool {
	array, ok := toInterfaceArray(value)
	if !ok {
		return defaultValue
	}
	ret := make([]bool, len(array))
	for i, v := range array {
		v2, ok := v.(bool)
		if !ok {
			return defaultValue
		}
		ret[i] = v2
	}
	return ret
}

func MustToBoolArray(value interface{}) []bool {
	array, ok := toInterfaceArray(value)
	if !ok {
		panic(fmt.Sprintf("failed to cast %+v to []bool", value))
	}
	ret := make([]bool, len(array))
	for i, v := range array {
		v2, ok := v.(bool)
		if !ok {
			panic(fmt.Sprintf("failed to cast %+v to []bool", value))
		}
		ret[i] = v2
	}
	return ret
}

func ToIntArray(value interface{}, defaultValue []int) []int {
	array, ok := toInterfaceArray(value)
	if !ok {
		return defaultValue
	}
	ret := make([]int, len(array))
	for i, v := range array {
		v2, ok := v.(int)
		if !ok {
			return defaultValue
		}
		ret[i] = v2
	}
	return ret
}

func MustToIntArray(value interface{}) []int {
	array, ok := toInterfaceArray(value)
	if !ok {
		panic(fmt.Sprintf("failed to cast %+v to []bool", value))
	}
	ret := make([]int, len(array))
	for i, v := range array {
		v2, ok := v.(int)
		if !ok {
			panic(fmt.Sprintf("failed to cast %+v to []int", value))
		}
		ret[i] = v2
	}
	return ret
}

func MustToInt8Array(value interface{}) []int8 {
	array, ok := toInterfaceArray(value)
	if !ok {
		panic(fmt.Sprintf("failed to cast %+v to []bool", value))
	}
	ret := make([]int8, len(array))
	for i, v := range array {
		v2, ok := v.(int8)
		if !ok {
			panic(fmt.Sprintf("failed to cast %+v to []int8", value))
		}
		ret[i] = v2
	}
	return ret
}

func ToInt8Array(value interface{}, defaultValue []int8) []int8 {
	array, ok := toInterfaceArray(value)
	if !ok {
		return defaultValue
	}
	ret := make([]int8, len(array))
	for i, v := range array {
		v2, ok := v.(int8)
		if !ok {
			return defaultValue
		}
		ret[i] = v2
	}
	return ret
}

func ToInt16Array(value interface{}, defaultValue []int16) []int16 {
	array, ok := toInterfaceArray(value)
	if !ok {
		return defaultValue
	}
	ret := make([]int16, len(array))
	for i, v := range array {
		v2, ok := v.(int16)
		if !ok {
			return defaultValue
		}
		ret[i] = v2
	}
	return ret
}

func MustToInt16Array(value interface{}) []int16 {
	array, ok := toInterfaceArray(value)
	if !ok {
		panic(fmt.Sprintf("failed to cast %+v to []bool", value))
	}
	ret := make([]int16, len(array))
	for i, v := range array {
		v2, ok := v.(int16)
		if !ok {
			panic(fmt.Sprintf("failed to cast %+v to []int16", value))
		}
		ret[i] = v2
	}
	return ret
}

func ToInt32Array(value interface{}, defaultValue []int32) []int32 {
	array, ok := toInterfaceArray(value)
	if !ok {
		return defaultValue
	}
	ret := make([]int32, len(array))
	for i, v := range array {
		v2, ok := v.(int32)
		if !ok {
			return defaultValue
		}
		ret[i] = v2
	}
	return ret
}

func MustToInt32Array(value interface{}) []int32 {
	array, ok := toInterfaceArray(value)
	if !ok {
		panic(fmt.Sprintf("failed to cast %+v to []bool", value))
	}
	ret := make([]int32, len(array))
	for i, v := range array {
		v2, ok := v.(int32)
		if !ok {
			panic(fmt.Sprintf("failed to cast %+v to []int32", value))
		}
		ret[i] = v2
	}
	return ret
}

func ToInt64Array(value interface{}, defaultValue []int64) []int64 {
	array, ok := toInterfaceArray(value)
	if !ok {
		return defaultValue
	}
	ret := make([]int64, len(array))
	for i, v := range array {
		v2, ok := v.(int64)
		if !ok {
			return defaultValue
		}
		ret[i] = v2
	}
	return ret
}

func MustToInt64Array(value interface{}) []int64 {
	array, ok := toInterfaceArray(value)
	if !ok {
		panic(fmt.Sprintf("failed to cast %+v to []bool", value))
	}
	ret := make([]int64, len(array))
	for i, v := range array {
		v2, ok := v.(int64)
		if !ok {
			panic(fmt.Sprintf("failed to cast %+v to []int64", value))
		}
		ret[i] = v2
	}
	return ret
}

func ToUintArray(value interface{}, defaultValue []uint) []uint {
	array, ok := toInterfaceArray(value)
	if !ok {
		return defaultValue
	}
	ret := make([]uint, len(array))
	for i, v := range array {
		v2, ok := v.(uint)
		if !ok {
			return defaultValue
		}
		ret[i] = v2
	}
	return ret
}

func MustToUintArray(value interface{}) []uint {
	array, ok := toInterfaceArray(value)
	if !ok {
		panic(fmt.Sprintf("failed to cast %+v to []bool", value))
	}
	ret := make([]uint, len(array))
	for i, v := range array {
		v2, ok := v.(uint)
		if !ok {
			panic(fmt.Sprintf("failed to cast %+v to []uint", value))
		}
		ret[i] = v2
	}
	return ret
}

func ToUint8Array(value interface{}, defaultValue []uint8) []uint8 {
	array, ok := toInterfaceArray(value)
	if !ok {
		return defaultValue
	}
	ret := make([]uint8, len(array))
	for i, v := range array {
		v2, ok := v.(uint8)
		if !ok {
			return defaultValue
		}
		ret[i] = v2
	}
	return ret
}

func MustToUint8Array(value interface{}) []uint8 {
	array, ok := toInterfaceArray(value)
	if !ok {
		panic(fmt.Sprintf("failed to cast %+v to []bool", value))
	}
	ret := make([]uint8, len(array))
	for i, v := range array {
		v2, ok := v.(uint8)
		if !ok {
			panic(fmt.Sprintf("failed to cast %+v to []uint8", value))
		}
		ret[i] = v2
	}
	return ret
}

func ToUint16Array(value interface{}, defaultValue []uint16) []uint16 {
	array, ok := toInterfaceArray(value)
	if !ok {
		return defaultValue
	}
	ret := make([]uint16, len(array))
	for i, v := range array {
		v2, ok := v.(uint16)
		if !ok {
			return defaultValue
		}
		ret[i] = v2
	}
	return ret
}

func MustToUint16Array(value interface{}) []uint16 {
	array, ok := toInterfaceArray(value)
	if !ok {
		panic(fmt.Sprintf("failed to cast %+v to []bool", value))
	}
	ret := make([]uint16, len(array))
	for i, v := range array {
		v2, ok := v.(uint16)
		if !ok {
			panic(fmt.Sprintf("failed to cast %+v to []uint16", value))
		}
		ret[i] = v2
	}
	return ret
}

func ToUint32Array(value interface{}, defaultValue []uint32) []uint32 {
	array, ok := toInterfaceArray(value)
	if !ok {
		return defaultValue
	}
	ret := make([]uint32, len(array))
	for i, v := range array {
		v2, ok := v.(uint32)
		if !ok {
			return defaultValue
		}
		ret[i] = v2
	}
	return ret
}

func MustToUint32Array(value interface{}) []uint32 {
	array, ok := toInterfaceArray(value)
	if !ok {
		panic(fmt.Sprintf("failed to cast %+v to []bool", value))
	}
	ret := make([]uint32, len(array))
	for i, v := range array {
		v2, ok := v.(uint32)
		if !ok {
			panic(fmt.Sprintf("failed to cast %+v to []uint32", value))
		}
		ret[i] = v2
	}
	return ret
}

func ToUint64Array(value interface{}, defaultValue []uint64) []uint64 {
	array, ok := toInterfaceArray(value)
	if !ok {
		return defaultValue
	}
	ret := make([]uint64, len(array))
	for i, v := range array {
		v2, ok := v.(uint64)
		if !ok {
			return defaultValue
		}
		ret[i] = v2
	}
	return ret
}

func MustToUint64Array(value interface{}) []uint64 {
	array, ok := toInterfaceArray(value)
	if !ok {
		panic(fmt.Sprintf("failed to cast %+v to []bool", value))
	}
	ret := make([]uint64, len(array))
	for i, v := range array {
		v2, ok := v.(uint64)
		if !ok {
			panic(fmt.Sprintf("failed to cast %+v to []uint64", value))
		}
		ret[i] = v2
	}
	return ret
}

func ToFloat32Array(value interface{}, defaultValue []float32) []float32 {
	array, ok := toInterfaceArray(value)
	if !ok {
		return defaultValue
	}
	ret := make([]float32, len(array))
	for i, v := range array {
		v2, ok := v.(float32)
		if !ok {
			return defaultValue
		}
		ret[i] = v2
	}
	return ret
}

func MustToFloat32Array(value interface{}) []float32 {
	array, ok := toInterfaceArray(value)
	if !ok {
		panic(fmt.Sprintf("failed to cast %+v to []bool", value))
	}
	ret := make([]float32, len(array))
	for i, v := range array {
		v2, ok := v.(float32)
		if !ok {
			panic(fmt.Sprintf("failed to cast %+v to []float32", value))
		}
		ret[i] = v2
	}
	return ret
}

func ToFloat64Array(value interface{}, defaultValue []float64) []float64 {
	array, ok := toInterfaceArray(value)
	if !ok {
		return defaultValue
	}
	ret := make([]float64, len(array))
	for i, v := range array {
		v2, ok := v.(float64)
		if !ok {
			return defaultValue
		}
		ret[i] = v2
	}
	return ret
}

func MustToFloat64Array(value interface{}) []float64 {
	array, ok := toInterfaceArray(value)
	if !ok {
		panic(fmt.Sprintf("failed to cast %+v to []bool", value))
	}
	ret := make([]float64, len(array))
	for i, v := range array {
		v2, ok := v.(float64)
		if !ok {
			panic(fmt.Sprintf("failed to cast %+v to []float64", value))
		}
		ret[i] = v2
	}
	return ret
}

func ToStringArray(value interface{}, defaultValue []string) []string {
	array, ok := toInterfaceArray(value)
	if !ok {
		return defaultValue
	}
	ret := make([]string, len(array))
	for i, v := range array {
		v2, ok := v.(string)
		if !ok {
			return defaultValue
		}
		ret[i] = v2
	}
	return ret
}

func MustToStringArray(value interface{}) []string {
	array, ok := toInterfaceArray(value)
	if !ok {
		panic(fmt.Sprintf("failed to cast %+v to []bool", value))
	}
	ret := make([]string, len(array))
	for i, v := range array {
		v2, ok := v.(string)
		if !ok {
			panic(fmt.Sprintf("failed to cast %+v to []string", value))
		}
		ret[i] = v2
	}
	return ret
}

func ToMapArray(value interface{}, defaultValue []map[string]interface{}) []map[string]interface{} {
	array, ok := toInterfaceArray(value)
	if !ok {
		return defaultValue
	}
	ret := make([]map[string]interface{}, len(array))
	for i, v := range array {
		v2, ok := v.(map[string]interface{})
		if !ok {
			return defaultValue
		}
		ret[i] = v2
	}
	return ret
}

func MustToMapArray(value interface{}) []map[string]interface{} {
	array, ok := toInterfaceArray(value)
	if !ok {
		panic(fmt.Sprintf("failed to cast %+v to []bool", value))
	}
	ret := make([]map[string]interface{}, len(array))
	for i, v := range array {
		v2, ok := v.(map[string]interface{})
		if !ok {
			panic(fmt.Sprintf("failed to cast %+v to []map[string]interface{}", value))
		}
		ret[i] = v2
	}
	return ret
}

func ToTimeArray(value interface{}, defaultValue []time.Time) []time.Time {
	array, ok := toInterfaceArray(value)
	if !ok {
		return defaultValue
	}
	ret := make([]time.Time, len(array))
	for i, v := range array {
		v2, ok := v.(time.Time)
		if !ok {
			return defaultValue
		}
		ret[i] = v2
	}
	return ret
}

func MustToTimeArray(value interface{}) []time.Time {
	array, ok := toInterfaceArray(value)
	if !ok {
		panic(fmt.Sprintf("failed to cast %+v to []bool", value))
	}
	ret := make([]time.Time, len(array))
	for i, v := range array {
		v2, ok := v.(time.Time)
		if !ok {
			panic(fmt.Sprintf("failed to cast %+v to []time.Time", value))
		}
		ret[i] = v2
	}
	return ret
}

func ToComplex64Array(value interface{}, defaultValue []complex64) []complex64 {
	array, ok := toInterfaceArray(value)
	if !ok {
		return defaultValue
	}
	ret := make([]complex64, len(array))
	for i, v := range array {
		v2, ok := v.(complex64)
		if !ok {
			return defaultValue
		}
		ret[i] = v2
	}
	return ret
}

func MustToComplex64Array(value interface{}) []complex64 {
	array, ok := toInterfaceArray(value)
	if !ok {
		panic(fmt.Sprintf("failed to cast %+v to []bool", value))
	}
	ret := make([]complex64, len(array))
	for i, v := range array {
		v2, ok := v.(complex64)
		if !ok {
			panic(fmt.Sprintf("failed to cast %+v to []complex64", value))
		}
		ret[i] = v2
	}
	return ret
}

func ToComplex128Array(value interface{}, defaultValue []complex128) []complex128 {
	array, ok := toInterfaceArray(value)
	if !ok {
		return defaultValue
	}
	ret := make([]complex128, len(array))
	for i, v := range array {
		v2, ok := v.(complex128)
		if !ok {
			return defaultValue
		}
		ret[i] = v2
	}
	return ret
}

func MustToComplex128Array(value interface{}) []complex128 {
	array, ok := toInterfaceArray(value)
	if !ok {
		panic(fmt.Sprintf("failed to cast %+v to []bool", value))
	}
	ret := make([]complex128, len(array))
	for i, v := range array {
		v2, ok := v.(complex128)
		if !ok {
			panic(fmt.Sprintf("failed to cast %+v to []complex128", value))
		}
		ret[i] = v2
	}
	return ret
}

func ToRuneArray(value interface{}, defaultValue []rune) []rune {
	array, ok := toInterfaceArray(value)
	if !ok {
		return defaultValue
	}
	ret := make([]rune, len(array))
	for i, v := range array {
		v2, ok := v.(rune)
		if !ok {
			return defaultValue
		}
		ret[i] = v2
	}
	return ret
}

func MustToRuneArray(value interface{}) []rune {
	array, ok := toInterfaceArray(value)
	if !ok {
		panic(fmt.Sprintf("failed to cast %+v to []bool", value))
	}
	ret := make([]rune, len(array))
	for i, v := range array {
		v2, ok := v.(rune)
		if !ok {
			panic(fmt.Sprintf("failed to cast %+v to []rune", value))
		}
		ret[i] = v2
	}
	return ret
}

func ToUintptrArray(value interface{}, defaultValue []uintptr) []uintptr {
	array, ok := toInterfaceArray(value)
	if !ok {
		return defaultValue
	}
	ret := make([]uintptr, len(array))
	for i, v := range array {
		v2, ok := v.(uintptr)
		if !ok {
			return defaultValue
		}
		ret[i] = v2
	}
	return ret
}

func MustToUintptrArray(value interface{}) []uintptr {
	array, ok := toInterfaceArray(value)
	if !ok {
		panic(fmt.Sprintf("failed to cast %+v to []bool", value))
	}
	ret := make([]uintptr, len(array))
	for i, v := range array {
		v2, ok := v.(uintptr)
		if !ok {
			panic(fmt.Sprintf("failed to cast %+v to []uintptr", value))
		}
		ret[i] = v2
	}
	return ret
}
