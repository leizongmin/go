package tryRecover

import (
	"log"
	"runtime/debug"
)

type LogFunction = func(format string, args ...interface{})

var logPrintf LogFunction

func init() {
	logPrintf = log.Printf
}

// 设置打印日志函数
func SetLogFunction(f LogFunction) {
	logPrintf = f
}

// 尝试捕捉panic，并吃掉
// 使用方法：defer Eat()
func Eat() {
	err := recover()
	if err != nil {
		logPrintf("try recover: %+v\n%s", err, string(debug.Stack()))
	}
}
