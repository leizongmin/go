package waitGroup

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestWithTimeout_WaitWithTimeout(t *testing.T) {
	wg := WithTimeout{}
	wg.Init(5)
	{
		sn := 1
		wg.Reset(sn)
		for i := 0; i < 5; i++ {
			go func() {
				time.Sleep(time.Millisecond * 1)
				wg.Done(sn)
			}()
		}
		count, isTimeout := wg.WaitWithTimeout(time.Millisecond * 20)
		assert.Equal(t, 5, count)
		assert.Equal(t, false, isTimeout)
	}
	time.Sleep(time.Millisecond * 200)
	{
		sn := 2
		wg.Reset(sn)
		for i := 0; i < 10; i++ {
			go func() {
				time.Sleep(time.Millisecond * 1)
				wg.Done(sn)
			}()
		}
		count, isTimeout := wg.WaitWithTimeout(time.Millisecond * 20)
		assert.Equal(t, 5, count)
		assert.Equal(t, false, isTimeout)
	}
	time.Sleep(time.Millisecond * 200)
	{
		sn := 3
		wg.Reset(sn)
		for i := 0; i < 10; i++ {
			go func() {
				time.Sleep(time.Millisecond * 50)
				wg.Done(sn)
			}()
		}
		count, isTimeout := wg.WaitWithTimeout(time.Millisecond * 20)
		assert.Equal(t, 0, count)
		assert.Equal(t, true, isTimeout)
	}
	time.Sleep(time.Millisecond * 200)
	{
		sn := 4
		wg.Reset(sn)
		for i := 0; i < 10; i++ {
			go func(i int) {
				time.Sleep(time.Millisecond * time.Duration(i*8))
				wg.Done(sn)
			}(i)
		}
		count, isTimeout := wg.WaitWithTimeout(time.Millisecond * 20)
		assert.Equal(t, true, count < 5)
		assert.Equal(t, true, isTimeout)
	}
}

func TestWithTimeout_Wait(t *testing.T) {
	wg := New()
	wg.Init(5)
	ch := make(chan interface{}, 0)
	go func() {
		sn := 1
		wg.Reset(sn)
		for i := 0; i < 5; i++ {
			go func(i int) {
				time.Sleep(time.Millisecond * 10 * time.Duration(i+1))
				wg.Done(sn)
			}(i)
		}
		count, isCancel := wg.Wait(ch)
		assert.Equal(t, 2, count)
		assert.Equal(t, true, isCancel)
	}()
	time.Sleep(time.Millisecond * 25)
	ch <- 1
	time.Sleep(time.Millisecond * 100)
}
