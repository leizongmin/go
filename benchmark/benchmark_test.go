package benchmark

import (
	"github.com/go-playground/assert/v2"
	"go.uber.org/atomic"
	"testing"
	"time"
)

func TestDoAndPrint(t *testing.T) {
	count := 0
	DoAndPrint(100, "测试", func() {
		count++
		time.Sleep(time.Nanosecond * 100)
	})
	assert.Equal(t, 100, count)
}

func TestDoParallelAndPrint(t *testing.T) {
	count := atomic.NewInt32(0)
	DoParallelAndPrint(100, 1000, "测试", func() {
		count.Inc()
		time.Sleep(time.Nanosecond * 100)
	})
	assert.Equal(t, int32(1000), count.Load())
}
