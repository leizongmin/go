package langutil

// 执行N次循环
func Loop(n int, f func()) {
	i := 0
	for i < n {
		f()
		i++
	}
}

// 执行N次循环，接收index参数
func LoopI(n int, f func(i int)) {
	i := 0
	for i < n {
		f(i)
		i++
	}
}
