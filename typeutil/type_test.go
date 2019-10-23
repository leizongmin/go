package typeutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMustToInt(t *testing.T) {
	assert.Equal(t, 123, ToInt(123, 666))
	assert.Equal(t, 666, ToInt(123.0, 666))
}
