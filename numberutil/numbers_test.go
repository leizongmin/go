package numberutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTryParseInt(t *testing.T) {
	assert.Equal(t, 12345, TryParseInt("12345", 0))
	assert.Equal(t, 54321, TryParseInt("aaa", 54321))
}

func TestTryParseInt64(t *testing.T) {
	assert.Equal(t, int64(12345), TryParseInt64("12345", 0))
	assert.Equal(t, int64(54321), TryParseInt64("aaa", 54321))
}
