package optional

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOfNullable(t *testing.T) {
	assert.Equal(t, None(), OfNullable(nil))
	assert.Equal(t, Some(123), OfNullable(123))
}

func TestOfZeroable(t *testing.T) {
	assert.Equal(t, None(), OfZeroable(nil))
	assert.Equal(t, None(), OfZeroable(0))
	assert.Equal(t, None(), OfZeroable(0.0))
	assert.Equal(t, None(), OfZeroable(int8(0)))
	assert.Equal(t, None(), OfZeroable(int16(0)))
	assert.Equal(t, None(), OfZeroable(int32(0)))
	assert.Equal(t, None(), OfZeroable(int64(0)))
	assert.Equal(t, None(), OfZeroable(uint8(0)))
	assert.Equal(t, None(), OfZeroable(uint16(0)))
	assert.Equal(t, None(), OfZeroable(uint32(0)))
	assert.Equal(t, None(), OfZeroable(uint64(0)))
	assert.Equal(t, None(), OfZeroable(float32(0)))
	assert.Equal(t, None(), OfZeroable(float64(0)))
	assert.Equal(t, None(), OfZeroable(""))
	assert.Equal(t, None(), OfZeroable([]int{}))
	assert.Equal(t, None(), OfZeroable([0]int{}))
	assert.Equal(t, None(), OfZeroable(map[int]int{}))

	assert.Equal(t, Some(123), OfZeroable(123))
	assert.Equal(t, Some([]int{123}), OfZeroable([]int{123}))
}
