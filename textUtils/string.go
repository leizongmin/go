package textUtils

import (
	"strings"
)

// 是否为空白字符
func IsWhitesapce(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

// 是否为空格
func IsSpace(ch rune) bool {
	return ch == ' '
}

// 是否为字母
func IsLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

// 是否为大写字母
func IsUpperLetter(ch rune) bool {
	return ch >= 'A' && ch <= 'Z'
}

// 是否为小写字母
func IsLowerLetter(ch rune) bool {
	return ch >= 'a' && ch <= 'z'
}

// 是否为数字
func IsDigit(ch rune) bool {
	return ch >= '0' && ch <= '9'
}

// 是否以指定字符串开头
func StartsWith(s string, substr string) bool {
	return strings.Index(s, substr) == 0
}
