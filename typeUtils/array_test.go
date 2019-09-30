package typeUtils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToBoolArray(t *testing.T) {
	ret := ToBoolArray([]interface{}{true, false}, nil)
	assert.True(t, ret[0])
	assert.False(t, ret[1])
	assert.Equal(t, []bool{true, false}, ret)
}

func TestToInterfaceArray(t *testing.T) {
	ret, ok := ToInterfaceArray([]interface{}{true, false, 123, "ok"})
	assert.True(t, ok)
	assert.Equal(t, true, ret[0])
	assert.Equal(t, false, ret[1])
	assert.Equal(t, 123, ret[2])
	assert.Equal(t, "ok", ret[3])
}

func TestToStringArray(t *testing.T) {
	ret := ToStringArray([]interface{}{"A", "B"}, nil)
	assert.NotNil(t, ret)
	assert.Equal(t, "A", ret[0])
	assert.Equal(t, "B", ret[1])
	assert.Equal(t, []string{"A", "B"}, ret)
}

func TestToStringArray2(t *testing.T) {
	ret := ToStringArray([]interface{}{interface{}("A"), interface{}("B")}, nil)
	assert.NotNil(t, ret)
	assert.Equal(t, "A", ret[0])
	assert.Equal(t, "B", ret[1])
	assert.Equal(t, []string{"A", "B"}, ret)
	fmt.Printf("%+v\n", ret)
}
