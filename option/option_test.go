package option

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNone(t *testing.T) {
	a := None()
	b := None()
	assert.Equal(t, a, b)
	assert.Equal(t, true, a.IsNone)
	assert.Equal(t, false, a.IsSome)
	assert.Equal(t, true, b.IsNone)
	assert.Equal(t, false, b.IsSome)
}

func TestSome(t *testing.T) {
	a := Some(123)
	b := Some(456)
	assert.NotEqual(t, a, b)
	assert.Equal(t, true, a.IsSome)
	assert.Equal(t, false, a.IsNone)
	assert.Equal(t, 123, a.Value)
	assert.Equal(t, true, b.IsSome)
	assert.Equal(t, false, b.IsNone)
	assert.Equal(t, 456, b.Value)
}
