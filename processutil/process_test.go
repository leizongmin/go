package processutil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSourceFileInfo(t *testing.T) {
	funcName, fileName, fileLine, err := SourceFileInfo()
	assert.NoError(t, err)
	assert.Equal(t, "TestSourceFileInfo", funcName)
	assert.Contains(t, fileName, "/processutil/process_test.go")
	assert.Equal(t, 9, fileLine)
}
