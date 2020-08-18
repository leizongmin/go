package timeutil

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSetTimeout(t *testing.T) {
	{
		var a bool
		b := SetTimeout(func() {
			a = true
		}, time.Millisecond*100)
		time.Sleep(time.Millisecond * 200)
		assert.True(t, a)
		assert.Equal(t, b.Counter, 1)
	}
	{
		var a bool
		b := SetTimeout(func() {
			a = true
		}, time.Millisecond*100)
		time.Sleep(time.Millisecond * 50)
		b.Cancel()
		assert.False(t, a)
		assert.True(t, b.Canceled)
		assert.False(t, b.Ended)
		time.Sleep(time.Millisecond * 100)
		assert.True(t, b.Ended)
	}
}

func TestSetInterval(t *testing.T) {
	{
		var counter int
		b := SetInterval(func() {
			counter++
		}, time.Millisecond*100)
		time.Sleep(time.Millisecond * 800)
		assert.True(t, counter > 5)
		assert.True(t, b.Counter > 5)
	}
	{
		var counter int
		b := SetInterval(func() {
			counter++
		}, time.Millisecond*100)
		time.Sleep(time.Millisecond * 250)
		b.Cancel()
		assert.Equal(t, counter, 2)
		assert.Equal(t, b.Counter, 2)
		assert.True(t, b.Canceled)
		assert.False(t, b.Ended)
		time.Sleep(time.Millisecond * 100)
		assert.True(t, b.Ended)
	}
}
