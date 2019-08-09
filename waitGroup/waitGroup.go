package waitGroup

import (
	"sync/atomic"
	"time"
)

type WaitGroupWithTimeout struct {
	sn       int64      // 序号，如果 done 的时候序号跟当前的不一致，则丢弃
	counter  int64      // 当前计数器
	maxCount int64      // 最大数量
	channel  chan int64 // 接收 done 的channel
}

// 重置计数
func (w *WaitGroupWithTimeout) Init(count int) {
	w.maxCount = int64(count)
	w.channel = make(chan int64, count)
	w.Reset(0)
}

// 重置计数，返回当前序号
func (w *WaitGroupWithTimeout) Reset(sn int) {
	atomic.StoreInt64(&w.sn, int64(sn))
	atomic.StoreInt64(&w.counter, 0)
}

// 等待结束
func (w *WaitGroupWithTimeout) Wait(timeout time.Duration) (count int, isTimeout bool) {
	t := time.After(timeout)
loop:
	for {
		select {
		case <-t:
			isTimeout = true
			break loop
		case sn := <-w.channel:
			if sn == w.sn {
				atomic.AddInt64(&w.counter, 1)
				count++
				// fmt.Printf("add %s %+v\n", time.Now(), w)
				if w.counter >= w.maxCount {
					break loop
				}
			} else {
				// fmt.Printf("ignore %+v\n", sn)
			}
		}
	}
	// fmt.Printf("end %+v\n%+v\n", w, t)
	return count, isTimeout
}

// 完成一个
func (w *WaitGroupWithTimeout) Done(sn int) {
	w.channel <- int64(sn)
}
