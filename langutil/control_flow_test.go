package langutil

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestLoop(t *testing.T) {
	i := 0
	Loop(100, func() {
		i++
	})
	assert.Equal(t, i, 100)
}

func TestLoopI(t *testing.T) {
	j := 0
	LoopI(100, func(i int) {
		assert.Equal(t, i, j)
		j++
	})
	assert.Equal(t, j, 100)
}

func TestTernary(t *testing.T) {
	assert.Equal(t, Ternary(true, 123, 456), 123)
	assert.Equal(t, Ternary(false, 123, 456), 456)
}

func TestTernaryDo(t *testing.T) {
	assert.Equal(t, TernaryDo(true, func() interface{} {
		return 123
	}, func() interface{} {
		return 456
	}), 123)
	assert.Equal(t, TernaryDo(false, func() interface{} {
		return 123
	}, func() interface{} {
		return 456
	}), 456)
}
