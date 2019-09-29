package textUtils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	assert.True(t, IsDigit('1'))
	assert.True(t, IsLetter('x'))
	assert.True(t, IsWhitesapce('\n'))
	assert.True(t, IsSpace(' '))
	assert.True(t, IsLowerLetter('x'))
	assert.True(t, IsUpperLetter('X'))

	assert.False(t, IsDigit('x'))
	assert.False(t, IsLetter('1'))
	assert.False(t, IsWhitesapce('2'))
	assert.False(t, IsSpace('3'))
	assert.False(t, IsLowerLetter('X'))
	assert.False(t, IsUpperLetter('x'))

	assert.True(t, StartsWith("abc", "a"))
	assert.False(t, StartsWith("abc", "A"))
	assert.False(t, StartsWith("abc", "b"))

	assert.Equal(t, `xxx`, AnythingToString("xxx"))
	assert.Equal(t, `123`, AnythingToString(123))
	assert.Equal(t, `123`, AnythingToString(123.0))
	assert.Equal(t, `{"a":123}`, AnythingToString(map[string]interface{}{"a": 123}))
	assert.Equal(t, `["a","b"]`, AnythingToString([]string{"a", "b"}))
	assert.Equal(t, `[123,456]`, AnythingToString([]int{123, 456}))
	assert.Equal(t, `true`, AnythingToString(true))
	assert.Equal(t, `false`, AnythingToString(false))
	assert.Equal(t, ``, AnythingToString(nil))
	assert.Equal(t, `{}`, AnythingToString(struct{}{}))
	assert.Equal(t, `{a:123}`, AnythingToString(struct{ a int }{a: 123}))
}
