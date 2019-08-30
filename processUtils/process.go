package processUtils

// 无限等待
func InfiniteWait() {
	c := make(chan bool, 0)
	<-c
}
