package benchmark

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSpent(t *testing.T) {
	ret := Spent("测试一下", func() {
		time.Sleep(time.Second)
	})
	assert.GreaterOrEqual(t, int64(ret), int64(time.Second))
}
