package typeutil

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAny(t *testing.T) {
	a := Any(map[string]interface{}{"a": 123})
	fmt.Println(a.Get(1))
	fmt.Println(a.Get("a"))
	fmt.Println(a.Get("b"))
	{
		v, ok := a.Get("a")
		assert.True(t, ok)
		assert.Equal(t, 123, v.MustToInt())
		v2, ok := a.Get("b")
		assert.False(t, ok)
		assert.Nil(t, v2.Value())
	}

	b := Any([2]string{"a", "b"})
	fmt.Println(b.Get(1))
	fmt.Println(b.Get(10))
	fmt.Println(b.Get("a"))
	fmt.Println(b.Get("b"))

	c := Any(append([]string{}, "a", "b"))
	fmt.Println(c.Get(1))
	fmt.Println(c.Get(10))
	fmt.Println(c.Get("a"))
	fmt.Println(c.Get("b"))
}
