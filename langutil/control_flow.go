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
func If(condition bool, thenValue interface{}, elseValue interface{}) interface{} {
	if condition {
		return thenValue
	}
	return elseValue
}

// 三目运算，执行函数
func IfDo(condition bool, thenDo func() interface{}, elseDo func() interface{}) interface{} {
	if condition {
		return thenDo()
	}
	return elseDo()
}
