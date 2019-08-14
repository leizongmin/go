package typeUtils

import (
	"fmt"
	"time"
)

func ToBoolArray(value interface{}, defaultValue []bool) []bool {
	if v, ok := value.([]bool); ok {
		return v
	} else {
		return defaultValue
	}
}

func MustToBoolArray(value interface{}) []bool {
	if v, ok := value.([]bool); ok {
		return v
	} else {
		panic(fmt.Sprintf("failed to cast %+v to []bool", value))
	}
}

func ToIntArray(value interface{}, defaultValue []int) []int {
	if v, ok := value.([]int); ok {
		return v
	} else {
		return defaultValue
	}
}

func MustToIntArray(value interface{}) []int {
	if v, ok := value.([]int); ok {
		return v
	} else {
		panic(fmt.Sprintf("failed to cast %+v to []int", value))
	}
}

func MustToInt8Array(value interface{}) []int8 {
	if v, ok := value.([]int8); ok {
		return v
	} else {
		panic(fmt.Sprintf("failed to cast %+v to []int8", value))
	}
}

func ToInt8Array(value interface{}, defaultValue []int8) []int8 {
	if v, ok := value.([]int8); ok {
		return v
	} else {
		return defaultValue
	}
}

func ToInt16Array(value interface{}, defaultValue []int16) []int16 {
	if v, ok := value.([]int16); ok {
		return v
	} else {
		return defaultValue
	}
}

func MustToInt16Array(value interface{}) []int16 {
	if v, ok := value.([]int16); ok {
		return v
	} else {
		panic(fmt.Sprintf("failed to cast %+v to []int16", value))
	}
}

func ToInt32Array(value interface{}, defaultValue []int32) []int32 {
	if v, ok := value.([]int32); ok {
		return v
	} else {
		return defaultValue
	}
}

func MustToInt32Array(value interface{}) []int32 {
	if v, ok := value.([]int32); ok {
		return v
	} else {
		panic(fmt.Sprintf("failed to cast %+v to []int32", value))
	}
}

func ToInt64Array(value interface{}, defaultValue []int64) []int64 {
	if v, ok := value.([]int64); ok {
		return v
	} else {
		return defaultValue
	}
}

func MustToInt64Array(value interface{}) []int64 {
	if v, ok := value.([]int64); ok {
		return v
	} else {
		panic(fmt.Sprintf("failed to cast %+v to []int64", value))
	}
}

func ToUintArray(value interface{}, defaultValue []uint) []uint {
	if v, ok := value.([]uint); ok {
		return v
	} else {
		return defaultValue
	}
}

func MustToUintArray(value interface{}) []uint {
	if v, ok := value.([]uint); ok {
		return v
	} else {
		panic(fmt.Sprintf("failed to cast %+v to []uint", value))
	}
}

func ToUint8Array(value interface{}, defaultValue []uint8) []uint8 {
	if v, ok := value.([]uint8); ok {
		return v
	} else {
		return defaultValue
	}
}

func MustToUint8Array(value interface{}) []uint8 {
	if v, ok := value.([]uint8); ok {
		return v
	} else {
		panic(fmt.Sprintf("failed to cast %+v to []uint8", value))
	}
}

func ToUint16Array(value interface{}, defaultValue []uint16) []uint16 {
	if v, ok := value.([]uint16); ok {
		return v
	} else {
		return defaultValue
	}
}

func MustToUint16Array(value interface{}) []uint16 {
	if v, ok := value.([]uint16); ok {
		return v
	} else {
		panic(fmt.Sprintf("failed to cast %+v to []uint16", value))
	}
}

func ToUint32Array(value interface{}, defaultValue []uint32) []uint32 {
	if v, ok := value.([]uint32); ok {
		return v
	} else {
		return defaultValue
	}
}

func MustToUint32Array(value interface{}) []uint32 {
	if v, ok := value.([]uint32); ok {
		return v
	} else {
		panic(fmt.Sprintf("failed to cast %+v to []uint32", value))
	}
}

func ToUint64Array(value interface{}, defaultValue []uint64) []uint64 {
	if v, ok := value.([]uint64); ok {
		return v
	} else {
		return defaultValue
	}
}

func MustToUint64Array(value interface{}) []uint64 {
	if v, ok := value.([]uint64); ok {
		return v
	} else {
		panic(fmt.Sprintf("failed to cast %+v to []uint64", value))
	}
}

func ToFloat32Array(value interface{}, defaultValue []float32) []float32 {
	if v, ok := value.([]float32); ok {
		return v
	} else {
		return defaultValue
	}
}

func MustToFloat32Array(value interface{}) []float32 {
	if v, ok := value.([]float32); ok {
		return v
	} else {
		panic(fmt.Sprintf("failed to cast %+v to []float32", value))
	}
}

func ToFloat64Array(value interface{}, defaultValue []float64) []float64 {
	if v, ok := value.([]float64); ok {
		return v
	} else {
		return defaultValue
	}
}

func MustToFloat64Array(value interface{}) []float64 {
	if v, ok := value.([]float64); ok {
		return v
	} else {
		panic(fmt.Sprintf("failed to cast %+v to []float64", value))
	}
}

func ToStringArray(value interface{}, defaultValue []string) []string {
	if v, ok := value.([]string); ok {
		return v
	} else {
		return defaultValue
	}
}

func MustToStringArray(value interface{}) []string {
	if v, ok := value.([]string); ok {
		return v
	} else {
		panic(fmt.Sprintf("failed to cast %+v to []string", value))
	}
}

func ToMapArray(value interface{}, defaultValue []map[string]interface{}) []map[string]interface{} {
	if v, ok := value.([]map[string]interface{}); ok {
		return v
	} else {
		return defaultValue
	}
}

func MustToMapArray(value interface{}) []map[string]interface{} {
	if v, ok := value.([]map[string]interface{}); ok {
		return v
	} else {
		panic(fmt.Sprintf("failed to cast %+v to []map[string]interface{}", value))
	}
}

func ToTimeArray(value interface{}, defaultValue []time.Time) []time.Time {
	if v, ok := value.([]time.Time); ok {
		return v
	} else {
		return defaultValue
	}
}

func MustToTimeArray(value interface{}) []time.Time {
	if v, ok := value.([]time.Time); ok {
		return v
	} else {
		panic(fmt.Sprintf("failed to cast %+v to []time.Time", value))
	}
}

func ToComplex64Array(value interface{}, defaultValue []complex64) []complex64 {
	if v, ok := value.([]complex64); ok {
		return v
	} else {
		return defaultValue
	}
}

func MustToComplex64Array(value interface{}) []complex64 {
	if v, ok := value.([]complex64); ok {
		return v
	} else {
		panic(fmt.Sprintf("failed to cast %+v to []complex64", value))
	}
}

func ToComplex128Array(value interface{}, defaultValue []complex128) []complex128 {
	if v, ok := value.([]complex128); ok {
		return v
	} else {
		return defaultValue
	}
}

func MustToComplex128Array(value interface{}) []complex128 {
	if v, ok := value.([]complex128); ok {
		return v
	} else {
		panic(fmt.Sprintf("failed to cast %+v to []complex128", value))
	}
}

func ToRuneArray(value interface{}, defaultValue []rune) []rune {
	if v, ok := value.([]rune); ok {
		return v
	} else {
		return defaultValue
	}
}

func MustToRuneArray(value interface{}) []rune {
	if v, ok := value.([]rune); ok {
		return v
	} else {
		panic(fmt.Sprintf("failed to cast %+v to []rune", value))
	}
}

func ToUintptrArray(value interface{}, defaultValue []uintptr) []uintptr {
	if v, ok := value.([]uintptr); ok {
		return v
	} else {
		return defaultValue
	}
}

func MustToUintptrArray(value interface{}) []uintptr {
	if v, ok := value.([]uintptr); ok {
		return v
	} else {
		panic(fmt.Sprintf("failed to cast %+v to []uintptr", value))
	}
}
