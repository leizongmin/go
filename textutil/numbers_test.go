package textutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTryParseInt(t *testing.T) {
	assert.Equal(t, 12345, TryParseInt("12345", 0))
	assert.Equal(t, 54321, TryParseInt("aaa", 54321))
}

func TestTryParseInt64(t *testing.T) {
	assert.Equal(t, int64(12345), TryParseInt64("12345", 0))
	assert.Equal(t, int64(54321), TryParseInt64("aaa", 54321))
}

func TestTryParseUint(t *testing.T) {
	assert.Equal(t, uint(12345), TryParseUint("12345", 0))
	assert.Equal(t, uint(54321), TryParseUint("aaa", 54321))
}

func TestTryParseUint64(t *testing.T) {
	assert.Equal(t, uint64(12345), TryParseUint64("12345", 0))
	assert.Equal(t, uint64(54321), TryParseUint64("aaa", 54321))
}

func TestTryParseFloat32(t *testing.T) {
	assert.Equal(t, float32(123.45), TryParseFloat32("123.45", 0))
	assert.Equal(t, float32(543.21), TryParseFloat32("aaa.bb", 543.21))
}

func TestTryParseFloat64(t *testing.T) {
	assert.Equal(t, 123.45, TryParseFloat64("123.45", 0))
	assert.Equal(t, 543.21, TryParseFloat64("aaa.bb", 543.21))
}

func TestTryParseBool(t *testing.T) {
	assert.Equal(t, true, TryParseBool("true", false))
	assert.Equal(t, false, TryParseBool("x", false))
}
