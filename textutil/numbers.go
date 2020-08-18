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
	return int64(v)
}
