package benchmark

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSpentTime(t *testing.T) {
	ret := SpentTime("测试一下", func() {
		time.Sleep(time.Second)
	})
	assert.GreaterOrEqual(t, int64(ret), int64(time.Second))
}
