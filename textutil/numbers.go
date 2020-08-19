package textutil

import (
	"strconv"
)

// 尝试解析int，如果失败则返回默认值
func TryParseInt(s string, defaultValue int) int {
	v, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return defaultValue
	}
	return int(v)
}

// 尝试解析int64，如果失败则返回默认值
func TryParseInt64(s string, defaultValue int64) int64 {
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return defaultValue
	}
	return v
}

// 尝试解析uint，如果失败则返回默认值
func TryParseUint(s string, defaultValue uint) uint {
	v, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return defaultValue
	}
	return uint(v)
}

// 尝试解析uint64，如果失败则返回默认值
func TryParseUint64(s string, defaultValue uint64) uint64 {
	v, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return defaultValue
	}
	return v
}

// 尝试解析float32，如果失败则返回默认值
func TryParseFloat32(s string, defaultValue float32) float32 {
	v, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return defaultValue
	}
	return float32(v)
}

// 尝试解析float64，如果失败则返回默认值
func TryParseFloat64(s string, defaultValue float64) float64 {
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return defaultValue
	}
	return v
}

// 尝试解析bool，如果失败则返回默认值
func TryParseBool(s string, defaultValue bool) bool {
	v, err := strconv.ParseBool(s)
	if err != nil {
		return defaultValue
	}
	return v
}
