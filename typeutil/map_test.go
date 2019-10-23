package typeutil

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStructToMap(t *testing.T) {
	a := struct {
		A int
		B float64
		C string
	}{
		A: 123,
		B: 666,
		C: "ok",
	}
	b, err := StructToMap(&a)
	assert.NoError(t, err)
	fmt.Printf("%+v\n", b)
	assert.Equal(t, float64(a.A), b["A"])
	assert.Equal(t, a.B, b["B"])
	assert.Equal(t, a.C, b["C"])
}
