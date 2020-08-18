package langutil

// 执行N次循环
func Loop(n int, f func()) {
	i := 0
	for i < n {
		i++
		f()
	}
}
