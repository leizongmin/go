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
