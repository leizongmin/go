package timer

import (
	"time"
)

type Timer struct {
	Counter  int  // 已执行次数
	Canceled bool // 是否已取消
	Ended    bool // 是否已结束
}

// Cancel 取消任务
func (t *Timer) Cancel() *Timer {
	t.Canceled = true
	return t
}

// SetTimeout 延时执行任务
func SetTimeout(h func(), d time.Duration) *Timer {
	t := &Timer{}
	go func() {
		time.Sleep(d)
		if !t.Canceled {
			h()
			t.Counter++
		}
		t.Ended = true
	}()
	return t
}

// SetInterval 循环某个周期执行任务
func SetInterval(h func(), d time.Duration) *Timer {
	t := &Timer{}
	go func() {
		for !t.Canceled {
			time.Sleep(d)
			h()
			t.Counter++
		}
		t.Ended = true
	}()
	return t
}
