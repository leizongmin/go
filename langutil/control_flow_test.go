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
