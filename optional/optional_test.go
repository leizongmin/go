package optional

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

func TestMap(t *testing.T) {
	assert.Equal(t, Some(123).Map(func(value interface{}) interface{} {
		assert.Equal(t, value, 123)
		return 456
	}), Some(456))
	assert.Equal(t, None().Map(func(_ interface{}) interface{} {
		return 456
	}), None())
}

func TestFlatMap(t *testing.T) {
	assert.Equal(t, Some(123).FlatMap(func(value interface{}) Optional {
		assert.Equal(t, value, 123)
		return Some(456)
	}), Some(456))
	assert.Equal(t, None().FlatMap(func(_ interface{}) Optional {
		return Some(456)
	}), None())
}

func TestOrElse(t *testing.T) {
	assert.Equal(t, None().OrElse(456), 456)
	assert.Equal(t, Some(123).OrElse(456), 123)
}

func TestOrElseGet(t *testing.T) {
	assert.Equal(t, None().OrElseGet(func() interface{} {
		return 456
	}), 456)
	assert.Equal(t, Some(123).OrElseGet(func() interface{} {
		return 456
	}), 123)
}

func TestFilter(t *testing.T) {
	assert.Equal(t, None().Filter(func(value interface{}) bool {
		return true
	}), None())
	assert.Equal(t, None().Filter(func(value interface{}) bool {
		return false
	}), None())
	assert.Equal(t, Some(123).Filter(func(value interface{}) bool {
		assert.Equal(t, value, 123)
		return true
	}), Some(123))
	assert.Equal(t, Some(123).Filter(func(value interface{}) bool {
		assert.Equal(t, value, 123)
		return false
	}), None())
}
