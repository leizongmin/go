package processUtils

import (
	"os"
)

// 无限等待
func InfiniteWait() {
	c := make(chan bool, 0)
	<-c
}

// 结束程序
func Exit(code int) {
	os.Exit(code)
}
