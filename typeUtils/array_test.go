package typeUtils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToBoolArray(t *testing.T) {
	ret := ToBoolArray([]interface{}{true, false}, nil)
	assert.True(t, ret[0])
	assert.False(t, ret[1])
}

func TestToInterfaceArray(t *testing.T) {
	ret, ok := ToInterfaceArray([]interface{}{true, false, 123, "ok"})
	assert.True(t, ok)
	assert.Equal(t, true, ret[0])
	assert.Equal(t, false, ret[1])
	assert.Equal(t, 123, ret[2])
	assert.Equal(t, "ok", ret[3])
}
