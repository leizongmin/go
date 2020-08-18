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

// 三目运算
func Ternary(ok bool, a interface{}, b interface{}) interface{} {
	if ok {
		return a
	}
	return b
}

// 三目运算，执行函数
func TernaryDo(ok bool, doA func() interface{}, doB func() interface{}) interface{} {
	if ok {
		return doA()
	}
	return doB()
}
