package iterutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	a := Array{1, 2, 3, 4, 5}
	b := Map(a, func(value Item) Item {
		return value.(int) * 2
	})
	assert.Equal(t, b, Array{2, 4, 6, 8, 10})
}

func TestFilter(t *testing.T) {
	a := Array{1, 2, 3, 4, 5}
	b := Filter(a, func(value Item) bool {
		if value.(int)%2 == 0 {
			return true
		}
		return false
	})
	assert.Equal(t, b, Array{2, 4})
}

func TestReduce(t *testing.T) {
	a := Array{1, 2, 3, 4, 5}
	b := Reduce(a, 1, func(total, currentValue Item, currentIndex int) Item {
		return total.(int) * currentValue.(int)
	})
	assert.Equal(t, b, 1*2*3*4*5)
}
