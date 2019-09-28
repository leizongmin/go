package typeUtils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsArrayOrSlice(t *testing.T) {
	assert.Equal(t, true, IsArray([1]string{"a"}))
	assert.Equal(t, true, IsSlice([]string{"a"}))
	assert.Equal(t, true, IsArrayOrSlice([1]string{"a"}))
	assert.Equal(t, true, IsArrayOrSlice([]string{"a"}))
}
