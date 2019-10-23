package typeutil

import (
	"fmt"
	"time"
)

func ToBool(value interface{}, defaultValue bool) bool {
	if v, ok := value.(bool); ok {
		return v
	} else {
		return defaultValue
	}
}

func MustToBool(value interface{}) bool {
	if v, ok := value.(bool); ok {
		return v
	} else {
		panic(fmt.Sprintf("failed to cast %+v to bool", value))
	}
}

func ToInt(value interface{}, defaultValue int) int {
	if v, ok := value.(int); ok {
		return v
	} else {
		return defaultValue
	}
}

func MustToInt(value interface{}) int {
	if v, ok := value.(int); ok {
		return v
	} else {
		panic(fmt.Sprintf("failed to cast %+v to int", value))
	}
}

func MustToInt8(value interface{}) int8 {
	if v, ok := value.(int8); ok {
		return v
	} else {
		panic(fmt.Sprintf("failed to cast %+v to int8", value))
	}
}

func ToInt8(value interface{}, defaultValue int8) int8 {
	if v, ok := value.(int8); ok {
		return v
	} else {
		return defaultValue
	}
}

func ToInt16(value interface{}, defaultValue int16) int16 {
	if v, ok := value.(int16); ok {
		return v
	} else {
		return defaultValue
	}
}

func MustToInt16(value interface{}) int16 {
	if v, ok := value.(int16); ok {
		return v
	} else {
		panic(fmt.Sprintf("failed to cast %+v to int16", value))
	}
}

func ToInt32(value interface{}, defaultValue int32) int32 {
	if v, ok := value.(int32); ok {
		return v
	} else {
		return defaultValue
	}
}

func MustToInt32(value interface{}) int32 {
	if v, ok := value.(int32); ok {
		return v
	} else {
		panic(fmt.Sprintf("failed to cast %+v to int32", value))
	}
}

func ToInt64(value interface{}, defaultValue int64) int64 {
	if v, ok := value.(int64); ok {
		return v
	} else {
		return defaultValue
	}
}

func MustToInt64(value interface{}) int64 {
	if v, ok := value.(int64); ok {
		return v
	} else {
		panic(fmt.Sprintf("failed to cast %+v to int64", value))
	}
}

func ToUint(value interface{}, defaultValue uint) uint {
	if v, ok := value.(uint); ok {
		return v
	} else {
		return defaultValue
	}
}

func MustToUint(value interface{}) uint {
	if v, ok := value.(uint); ok {
		return v
	} else {
		panic(fmt.Sprintf("failed to cast %+v to uint", value))
	}
}

func ToUint8(value interface{}, defaultValue uint8) uint8 {
	if v, ok := value.(uint8); ok {
		return v
	} else {
		return defaultValue
	}
}

func MustToUint8(value interface{}) uint8 {
	if v, ok := value.(uint8); ok {
		return v
	} else {
		panic(fmt.Sprintf("failed to cast %+v to uint8", value))
	}
}

func ToUint16(value interface{}, defaultValue uint16) uint16 {
	if v, ok := value.(uint16); ok {
		return v
	} else {
		return defaultValue
	}
}

func MustToUint16(value interface{}) uint16 {
	if v, ok := value.(uint16); ok {
		return v
	} else {
		panic(fmt.Sprintf("failed to cast %+v to uint16", value))
	}
}

func ToUint32(value interface{}, defaultValue uint32) uint32 {
	if v, ok := value.(uint32); ok {
		return v
	} else {
		return defaultValue
	}
}

func MustToUint32(value interface{}) uint32 {
	if v, ok := value.(uint32); ok {
		return v
	} else {
		panic(fmt.Sprintf("failed to cast %+v to uint32", value))
	}
}

func ToUint64(value interface{}, defaultValue uint64) uint64 {
	if v, ok := value.(uint64); ok {
		return v
	} else {
		return defaultValue
	}
}

func MustToUint64(value interface{}) uint64 {
	if v, ok := value.(uint64); ok {
		return v
	} else {
		panic(fmt.Sprintf("failed to cast %+v to uint64", value))
	}
}

func ToFloat32(value interface{}, defaultValue float32) float32 {
	if v, ok := value.(float32); ok {
		return v
	} else {
		return defaultValue
	}
}

func MustToFloat32(value interface{}) float32 {
	if v, ok := value.(float32); ok {
		return v
	} else {
		panic(fmt.Sprintf("failed to cast %+v to float32", value))
	}
}

func ToFloat64(value interface{}, defaultValue float64) float64 {
	if v, ok := value.(float64); ok {
		return v
	} else {
		return defaultValue
	}
}

func MustToFloat64(value interface{}) float64 {
	if v, ok := value.(float64); ok {
		return v
	} else {
		panic(fmt.Sprintf("failed to cast %+v to float64", value))
	}
}

func ToString(value interface{}, defaultValue string) string {
	if v, ok := value.(string); ok {
		return v
	} else {
		return defaultValue
	}
}

func MustToString(value interface{}) string {
	if v, ok := value.(string); ok {
		return v
	} else {
		panic(fmt.Sprintf("failed to cast %+v to string", value))
	}
}

func ToMap(value interface{}, defaultValue map[string]interface{}) map[string]interface{} {
	if v, ok := value.(map[string]interface{}); ok {
		return v
	} else {
		return defaultValue
	}
}

func MustToMap(value interface{}) map[string]interface{} {
	if v, ok := value.(map[string]interface{}); ok {
		return v
	} else {
		panic(fmt.Sprintf("failed to cast %+v to map[string]interface{}", value))
	}
}

func ToTime(value interface{}, defaultValue time.Time) time.Time {
	if v, ok := value.(time.Time); ok {
		return v
	} else {
		return defaultValue
	}
}

func MustToTime(value interface{}) time.Time {
	if v, ok := value.(time.Time); ok {
		return v
	} else {
		panic(fmt.Sprintf("failed to cast %+v to time.Time", value))
	}
}

func ToComplex64(value interface{}, defaultValue complex64) complex64 {
	if v, ok := value.(complex64); ok {
		return v
	} else {
		return defaultValue
	}
}

func MustToComplex64(value interface{}) complex64 {
	if v, ok := value.(complex64); ok {
		return v
	} else {
		panic(fmt.Sprintf("failed to cast %+v to complex64", value))
	}
}

func ToComplex128(value interface{}, defaultValue complex128) complex128 {
	if v, ok := value.(complex128); ok {
		return v
	} else {
		return defaultValue
	}
}

func MustToComplex128(value interface{}) complex128 {
	if v, ok := value.(complex128); ok {
		return v
	} else {
		panic(fmt.Sprintf("failed to cast %+v to complex128", value))
	}
}

func ToByte(value interface{}, defaultValue byte) byte {
	if v, ok := value.(byte); ok {
		return v
	} else {
		return defaultValue
	}
}

func MustToByte(value interface{}) byte {
	if v, ok := value.(byte); ok {
		return v
	} else {
		panic(fmt.Sprintf("failed to cast %+v to byte", value))
	}
}

func ToRune(value interface{}, defaultValue rune) rune {
	if v, ok := value.(rune); ok {
		return v
	} else {
		return defaultValue
	}
}

func MustToRune(value interface{}) rune {
	if v, ok := value.(rune); ok {
		return v
	} else {
		panic(fmt.Sprintf("failed to cast %+v to rune", value))
	}
}

func ToUintptr(value interface{}, defaultValue uintptr) uintptr {
	if v, ok := value.(uintptr); ok {
		return v
	} else {
		return defaultValue
	}
}

func MustToUintptr(value interface{}) uintptr {
	if v, ok := value.(uintptr); ok {
		return v
	} else {
		panic(fmt.Sprintf("failed to cast %+v to uintptr", value))
	}
}
