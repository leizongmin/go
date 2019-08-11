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
}
