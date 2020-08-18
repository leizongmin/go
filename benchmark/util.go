package benchmark

import (
	"fmt"
	"time"
)

// 打印执行函数消耗的时间
func SpentTime(msg string, f func()) time.Duration {
	start := time.Now()
	f()
	spent := time.Now().Sub(start)
	fmt.Printf("[spent:%s]\t%s\n", spent, msg)
	return spent
}
