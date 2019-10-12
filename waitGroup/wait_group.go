package waitGroup

import (
	"sync/atomic"
	"time"
)

type WithTimeout struct {
	sn       int64      // 序号，如果 done 的时候序号跟当前的不一致，则丢弃
	counter  int64      // 当前计数器
	maxCount int64      // 最大数量
	channel  chan int64 // 接收 done 的channel
}

// 新建waitGroup
func New() *WithTimeout {
	return &WithTimeout{}
}

// 重置计数
func (w *WithTimeout) Init(count int) *WithTimeout {
	w.maxCount = int64(count)
	w.channel = make(chan int64, count)
	w.Reset(0)
	return w
}

// 重置计数，返回当前序号
func (w *WithTimeout) Reset(sn int) *WithTimeout {
	atomic.StoreInt64(&w.sn, int64(sn))
	atomic.StoreInt64(&w.counter, 0)
	return w
}

// 等待结束，如果超时则立刻返回
func (w *WithTimeout) WaitWithTimeout(timeout time.Duration) (count int, isTimeout bool) {
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

// 等待结束，如果收到指定信号则立刻返回
func (w *WithTimeout) Wait(ch <-chan interface{}) (count int, isCancel bool) {
loop:
	for {
		select {
		case <-ch:
			isCancel = true
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
	return count, isCancel
}

// 无限等待结束
func (w *WithTimeout) WaitInfinity() (count int) {
	count, _ = w.Wait(make(chan interface{}))
	return count
}

// 完成一个
func (w *WithTimeout) Done(sn int) *WithTimeout {
	w.channel <- int64(sn)
	return w
}
