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
